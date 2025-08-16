---
layout: default
title: First Run
parent: Getting Started
nav_order: 2
---

# First Run

Get Kalco running and export your first Kubernetes cluster in minutes.

## Prerequisites

Before running Kalco for the first time, ensure you have:

- ✅ **Kalco installed** - [Installation guide]({{ site.baseurl }}/docs/getting-started/installation)
- ✅ **Kubernetes access** - Valid kubeconfig or in-cluster access
- ✅ **Cluster running** - A Kubernetes cluster to export

## Verify Installation

First, let's confirm Kalco is properly installed:

```bash
# Check version
kalco --version

# Check help
kalco --help

# Verify binary location
which kalco
```

You should see output similar to:
```
Kalco v1.0.0
Kubernetes Analysis & Lifecycle Control
```

## Connect to Your Cluster

Kalco automatically detects your Kubernetes configuration:

```bash
# Check current context
kubectl config current-context

# List available contexts
kubectl config get-contexts

# Switch context if needed
kubectl config use-context your-cluster-name
```

## Your First Export

### Basic Export

Start with a simple export to see Kalco in action:

```bash
# Export all resources to default directory
kalco export
```

This will:
- Create a timestamped output directory
- Export all cluster resources
- Organize them by namespace and type
- Initialize a Git repository (if Git is available)

### Custom Output Directory

Specify where to save your export:

```bash
# Export to specific directory
kalco export --output-dir ./my-cluster-backup

# Export to absolute path
kalco export --output-dir /path/to/backups/production-cluster
```

### Git Integration

Enable version control for your exports:

```bash
# Initialize Git and commit changes
kalco export --git-push

# Custom commit message
kalco export --commit-message "Initial cluster export $(date)"

# Push to remote (if configured)
kalco export --git-push --remote origin
```

## Understanding the Output

After running `kalco export`, you'll see a directory structure like this:

```
my-cluster-backup/
├── _cluster/                    # Cluster-scoped resources
│   ├── ClusterRole/
│   ├── ClusterRoleBinding/
│   ├── CustomResourceDefinition/
│   └── StorageClass/
├── default/                     # Default namespace
│   ├── ConfigMap/
│   ├── Deployment/
│   ├── Service/
│   └── ServiceAccount/
├── kube-system/                 # System namespace
│   ├── ConfigMap/
│   ├── DaemonSet/
│   ├── Deployment/
│   └── Service/
└── kalco-reports/               # Generated reports
    └── Initial-cluster-export.md
```

## What Gets Exported

Kalco automatically discovers and exports:

- **Native Kubernetes Resources** - Pods, Deployments, Services, etc.
- **Custom Resources (CRDs)** - Any CRDs installed in your cluster
- **All Namespaces** - Including system namespaces
- **Resource Metadata** - Labels, annotations, and relationships

## Next Steps

Now that you've exported your first cluster:

1. **[Explore the Output]({{ site.baseurl }}/docs/getting-started/configuration)** - Understand the exported structure
2. **[Validate Resources]({{ site.baseurl }}/docs/commands/validation)** - Check for configuration issues
3. **[Analyze Cluster]({{ site.baseurl }}/docs/commands/analysis)** - Find optimization opportunities
4. **[Generate Reports]({{ site.baseurl }}/docs/commands/reports)** - Create comprehensive documentation

## Common First-Run Scenarios

### Development Cluster

```bash
# Export development cluster
kalco export --output-dir ./dev-cluster --commit-message "Dev cluster snapshot"

# Exclude temporary resources
kalco export --exclude events,pods --output-dir ./dev-cluster-clean
```

### Production Cluster

```bash
# Export production with detailed logging
kalco export --verbose --output-dir ./prod-backup

# Export with Git versioning
kalco export --git-push --commit-message "Production backup $(date)" --output-dir ./prod-backup
```

### Multi-Cluster Setup

```bash
# Export staging cluster
kalco export --context staging --output-dir ./staging-cluster

# Export production cluster
kalco export --context production --output-dir ./production-cluster
```

## Troubleshooting

### Permission Issues

```bash
# Check cluster access
kubectl get nodes

# Verify kubeconfig
kubectl config view
```

### Resource Export Issues

```bash
# Enable verbose output
kalco export --verbose

# Check specific resource types
kubectl api-resources
```

### Git Issues

```bash
# Check Git status
git status

# Initialize Git manually
git init
git add .
git commit -m "Initial commit"
```

## Success Indicators

You'll know your first run was successful when you see:

- ✅ **Export completed** - No error messages
- ✅ **Output directory created** - With organized resource files
- ✅ **Git repository initialized** - If Git integration enabled
- ✅ **Resources exported** - YAML files for each resource
- ✅ **Report generated** - Summary of the export process

Congratulations! You've successfully exported your first Kubernetes cluster with Kalco. 🎉
