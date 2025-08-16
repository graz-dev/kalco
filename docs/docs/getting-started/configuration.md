---
layout: default
title: Configuration
nav_order: 3
parent: Getting Started
---

# Configuration

Customize Kalco to fit your workflow and requirements.

## ⚙️ Configuration Methods

Kalco supports multiple configuration methods in order of precedence:

1. **Command-line flags** (highest priority)
2. **Environment variables**
3. **Configuration files**
4. **Default values** (lowest priority)

## 🚩 Command-Line Options

### Basic Options

```bash
# Output directory
kalco export --output ./cluster-backup

# Specific namespaces
kalco export --namespaces default,kube-system,production

# Resource filtering
kalco export --resources pods,services,deployments
kalco export --exclude events,replicasets,endpoints

# Verbose output
kalco export --verbose
```

### Git Integration

```bash
# Enable Git operations
kalco export --git-push

# Custom commit message
kalco export --commit-message "Weekly backup - $(date)"

# Dry run (no actual export)
kalco export --dry-run
```

## 🌍 Environment Variables

Set these environment variables for persistent configuration:

```bash
# Output directory
export KALCO_OUTPUT_DIR="./cluster-exports"

# Default namespaces
export KALCO_NAMESPACES="default,kube-system,production"

# Git integration
export KALCO_GIT_PUSH="true"
export KALCO_COMMIT_MESSAGE="Cluster snapshot"

# Resource filtering
export KALCO_RESOURCES="pods,services,deployments,configmaps"
export KALCO_EXCLUDE="events,replicasets,endpoints"
```

## 📁 Configuration Files

Create a configuration file for complex setups:

```yaml
# ~/.kalco/config.yaml
output:
  directory: "./cluster-backups"
  format: "yaml"
  compress: false

filtering:
  namespaces:
    - "default"
    - "kube-system"
    - "production"
  resources:
    include:
      - "pods"
      - "services"
      - "deployments"
      - "configmaps"
    exclude:
      - "events"
      - "replicasets"
      - "endpoints"

git:
  enabled: true
  auto_push: false
  commit_message: "Cluster export - {timestamp}"
  remote_origin: "origin"

validation:
  cross_references: true
  orphaned_resources: true
  detailed_reporting: true

output:
  reports:
    enabled: true
    format: "markdown"
    include_changes: true
    include_validation: true
```

## 🔧 Advanced Configuration

### Custom Resource Types

```bash
# Include Custom Resource Definitions
kalco export --resources pods,services,mycustomresource

# Exclude specific CRDs
kalco export --exclude events,replicasets,mycustomresource
```

### Output Formatting

```bash
# Compressed output
kalco export --compress

# Custom file naming
kalco export --output "./backups/cluster-{date}-{time}"
```

### Validation Options

```bash
# Skip validation for faster export
kalco export --skip-validation

# Custom validation rules
kalco export --validation-rules ./rules.yaml
```

## 📋 Configuration Examples

### Development Environment

```bash
# Quick export for development
kalco export \
  --output "./dev-cluster" \
  --namespaces "default,development" \
  --exclude "events,replicasets" \
  --verbose
```

### Production Backup

```bash
# Comprehensive production backup
kalco export \
  --output "./production-backup-$(date +%Y%m%d)" \
  --namespaces "production,monitoring,security" \
  --git-push \
  --commit-message "Production backup - $(date)"
```

### CI/CD Pipeline

```bash
# Automated export in CI/CD
kalco export \
  --output "./cluster-state" \
  --git-push \
  --commit-message "Automated backup - Build ${BUILD_NUMBER}"
```

## 🔍 Configuration Validation

Validate your configuration:

```bash
# Check configuration
kalco config validate

# Show current configuration
kalco config show

# Test configuration
kalco export --dry-run --verbose
```

## 📚 Next Steps

1. **[Commands Reference]({{ site.baseurl }}/docs/commands/)** - Complete command documentation
2. **[Use Cases]({{ site.baseurl }}/docs/use-cases/)** - Common scenarios and workflows
3. **[Troubleshooting]({{ site.baseurl }}/docs/getting-started/troubleshooting)** - Solve common issues
