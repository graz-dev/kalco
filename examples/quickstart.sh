#!/bin/bash

# Kalco Quickstart Script
# This script demonstrates the core functionality of Kalco

set -e

echo "Kalco Quickstart - Kubernetes Analysis & Lifecycle Control"
echo "=========================================================="
echo

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

# Check if kalco is installed
if ! command -v kalco &> /dev/null; then
    print_error "Kalco is not installed. Please install it first."
    print_info "Run: curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash"
    exit 1
fi

print_success "Kalco is installed: $(kalco version)"

# Create a temporary directory for this demo
DEMO_DIR=$(mktemp -d)
print_status "Created demo directory: $DEMO_DIR"
cd "$DEMO_DIR"

# Create a test context
print_status "Creating test context..."
kalco context set demo \
    --kubeconfig ~/.kube/config \
    --output ./demo-exports \
    --description "Demo context for quickstart" \
    --labels env=demo,team=platform

print_success "Context 'demo' created successfully"

# List contexts
print_status "Listing available contexts..."
kalco context list

# Use the demo context
print_status "Switching to demo context..."
kalco context use demo

print_success "Now using context 'demo'"

# Show current context
print_status "Current context details:"
kalco context current

# Create a simple test file to simulate cluster resources
print_status "Creating test cluster resources..."
mkdir -p demo-exports/default/pods
mkdir -p demo-exports/default/services
mkdir -p demo-exports/_cluster/namespaces

cat > demo-exports/default/pods/nginx.yaml << 'EOF'
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  namespace: default
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
EOF

cat > demo-exports/default/services/nginx-service.yaml << 'EOF'
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  namespace: default
spec:
  selector:
    app: nginx
  ports:
  - port: 80
    targetPort: 80
  type: ClusterIP
EOF

cat > demo-exports/_cluster/namespaces/default.yaml << 'EOF'
apiVersion: v1
kind: Namespace
metadata:
  name: default
EOF

print_success "Test resources created"

# Initialize Git repository manually to simulate the export process
print_status "Initializing Git repository..."
cd demo-exports
git init
git add .
git commit -m "Initial cluster snapshot: $(date)"

print_success "Git repository initialized with initial snapshot"

# Modify a resource to simulate changes
print_status "Simulating resource changes..."
sed -i 's/nginx:latest/nginx:1.21/g' default/pods/nginx.yaml

# Add the change
git add .
git commit -m "Updated nginx version: $(date)"

print_success "Resource modification committed"

# Go back to demo directory
cd ..

# Now demonstrate the export command (which would normally connect to a real cluster)
print_status "Demonstrating export command structure..."
print_info "Note: This is a simulation. In a real environment, kalco export would:"
print_info "1. Connect to your Kubernetes cluster"
print_info "2. Discover all available resources"
print_info "3. Export them to organized YAML files"
print_info "4. Initialize Git repository if needed"
print_info "5. Commit changes with timestamp"
print_info "6. Generate comprehensive change report"

# Show the directory structure
print_status "Current export directory structure:"
tree demo-exports || find demo-exports -type f

# Show the Git log
print_status "Git commit history:"
cd demo-exports
git log --oneline

# Show the kalco-config.json
print_status "Kalco configuration:"
if [ -f kalco-config.json ]; then
    cat kalco-config.json | jq . 2>/dev/null || cat kalco-config.json
else
    print_warning "kalco-config.json not found (this would be created by kalco export)"
fi

cd ..

# Cleanup
print_status "Cleaning up demo environment..."
if [ "$1" != "--keep" ]; then
    rm -rf "$DEMO_DIR"
    print_success "Demo directory cleaned up"
else
    print_info "Demo directory kept at: $DEMO_DIR"
    print_info "Run: rm -rf $DEMO_DIR to clean up manually"
fi

echo
print_success "Quickstart completed successfully!"
echo
print_info "What you learned:"
print_info "• How to create and manage contexts with kalco context"
print_info "• How contexts store cluster configuration and output directories"
print_info "• How kalco organizes exported resources in a structured way"
print_info "• How Git integration works for version control"
print_info "• How kalco-config.json stores context information"
echo
print_info "Next steps:"
print_info "• Try kalco export with a real Kubernetes cluster"
print_info "• Explore the generated reports in kalco-reports/ directory"
print_info "• Use kalco context load to import existing exports"
print_info "• Read the documentation: https://graz-dev.github.io/kalco"
echo
print_info "For more examples and documentation, visit:"
print_info "https://github.com/graz-dev/kalco"
