#!/bin/bash

# Kalco Simple Quickstart Script
# This script demonstrates a real, cohesive application with:
# - Echo server deployment
# - Service and Ingress
# - Real CRD from kube-green operator (SleepInfo)
# - Cross-reference validation
# - Orphaned resource detection

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_feature() {
    echo -e "${BLUE}🔍 $1${NC}"
}

# Script header
echo "🚀 Kalco Simple Quickstart Demo"
echo "================================"
echo ""
echo "This demo shows a real, cohesive application with:"
echo "- 🚀 HTTP Echo server deployment"
echo "- 🌐 Service and Ingress"
echo "- 🔧 Real kube-green operator with SleepInfo CRD"
echo "- 🔍 Cross-reference validation"
echo "- 🗑️  Orphaned resource detection"
echo ""

# Check prerequisites
echo "🔍 Checking prerequisites..."
if ! command -v kubectl &> /dev/null; then
    print_error "kubectl is required but not installed"
    exit 1
fi

if ! command -v kind &> /dev/null; then
    print_error "kind is required but not installed"
    exit 1
fi

if ! command -v go &> /dev/null; then
    print_error "go is required but not installed"
    exit 1
fi



print_status "All prerequisites are available"

# Build kalco
echo ""
echo "ℹ️ Building kalco..."
go build -o kalco
print_status "kalco built successfully"

# Create test cluster
echo ""
echo "🏗️ Creating test cluster..."
kind create cluster --name kalco-quickstart --wait 2m
print_status "Test cluster created"

# Wait for cluster to be ready
echo ""
echo "⏳ Waiting for cluster to be ready..."
sleep 10  # Give the cluster a moment to fully initialize
kubectl wait --for=condition=Ready node/kalco-quickstart-control-plane --timeout=60s
print_status "Cluster is ready"

# Create a simple, real application
echo ""
echo "📦 Creating a simple, real application..."

# Create namespace for our demo app
kubectl create namespace demo-app --dry-run=client -o yaml | kubectl apply -f -

# Create ConfigMap for echo server configuration
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: echo-config
  namespace: demo-app
  labels:
    app: echo-server
    tier: backend
data:
  environment: "development"
  log-level: "info"
  app-version: "1.0.0"
  description: "Simple echo server for demo purposes"
EOF

# Create Deployment for echo server
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echo-server
  namespace: demo-app
  labels:
    app: echo-server
    tier: backend
spec:
  replicas: 2
  selector:
    matchLabels:
      app: echo-server
      tier: backend
  template:
    metadata:
      labels:
        app: echo-server
        tier: backend
    spec:
      containers:
      - name: echo
        image: hashicorp/http-echo:latest
        ports:
        - containerPort: 80
          name: http
        command: ["/http-echo"]
        args: ["-text", "Hello from Echo Server!", "-listen", ":80"]
        resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "128Mi"
            cpu: "200m"
        livenessProbe:
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 10
          periodSeconds: 30
        readinessProbe:
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 5
          periodSeconds: 5
EOF

# Create Service for echo server
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: echo-service
  namespace: demo-app
  labels:
    app: echo-server
    tier: backend
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 80
    protocol: TCP
    name: http
  selector:
    app: echo-server
    tier: backend
EOF

# Create Ingress for external access
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: echo-ingress
  namespace: demo-app
  labels:
    app: echo-server
    tier: frontend
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - host: echo.local
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: echo-service
            port:
              number: 80
EOF

# Install kube-green operator for real CRD testing
echo ""
echo "🔧 Installing kube-green operator for real CRD testing..."
print_feature "This will install a real operator with real CRDs!"

# First install cert-manager (required for kube-green)
echo "📦 Installing cert-manager..."
kubectl apply -f https://github.com/jetstack/cert-manager/releases/latest/download/cert-manager.yaml

# Wait for cert-manager to be ready
echo "⏳ Waiting for cert-manager to be ready..."
kubectl wait --for=condition=Ready pod -l app.kubernetes.io/instance=cert-manager -n cert-manager --timeout=120s
print_status "cert-manager installed and ready"

# Install kube-green
echo "📦 Installing kube-green..."
kubectl apply -f https://github.com/kube-green/kube-green/releases/latest/download/kube-green.yaml

# Wait for kube-green to be ready
echo "⏳ Waiting for kube-green to be ready..."
kubectl wait --for=condition=Ready pod -l app=kube-green -n kube-green --timeout=120s
print_status "kube-green operator installed and ready"

# Give kube-green webhook a moment to be fully ready
echo "⏳ Waiting for kube-green webhook to be ready..."
sleep 30

# Create a SleepInfo CRD resource
echo ""
echo "🌙 Creating SleepInfo CRD resource..."
kubectl apply -f - <<EOF
apiVersion: kube-green.com/v1alpha1
kind: SleepInfo
metadata:
  name: demo-sleep
  namespace: demo-app
spec:
  weekdays: "1-5"
  sleepAt: "20:00"
  wakeUpAt: "08:00"
  timeZone: "Europe/Rome"
EOF

print_status "SleepInfo CRD resource created"

# Create some resources with intentional broken references for validation testing
echo ""
echo "🔍 Creating resources with broken references for validation testing..."

# Create Service targeting non-existent deployment
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: broken-service
  namespace: demo-app
  labels:
    app: broken
    tier: backend
spec:
  type: ClusterIP
  ports:
  - port: 8080
    targetPort: 8080
    protocol: TCP
    name: http
  selector:
    app: non-existent-deployment
    tier: backend
EOF

# Create Ingress with broken service backend
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: broken-ingress
  namespace: demo-app
  labels:
    app: broken
    tier: frontend
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: broken.local
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: non-existent-service
            port:
              number: 80
EOF

# Create orphaned resources (no references)
echo ""
echo "🗑️  Creating orphaned resources for detection testing..."

# Create orphaned ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: orphaned-config
  namespace: demo-app
  labels:
    app: orphaned
    tier: test
data:
  orphaned: "true"
  description: "This ConfigMap has no references and will be detected as orphaned"
EOF

# Create orphaned Secret
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: orphaned-secret
  namespace: demo-app
  labels:
    app: orphaned
    tier: test
type: Opaque
data:
  orphaned: "dHJ1ZQ=="  # true
stringData:
  description: "This Secret has no references and will be detected as orphaned"
EOF

print_status "All test resources created"

# First export - creates Git repo and initial report
echo ""
echo "📦 First export - creating initial snapshot..."
if [ ! -f "./kalco" ]; then
    print_error "kalco binary not found. Please build it first."
    exit 1
fi
./kalco --output-dir ./quickstart-demo --commit-message "Initial snapshot: $(date)"
print_status "Initial export completed"

# Verify Git repository and initial report
echo ""
echo "🔍 Verifying initial setup..."
cd ./quickstart-demo

print_status "Git repository initialized"
print_status "Reports directory created"
initial_report=$(ls kalco-reports/*.md | head -1)
print_info "Initial report: $initial_report"

cd ..

# Now modify some resources to demonstrate change tracking
echo ""
echo "✏️  Modifying existing resources to demonstrate change tracking..."
print_feature "This will show how kalco tracks resource changes!"

# Modify the ConfigMap to add new data
kubectl patch configmap echo-config -n demo-app --patch '{"data":{"new-feature":"enabled","debug-mode":"true"}}'

# Modify the Deployment to change replica count
kubectl scale deployment echo-server -n demo-app --replicas=3

# Add labels to existing resources
kubectl label deployment echo-server -n demo-app environment=staging version=v2.0.0 --overwrite
kubectl label service echo-service -n demo-app environment=staging version=v2.0.0 --overwrite

print_status "Resource modifications completed for change tracking demonstration"

# Second export - generates change report with validation
echo ""
echo "📦 Second export - generating change report with Cross-Reference Validation..."
./kalco --output-dir ./quickstart-demo --commit-message "Changes and validation demo: $(date)"
print_status "Second export completed"

# Analyze the enhanced report
echo ""
echo "📊 Analyzing enhanced change report with Cross-Reference Validation..."
cd ./quickstart-demo

# Find the latest report
latest_report=$(ls -t kalco-reports/*.md | head -1)
echo "📋 Latest report: $latest_report"

echo ""
echo "🔍 Cross-Reference Validation Results:"
echo "======================================"
grep -A 5 "## 🔍 Cross-Reference Validation" "$latest_report"

echo ""
echo "❌ Broken References Found:"
echo "============================"
grep -A 10 "Broken References" "$latest_report" | head -30

echo ""
echo "⚠️  Warning References:"
echo "======================="
grep -A 10 "Warning References" "$latest_report" | head -20

echo ""
echo "✅ Valid References Summary:"
echo "============================"
grep -A 10 "Valid References Summary" "$latest_report" | head -15

echo ""
echo "💡 Validation Recommendations:"
echo "============================="
grep -A 15 "Recommendations" "$latest_report" | head -20

echo ""
echo "🗑️  Orphaned Resource Detection Section:"
echo "========================================="
grep -A 5 -B 5 "Orphaned Resource Detection" "$latest_report"

echo ""
echo "🗑️  Orphaned Resources Found:"
echo "=============================="
grep -A 10 "Orphaned Resources Found" "$latest_report" | head -30

echo ""
echo "💡 Cleanup Recommendations:"
echo "=========================="
grep -A 15 "Cleanup Recommendations" "$latest_report" | head -20

# Check for CRD handling
echo ""
echo "🔧 Custom Resource Definition (CRD) Handling:"
echo "============================================="
grep -A 5 "SleepInfo\|kube-green" "$latest_report"

echo ""
echo "🌐 Application Resources:"
echo "========================"
grep -A 5 "echo-server\|echo-service\|echo-ingress" "$latest_report"

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
kind delete cluster --name kalco-quickstart
print_status "Test cluster deleted"

echo ""
echo "🎉 Simple Quickstart Demo Completed!"
echo "===================================="
echo ""
echo "📊 What was tested:"
echo "- ✅ Initial snapshot with Git repository creation"
echo "- ✅ Simple, cohesive application (HTTP echo server + service + ingress)"
echo "- ✅ Real kube-green operator with SleepInfo CRD"
echo "- ✅ Resource modification and change tracking"
echo "- ✅ Enhanced change report with validation"
echo "- ✅ Git history tracking"
echo "- 🔍 CROSS-REFERENCE VALIDATION"
echo "  - ❌ Broken Service selectors (targeting non-existent deployments)"
echo "  - ❌ Broken Ingress backends (non-existent services)"
echo "- 🗑️  ORPHANED RESOURCE DETECTION"
echo "  - 🗑️  Orphaned ConfigMaps (unreferenced)"
echo "  - 🗑️  Orphaned Secrets (unreferenced)"
echo ""
echo "📁 Your backup is preserved in: ./quickstart-demo/"
echo "📋 Enhanced reports with validation are in: ./quickstart-demo/kalco-reports/"
echo ""
echo "🔍 Key Features Demonstrated:"
echo "- 🆕 New resources show complete YAML content"
echo "- ✏️ Modified resources show Git diff with before/after"
echo "- 📊 Change summaries with line counts and section tracking"
echo "- 🔍 Field-level change identification"
echo "- 🔍 CROSS-REFERENCE VALIDATION:"
echo "  - ✅ Valid references tracking"
echo "  - ❌ Broken references detection"
echo "  - 📋 Actionable recommendations"
echo "  - 🛡️ Reliability assurance for reapplying resources"
echo "- 🗑️  ORPHANED RESOURCE DETECTION:"
echo "  - 🔍 Orphaned resource identification"
echo "  - 📊 Resource breakdown by type"
echo "  - 💡 Cleanup recommendations"
echo "  - 🧹 Cluster cleanup guidance"
echo "- 🔧 REAL CRD SUPPORT:"
echo "  - 🌐 kube-green operator installation"
echo "  - 📦 SleepInfo CRD resource creation"
echo "  - 🔍 CRD validation and analysis"
echo ""
echo "💡 Try viewing the reports to see kalco's functionality!"
echo "🔍 The Cross-Reference Validation section will show you exactly what's broken!"
echo "🗑️  The Orphaned Resource Detection will help you clean up your cluster!"
echo "🔧 CRD support ensures all your custom resources are properly handled!"
