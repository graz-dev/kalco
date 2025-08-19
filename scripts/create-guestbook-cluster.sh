#!/bin/bash

# Kalco Guestbook Cluster Setup Script
# This script creates a Kind cluster with the Guestbook application
# Based on: https://kubernetes.io/docs/tutorials/stateless-application/guestbook/

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to wait for pods to be ready
wait_for_pods() {
    local namespace=$1
    local label_selector=$2
    local expected_count=$3
    local timeout=300
    local elapsed=0
    
    print_status "Waiting for pods in namespace '$namespace' with selector '$label_selector' to be ready..."
    
    while [ $elapsed -lt $timeout ]; do
        local ready_count=$(kubectl get pods -n "$namespace" -l "$label_selector" -o jsonpath='{.items[*].status.containerStatuses[*].ready}' 2>/dev/null | tr ' ' '\n' | grep -c "true" 2>/dev/null || echo "0")
        
        # Ensure ready_count is a valid number
        if [[ "$ready_count" =~ ^[0-9]+$ ]]; then
            if [ "$ready_count" -eq "$expected_count" ]; then
                print_success "All $expected_count pods are ready!"
                return 0
            fi
            
            print_status "Ready: $ready_count/$expected_count pods (waiting...)"
        else
            print_status "Waiting for pods to initialize... (ready: 0/$expected_count)"
        fi
        
        sleep 10
        elapsed=$((elapsed + 10))
    done
    
    print_error "Timeout waiting for pods to be ready after ${timeout}s"
    return 1
}

# Function to check prerequisites
check_prerequisites() {
    print_status "Checking prerequisites..."
    
    # Check if kind is installed
    if ! command_exists kind; then
        print_error "Kind is not installed. Please install it first:"
        echo "  curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64"
        echo "  chmod +x ./kind"
        echo "  sudo mv ./kind /usr/local/bin/kind"
        exit 1
    fi
    
    # Check if kubectl is installed
    if ! command_exists kubectl; then
        print_error "kubectl is not installed. Please install it first:"
        echo "  curl -LO \"https://dl.k8s.io/release/\$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\""
        echo "  chmod +x kubectl"
        echo "  sudo mv kubectl /usr/local/bin/kubectl"
        exit 1
    fi
    
    # Check if Docker is running
    if ! docker info >/dev/null 2>&1; then
        print_error "Docker is not running. Please start Docker first."
        exit 1
    fi
    
    print_success "All prerequisites are satisfied!"
}

# Function to create Kind cluster
create_cluster() {
    local cluster_name=${1:-"guestbook-cluster"}
    
    print_status "Creating Kind cluster '$cluster_name'..."
    
    # Check if cluster already exists
    if kind get clusters | grep -q "^$cluster_name$"; then
        print_warning "Cluster '$cluster_name' already exists!"
        read -p "Do you want to delete and recreate it? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            print_status "Deleting existing cluster..."
            kind delete cluster --name "$cluster_name"
        else
            print_status "Using existing cluster '$cluster_name'"
            return 0
        fi
    fi
    
    # Create cluster configuration
    cat > /tmp/kind-config.yaml << EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: $cluster_name
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 80
    hostPort: 8080
    protocol: TCP
- role: worker
- role: worker
EOF
    
    # Create the cluster
    kind create cluster --config /tmp/kind-config.yaml --name "$cluster_name"
    
    # Clean up config file
    rm -f /tmp/kind-config.yaml
    
    print_success "Cluster '$cluster_name' created successfully!"
}

# Function to deploy Redis
deploy_redis() {
    print_status "Deploying Redis database..."
    
    # Create Redis leader deployment
    cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-leader
  labels:
    app: redis
    tier: backend
    role: leader
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
      tier: backend
      role: leader
  template:
    metadata:
      labels:
        app: redis
        tier: backend
        role: leader
    spec:
      containers:
      - name: redis
        image: redis:6.0.5
        ports:
        - containerPort: 6379
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
EOF
    
    # Create Redis leader service
    cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: redis-leader
  labels:
    app: redis
    tier: backend
    role: leader
spec:
  ports:
  - port: 6379
    targetPort: 6379
  selector:
    app: redis
    tier: backend
    role: leader
EOF
    
    # Create Redis follower deployment
    cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-follower
  labels:
    app: redis
    tier: backend
    role: follower
spec:
  replicas: 2
  selector:
    matchLabels:
      app: redis
      tier: backend
      role: follower
  template:
    metadata:
      labels:
        app: redis
        tier: backend
        role: follower
    spec:
      containers:
      - name: redis
        image: redis:6.0.5
        ports:
        - containerPort: 6379
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
EOF
    
    # Create Redis follower service
    cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: redis-follower
  labels:
    app: redis
    tier: backend
    role: follower
spec:
  ports:
  - port: 6379
    targetPort: 6379
  selector:
    app: redis
    tier: backend
    role: follower
EOF
    
    print_success "Redis deployed successfully!"
}

# Function to deploy Guestbook frontend
deploy_guestbook() {
    print_status "Deploying Guestbook frontend..."
    
    # Create frontend deployment
    cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  labels:
    app: guestbook
    tier: frontend
spec:
  replicas: 3
  selector:
    matchLabels:
        app: guestbook
        tier: frontend
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: us-docker.pkg.dev/google-samples/containers/gke/gb-frontend:v5
        env:
        - name: GET_HOSTS_FROM
          value: "dns"
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        ports:
        - containerPort: 80
EOF
    
    # Create frontend service
    cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: frontend
  labels:
    app: guestbook
    tier: frontend
spec:
  ports:
  - port: 80
  selector:
    app: guestbook
    tier: frontend
EOF
    
    print_success "Guestbook frontend deployed successfully!"
}

# Function to wait for all components
wait_for_components() {
    print_status "Waiting for all components to be ready..."
    
    # Wait for Redis leader
    wait_for_pods "default" "app=redis,tier=backend,role=leader" 1
    
    # Wait for Redis followers
    wait_for_pods "default" "app=redis,tier=backend,role=follower" 2
    
    # Wait for frontend
    wait_for_pods "default" "app=guestbook,tier=frontend" 3
    
    print_success "All components are ready!"
}

# Function to show cluster information
show_cluster_info() {
    print_status "Cluster Information:"
    echo "  Cluster Name: $(kind get clusters | head -1)"
    echo "  Kubernetes Version: $(kubectl version --short 2>/dev/null | grep 'Server Version' | awk '{print $3}' || echo 'Unknown')"
    echo "  Nodes: $(kubectl get nodes --no-headers | wc -l)"
    
    echo
    print_status "Application Status:"
    kubectl get pods -l app=redis
    kubectl get pods -l app=guestbook
    kubectl get services
    
    echo
    print_status "Access Information:"
    echo "  Frontend Service: kubectl port-forward svc/frontend 8080:80"
    echo "  Redis Leader: kubectl port-forward svc/redis-leader 6379:6379"
    echo "  Redis Follower: kubectl port-forward svc/redis-follower 6380:6379"
    echo
    echo "  Open http://localhost:8080 in your browser to access the Guestbook application"
}

# Function to create Kalco context
create_kalco_context() {
    # Try to find kalco in PATH or use local binary
    local kalco_cmd="kalco"
    if ! command_exists kalco; then
        if [ -f "./kalco" ]; then
            kalco_cmd="./kalco"
            print_status "Using local Kalco binary"
        else
            print_warning "Kalco is not installed and no local binary found. Install it to manage this cluster with Kalco."
            return 0
        fi
    fi
    
    print_status "Creating Kalco context for the guestbook cluster..."
    
    # Create context
    if $kalco_cmd context set guestbook-cluster \
      --kubeconfig ~/.kube/config \
      --output ./guestbook-exports \
      --description "Kind cluster with Guestbook application" \
      --labels env=dev,app=guestbook,cluster=kind; then
        
        print_success "Kalco context 'guestbook-cluster' created!"
        echo "  Use '$kalco_cmd context use guestbook-cluster' to switch to this context"
        echo "  Use '$kalco_cmd export' to export the cluster resources"
    else
        print_warning "Failed to create Kalco context. You can create it manually later."
    fi
}

# Main execution
main() {
    echo "=========================================="
    echo "  Kalco Guestbook Cluster Setup Script"
    echo "=========================================="
    echo
    
    # Check prerequisites
    check_prerequisites
    
    # Create cluster
    create_cluster
    
    # Deploy Redis
    deploy_redis
    
    # Deploy Guestbook frontend
    deploy_guestbook
    
    # Wait for components
    wait_for_components
    
    # Show cluster information
    show_cluster_info
    
    # Create Kalco context
    create_kalco_context
    
    echo
    echo "=========================================="
    print_success "Guestbook cluster setup completed!"
    echo "=========================================="
    echo
    echo "The cluster will remain running. To stop it manually, run:"
    echo "  kind delete cluster --name guestbook-cluster"
    echo
    echo "To access the application:"
    echo "  kubectl port-forward svc/frontend 8080:80"
    echo "  Then open http://localhost:8080 in your browser"
    echo
    echo "To export the cluster with Kalco:"
    echo "  kalco context use guestbook-cluster"
    echo "  kalco export --git-push --commit-message 'Initial guestbook cluster export'"
}

# Run main function
main "$@"
