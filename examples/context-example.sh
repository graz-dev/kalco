#!/bin/bash

# Kalco Context Management Example
# This script demonstrates how to use Kalco contexts for managing multiple clusters

set -e

echo "🚀 Kalco Context Management Example"
echo "=================================="
echo

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
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

# Check if kalco is available (use local binary if available, otherwise system)
KALCO_CMD="kalco"
if [ -f "./kalco" ]; then
    KALCO_CMD="./kalco"
    print_info "Using local kalco binary"
elif command -v kalco &> /dev/null; then
    print_info "Using system kalco binary"
else
    print_error "Kalco is not available. Please build it first with 'make build' or install it."
    exit 1
fi

print_success "Kalco is available and ready to use"
echo

# Create example contexts for different environments
print_info "Creating example contexts for different environments..."

# Production context
print_info "Setting up production context..."
$KALCO_CMD context set production \
  --kubeconfig ~/.kube/config \
  --output ./exports/production \
  --description "Production cluster for live workloads" \
  --labels env=production \
  --labels team=platform \
  --labels region=eu-west

# Staging context
print_info "Setting up staging context..."
$KALCO_CMD context set staging \
  --kubeconfig ~/.kube/config \
  --output ./exports/staging \
  --description "Staging cluster for testing" \
  --labels env=staging \
  --labels team=qa \
  --labels purpose=testing

# Development context
print_info "Setting up development context..."
$KALCO_CMD context set development \
  --kubeconfig ~/.kube/config \
  --output ./exports/development \
  --description "Development cluster for engineers" \
  --labels env=development \
  --labels team=engineering \
  --labels purpose=development

print_success "All contexts created successfully"
echo

# List all contexts
print_info "Listing all available contexts..."
$KALCO_CMD context list
echo

# Switch between contexts and show current
print_info "Demonstrating context switching..."

print_info "Switching to production context..."
$KALCO_CMD context use production
$KALCO_CMD context current
echo

print_info "Switching to staging context..."
$KALCO_CMD context use staging
$KALCO_CMD context current
echo

print_info "Switching to development context..."
$KALCO_CMD context use development
$KALCO_CMD context current
echo

# Show how export uses context automatically
print_info "Demonstrating automatic context usage in export command..."
print_info "Current context: development"
$KALCO_CMD export --dry-run
echo

# Switch back to production and export
print_info "Switching to production and exporting..."
$KALCO_CMD context use production
$KALCO_CMD export --dry-run
echo

# Show context override with flags
print_info "Demonstrating context override with command-line flags..."
$KALCO_CMD export --dry-run --output ./override-output
echo

# Show context details
print_info "Showing detailed context information..."
$KALCO_CMD context show production
echo

# Cleanup example contexts
print_warning "Cleaning up example contexts..."
$KALCO_CMD context use development
$KALCO_CMD context delete production
$KALCO_CMD context use staging
$KALCO_CMD context delete development
$KALCO_CMD context use temp
$KALCO_CMD context delete staging

print_success "Example contexts cleaned up"
echo

print_success "Context management example completed successfully!"
echo
print_info "Key takeaways:"
echo "  • Use 'kalco context set' to create contexts"
echo "  • Use 'kalco context use' to switch between contexts"
echo "  • Contexts automatically configure kubeconfig and output directory"
echo "  • Command-line flags can override context settings"
echo "  • Use 'kalco context list' to see all available contexts"
echo "  • Use 'kalco context current' to see active context"
echo
print_info "For more information, run: kalco context --help"
