#!/bin/bash

# Kalco Quickstart Script
# This comprehensive script demonstrates all of kalco's capabilities

set -e

echo "🚀 Kalco Quickstart Demo"
echo "========================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✅${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ️${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

print_error() {
    echo -e "${RED}❌${NC} $1"
}

# Check prerequisites
echo "🔍 Checking prerequisites..."

if ! command -v kind &> /dev/null; then
    print_error "KIND is not installed. Please install it first: https://kind.sigs.k8s.io/docs/user/quick-start/"
    exit 1
fi

if ! command -v kubectl &> /dev/null; then
    print_error "kubectl is not installed. Please install it first."
    exit 1
fi

if ! command -v git &> /dev/null; then
    print_error "git is not installed. Please install it first."
    exit 1
fi

print_status "All prerequisites are available"

# Build kalco if not exists
if [ ! -f "./kalco" ]; then
    print_info "Building kalco..."
    go build -o kalco
    print_status "kalco built successfully"
else
    print_status "kalco binary found"
fi

# Create test cluster
echo ""
echo "🏗️ Creating test cluster..."
kind create cluster --name kalco-enhanced-reports-test --wait 2m
print_status "Test cluster created"

# Wait for cluster to be ready
echo "⏳ Waiting for cluster to be ready..."
kubectl wait --for=condition=Ready nodes --all --timeout=300s
print_status "Cluster is ready"

# Create test namespace and resources
echo ""
echo "📦 Creating test resources..."

# Create namespace
kubectl create namespace enhanced-test

# Create initial ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: enhanced-test
data:
  environment: "development"
  log-level: "info"
  version: "1.0.0"
EOF

# Create initial Deployment
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: enhanced-test
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

print_status "Initial test resources created"

# First export - creates Git repo and initial report
echo ""
echo "📦 First export - creating initial snapshot..."
./kalco --output-dir ./enhanced-test-backup --commit-message "Initial snapshot: $(date)"
print_status "Initial export completed"

# Verify Git repository and initial report
echo ""
echo "🔍 Verifying initial setup..."
cd ./enhanced-test-backup

if [ -d ".git" ]; then
    print_status "Git repository initialized"
else
    print_error "Git repository not found"
    exit 1
fi

if [ -d "kalco-reports" ]; then
    print_status "Reports directory created"
    initial_report=$(ls kalco-reports/*.md | head -1)
    print_info "Initial report: $initial_report"
else
    print_error "Reports directory not found"
    exit 1
fi

cd ..

# Modify cluster resources to test change tracking
echo ""
echo "🔄 Modifying cluster resources..."

# Update ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: enhanced-test
data:
  environment: "staging"
  log-level: "debug"
  version: "1.1.0"
  feature-flags: "new-feature=true"
  database-url: "postgres://staging:5432"
EOF

# Scale deployment
kubectl scale deployment nginx-deployment --namespace enhanced-test --replicas=3

# Create new Secret
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
  namespace: enhanced-test
type: Opaque
data:
  api-key: YXBpLWtleS1zdGFnaW5n
  password: cGFzc3dvcmQtc3RhZ2luZw==
EOF

# Create new Service
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  namespace: enhanced-test
spec:
  selector:
    app: nginx
  ports:
  - port: 80
    targetPort: 80
  type: ClusterIP
EOF

print_status "Resources modified"

# Second export - updates Git repo and generates change report
echo ""
echo "📦 Second export - generating change report..."
./kalco --output-dir ./enhanced-test-backup --commit-message "Enhanced resources: $(date)"
print_status "Second export completed"

# Analyze the enhanced report
echo ""
echo "📊 Analyzing enhanced change report..."
cd ./enhanced-test-backup

# Find the latest report
latest_report=$(ls -t kalco-reports/*.md | head -1)
print_info "Latest report: $latest_report"

# Display report summary
echo ""
echo "📋 Report Summary:"
echo "=================="
grep -E "^## |^### |^#### " "$latest_report" | head -20

echo ""
echo "🔍 Detailed Changes Section:"
echo "============================"
grep -A 5 -B 5 "Detailed Resource Changes" "$latest_report" || echo "Detailed changes section not found"

echo ""
echo "📊 Change Details for Modified Resources:"
echo "========================================="
grep -A 10 "Resource Modified" "$latest_report" | head -20

echo ""
echo "💻 Git History:"
echo "==============="
git log --oneline -3

echo ""
echo "🔄 What Changed:"
echo "================"
git diff HEAD~1 HEAD --name-status

cd ..

# Cleanup
echo ""
echo "🧹 Cleaning up..."
kind delete cluster --name kalco-enhanced-reports-test
print_status "Test cluster deleted"

echo ""
echo "🎉 Quickstart Demo Completed!"
echo "============================="
echo ""
echo "📊 What was tested:"
echo "- ✅ Initial snapshot with Git repository creation"
echo "- ✅ Initial change report generation"
echo "- ✅ Resource modification (ConfigMap, Deployment, Secret, Service)"
echo "- ✅ Enhanced change report with detailed diffs"
echo "- ✅ Git history tracking"
echo ""
echo "📁 Your enhanced backup is preserved in: ./enhanced-test-backup/"
echo "📋 Enhanced reports are in: ./enhanced-test-backup/kalco-reports/"
echo ""
echo "🔍 Key Features Demonstrated:"
echo "- 🆕 New resources show complete YAML content"
echo "- ✏️ Modified resources show Git diff with before/after"
echo "- 🗑️ Deleted resources show what was removed"
echo "- 📊 Change summaries with line counts and section tracking"
echo "- 🔍 Field-level change identification"
echo ""
echo "💡 Try viewing the reports to see kalco's enhanced functionality!"
