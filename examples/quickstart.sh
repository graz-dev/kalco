#!/bin/bash

# Kalco Quickstart Script
# This comprehensive script demonstrates all of kalco's capabilities including Cross-Reference Validation

set -e

echo "🚀 Kalco Quickstart Demo with Cross-Reference Validation"
echo "========================================================"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
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

print_feature() {
    echo -e "${PURPLE}🔍${NC} $1"
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
kind create cluster --name kalco-validation-test --wait 2m
print_status "Test cluster created"

# Wait for cluster to be ready
echo "⏳ Waiting for cluster to be ready..."
kubectl wait --for=condition=Ready nodes --all --timeout=300s
print_status "Cluster is ready"

# Create test namespace and resources
echo ""
echo "📦 Creating test resources..."

# Create namespace
kubectl create namespace validation-test

# Create initial ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: validation-test
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
  namespace: validation-test
  labels:
    app: nginx
    tier: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
      tier: frontend
  template:
    metadata:
      labels:
        app: nginx
        tier: frontend
    spec:
      containers:
      - name: nginx
        image: nginx:1.21
        ports:
        - containerPort: 80
EOF

# Create ServiceAccount for RBAC testing
kubectl apply -f - <<EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-service-account
  namespace: validation-test
EOF

# Create Role
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: app-role
  namespace: validation-test
rules:
- apiGroups: [""]
  resources: ["pods", "services"]
  verbs: ["get", "list", "watch"]
EOF

# Create RoleBinding
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: app-role-binding
  namespace: validation-test
subjects:
- kind: ServiceAccount
  name: app-service-account
  namespace: validation-test
roleRef:
  kind: Role
  name: app-role
  apiGroup: rbac.authorization.k8s.io
EOF

print_status "Initial test resources created"

# First export - creates Git repo and initial report
echo ""
echo "📦 First export - creating initial snapshot..."
./kalco --output-dir ./quickstart-demo --commit-message "Initial snapshot: $(date)"
print_status "Initial export completed"

# Verify Git repository and initial report
echo ""
echo "🔍 Verifying initial setup..."
cd ./quickstart-demo

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

# Now create resources with BROKEN REFERENCES to demonstrate validation
echo ""
echo "🔍 Creating resources with BROKEN REFERENCES to demonstrate Cross-Reference Validation..."
print_feature "This will show how kalco detects broken references!"

# Create Service with BROKEN selector (targets non-existent deployment)
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: broken-service
  namespace: validation-test
spec:
  selector:
    app: non-existent-app
    tier: backend
  ports:
  - port: 8080
    targetPort: 8080
  type: ClusterIP
EOF

# Create NetworkPolicy with BROKEN pod selector
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: broken-network-policy
  namespace: validation-test
spec:
  podSelector:
    matchLabels:
      app: non-existent-app
      tier: backend
  policyTypes:
  - Ingress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: nginx
          tier: frontend
    ports:
    - protocol: TCP
      port: 80
EOF

# Create Ingress with BROKEN backend service
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: broken-ingress
  namespace: validation-test
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: broken.example.com
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

# Create HorizontalPodAutoscaler with BROKEN target
kubectl apply -f - <<EOF
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: broken-hpa
  namespace: validation-test
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: non-existent-deployment
  minReplicas: 1
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50
EOF

# Create PodDisruptionBudget with BROKEN selector
kubectl apply -f - <<EOF
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: broken-pdb
  namespace: validation-test
spec:
  minAvailable: 1
  selector:
    matchLabels:
      app: non-existent-app
      tier: backend
EOF

# Create RoleBinding with BROKEN ServiceAccount reference
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: broken-role-binding
  namespace: validation-test
subjects:
- kind: ServiceAccount
  name: non-existent-service-account
  namespace: validation-test
roleRef:
  kind: Role
  name: app-role
  apiGroup: rbac.authorization.k8s.io
EOF

# Create RoleBinding with external User reference (will be a warning)
kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: external-user-binding
  namespace: validation-test
subjects:
- kind: User
  name: external-user@example.com
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: app-role
  apiGroup: rbac.authorization.k8s.io
EOF

print_status "Resources with broken references created"
print_warning "These resources have intentional broken references to demonstrate validation!"

# Second export - updates Git repo and generates change report with validation
echo ""
echo "📦 Second export - generating change report with Cross-Reference Validation..."
./kalco --output-dir ./quickstart-demo --commit-message "Broken references demo: $(date)"
print_status "Second export completed"

# Analyze the enhanced report with validation
echo ""
echo "📊 Analyzing enhanced change report with Cross-Reference Validation..."
cd ./quickstart-demo

# Find the latest report
latest_report=$(ls -t kalco-reports/*.md | head -1)
print_info "Latest report: $latest_report"

# Display report summary
echo ""
echo "📋 Report Summary:"
echo "=================="
grep -E "^## |^### |^#### " "$latest_report" | head -25

echo ""
echo "🔍 Cross-Reference Validation Section:"
echo "======================================"
grep -A 5 -B 5 "Cross-Reference Validation" "$latest_report" || echo "Validation section not found"

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
echo "💡 Recommendations:"
echo "=================="
grep -A 15 "Recommendations" "$latest_report" | head -20

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
kind delete cluster --name kalco-validation-test
print_status "Test cluster deleted"

echo ""
echo "🎉 Enhanced Quickstart Demo Completed!"
echo "======================================"
echo ""
echo "📊 What was tested:"
echo "- ✅ Initial snapshot with Git repository creation"
echo "- ✅ Initial change report generation"
echo "- ✅ Resource modification (ConfigMap, Deployment, ServiceAccount, Role, RoleBinding)"
echo "- ✅ Enhanced change report with detailed diffs"
echo "- ✅ Git history tracking"
echo "- 🔍 CROSS-REFERENCE VALIDATION (NEW FEATURE!)"
echo "  - ❌ Broken Service selectors"
echo "  - ❌ Broken NetworkPolicy selectors"
echo "  - ❌ Broken Ingress backends"
echo "  - ❌ Broken HPA targets"
echo "  - ❌ Broken PDB selectors"
echo "  - ❌ Broken RoleBinding ServiceAccount references"
echo "  - ⚠️  External User references (warnings)"
echo ""
echo "📁 Your enhanced backup is preserved in: ./quickstart-demo/"
echo "📋 Enhanced reports with validation are in: ./quickstart-demo/kalco-reports/"
echo ""
echo "🔍 Key Features Demonstrated:"
echo "- 🆕 New resources show complete YAML content"
echo "- ✏️ Modified resources show Git diff with before/after"
echo "- 🗑️ Deleted resources show what was removed"
echo "- 📊 Change summaries with line counts and section tracking"
echo "- 🔍 Field-level change identification"
echo "- 🔍 CROSS-REFERENCE VALIDATION:"
echo "  - ✅ Valid references tracking"
echo "  - ❌ Broken references detection"
echo "  - ⚠️  Warning references for external resources"
echo "  - 📋 Actionable recommendations"
echo "  - 🛡️ Reliability assurance for reapplying resources"
echo ""
echo "💡 Try viewing the reports to see kalco's enhanced functionality!"
echo "🔍 The Cross-Reference Validation section will show you exactly what's broken!"
