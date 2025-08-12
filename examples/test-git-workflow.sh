#!/bin/bash

echo "🚀 Kalco Git Workflow Test with KIND Cluster"
echo "============================================="
echo

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_step() {
    echo -e "${BLUE}📋 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Configuration
CLUSTER_NAME="kalco-git-test"
OUTPUT_DIR="./cluster-history"
NAMESPACE="demo-apps"

print_step "Step 1: Creating KIND cluster for testing"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step creates a new Kubernetes cluster using KIND (Kubernetes in Docker)."
echo "The cluster will be used to demonstrate kalco's Git version control capabilities."
echo

if ! command -v kind &> /dev/null; then
    print_error "KIND is not installed. Please install KIND first:"
    echo "   brew install kind  # macOS"
    echo "   # or visit: https://kind.sigs.k8s.io/docs/user/quick-start/"
    exit 1
fi

echo "Creating KIND cluster: $CLUSTER_NAME"
kind create cluster --name "$CLUSTER_NAME" --config - <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
EOF

if [ $? -eq 0 ]; then
    print_success "KIND cluster created successfully"
else
    print_error "Failed to create KIND cluster"
    exit 1
fi

echo
print_step "Step 2: Creating initial test resources"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step creates a namespace and some basic Kubernetes resources."
echo "These resources will be exported by kalco to demonstrate Git version control."
echo

# Create namespace
kubectl create namespace "$NAMESPACE"
print_success "Created namespace: $NAMESPACE"

# Create ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: $NAMESPACE
data:
  environment: "development"
  log-level: "info"
  version: "1.0.0"
EOF
print_success "Created ConfigMap: app-config"

# Create Deployment
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: $NAMESPACE
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.21
        ports:
        - containerPort: 80
EOF
print_success "Created Deployment: nginx-deployment"

# Create Service
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  namespace: $NAMESPACE
spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  type: ClusterIP
EOF
print_success "Created Service: nginx-service"

echo
print_step "Step 3: First export with kalco (creates new Git repo)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step exports all cluster resources to a new directory."
echo "kalco will automatically initialize a Git repository and commit the changes."
echo "This demonstrates the automatic Git setup for new output directories."
echo

echo "Exporting cluster resources to: $OUTPUT_DIR"
echo "Note: This directory doesn't exist yet - kalco will create it automatically"
echo

if ! command -v ./kalco &> /dev/null; then
    print_error "kalco binary not found. Building it first..."
    go build -o kalco .
fi

./kalco --output-dir "$OUTPUT_DIR" --commit-message "Initial cluster snapshot: $(date '+%Y-%m-%d %H:%M:%S')"

if [ $? -eq 0 ]; then
    print_success "First export completed successfully"
else
    print_error "First export failed"
    exit 1
fi

echo
print_step "Step 4: Verifying Git repository and first commit"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step verifies that kalco automatically created a Git repository"
echo "and committed the initial cluster snapshot."
echo

cd "$OUTPUT_DIR"

echo "📊 Git repository status:"
git status --short

echo
echo "📋 Git commit history:"
git log --oneline

echo
echo "📁 Directory structure:"
ls -la

cd ..

echo
print_step "Step 5: Modifying cluster resources"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step modifies the existing resources to demonstrate change detection."
echo "We'll update the ConfigMap and scale the Deployment to show differences."
echo

# Update ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: $NAMESPACE
data:
  environment: "staging"
  log-level: "debug"
  version: "1.1.0"
  feature-flags: "new-feature=true"
EOF
print_success "Updated ConfigMap: app-config"

# Scale Deployment
kubectl scale deployment nginx-deployment --namespace "$NAMESPACE" --replicas=3
print_success "Scaled Deployment: nginx-deployment (1 → 3 replicas)"

# Create new Secret
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
  namespace: $NAMESPACE
type: Opaque
data:
  api-key: YXBpLWtleS1zdGFnaW5n
  password: cGFzc3dvcmQtc3RhZ2luZw==
EOF
print_success "Created Secret: app-secret"

echo
print_step "Step 6: Second export with kalco (updates existing Git repo)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step exports the modified cluster resources to the same directory."
echo "kalco will detect changes and create a new commit in the existing Git repo."
echo "This demonstrates incremental version control and change tracking."
echo

echo "Exporting updated cluster resources to: $OUTPUT_DIR"
echo "Note: This directory already exists with a Git repo - kalco will update it"
echo

./kalco --output-dir "$OUTPUT_DIR" --commit-message "Cluster changes: $(date '+%Y-%m-%d %H:%M:%S')"

if [ $? -eq 0 ]; then
    print_success "Second export completed successfully"
else
    print_error "Second export failed"
    exit 1
fi

echo
print_step "Step 7: Verifying Git history and changes"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step verifies that kalco properly tracked changes in Git."
echo "We can see the complete history of cluster snapshots."
echo

cd "$OUTPUT_DIR"

echo "📊 Git repository status:"
git status --short

echo
echo "📋 Git commit history (showing both snapshots):"
git log --oneline

echo
echo "🔄 Changes between commits:"
git diff HEAD~1 HEAD --name-only

echo
echo "📊 Detailed changes in ConfigMap:"
git diff HEAD~1 HEAD -- "demo-apps/ConfigMap/app-config.yaml" || echo "No ConfigMap changes found"

cd ..

echo
print_step "Step 8: Demonstrating remote origin detection"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step shows how kalco detects remote origins and guides users."
echo "If you have a remote Git repository, you can test the --git-push flag."
echo

cd "$OUTPUT_DIR"

if git remote get-url origin &> /dev/null; then
    print_success "Remote origin detected: $(git remote get-url origin)"
    echo
    echo "💡 To automatically push changes, use:"
    echo "   ./kalco --output-dir $OUTPUT_DIR --git-push"
    echo
    echo "🌐 Current remote branches:"
    git branch -r
else
    print_warning "No remote origin configured"
    echo
    echo "💡 To add a remote origin:"
    echo "   git remote add origin <your-repo-url>"
    echo "   git push -u origin main"
    echo
    echo "💡 Then use kalco with auto-push:"
    echo "   ./kalco --output-dir $OUTPUT_DIR --git-push"
fi

cd ..

echo
print_step "Step 9: Cleanup and summary"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "This step cleans up the test environment and summarizes what we learned."
echo

echo "🧹 Cleaning up KIND cluster..."
kind delete cluster --name "$CLUSTER_NAME"

if [ $? -eq 0 ]; then
    print_success "KIND cluster deleted successfully"
else
    print_warning "Failed to delete KIND cluster - you may need to clean it up manually"
fi

echo
echo "🎉 Git Workflow Test Completed Successfully!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo
echo "📚 What We Demonstrated:"
echo "  ✅ Automatic Git repository initialization for new directories"
echo "  ✅ Seamless integration with existing Git repositories"
echo "  ✅ Change detection and incremental commits"
echo "  ✅ Complete cluster snapshot history"
echo "  ✅ Remote origin detection and guidance"
echo "  ✅ Professional Git workflow integration"
echo
echo "📁 Your cluster history is preserved in: $OUTPUT_DIR"
echo "💡 You can continue using this directory for future exports:"
echo "   ./kalco --output-dir $OUTPUT_DIR --commit-message 'Daily backup'"
echo
echo "🚀 kalco now provides complete Git version control for your Kubernetes clusters!"
