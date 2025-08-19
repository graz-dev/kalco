---
layout: default
title: First Run
nav_order: 2
parent: Getting Started
---

# First Run

This guide walks you through your first Kalco export, from setting up a context to generating your first cluster snapshot.

## Prerequisites

Before starting, ensure you have:

- **Kalco installed** - See [Installation](installation.md) if needed
- **Kubernetes access** - Valid kubeconfig or in-cluster access
- **Git installed** (optional) - For version control functionality

## Quick Start

### 1. Verify Installation

First, confirm Kalco is working:

```bash
kalco version
```

You should see version information and build details.

### 2. Check Available Commands

Explore what Kalco can do:

```bash
kalco --help
```

This shows the available commands: `context`, `export`, `completion`, and `version`.

### 3. Set Up Your First Context

Create a context for your cluster:

```bash
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-cluster-exports \
  --description "My first Kubernetes cluster" \
  --labels env=dev,team=personal
```

### 4. Use the Context

Activate your context:

```bash
kalco context use my-cluster
```

### 5. Export Your Cluster

Perform your first export:

```bash
kalco export --git-push --commit-message "Initial cluster snapshot"
```

## Detailed Walkthrough

### Understanding Contexts

Contexts in Kalco store cluster-specific information:

- **Kubeconfig path** - How to connect to your cluster
- **Output directory** - Where exported resources are saved
- **Description** - Human-readable context description
- **Labels** - Key-value pairs for organization

### Creating Your First Context

```bash
# Basic context creation
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-cluster-exports

# With additional metadata
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-cluster-exports \
  --description "Development cluster for testing" \
  --labels env=dev,team=developers,region=local
```

**Context Naming Tips:**
- Use descriptive names (e.g., `prod-eu-west`, `staging-us-east`)
- Include environment information
- Use consistent naming patterns

### Managing Contexts

View your contexts:

```bash
# List all contexts
kalco context list

# Show current context
kalco context current

# Show specific context details
kalco context show my-cluster
```

Switch between contexts:

```bash
# Switch to a different context
kalco context use another-cluster

# Switch back
kalco context use my-cluster
```

### Your First Export

The export command is Kalco's core functionality:

```bash
# Basic export
kalco export

# Export with Git integration
kalco export --git-push --commit-message "Initial backup"

# Export to specific directory
kalco export --output ./custom-backup

# Export specific namespaces
kalco export --namespaces default,kube-system
```

### Understanding Export Output

After export, you'll see:

1. **Resource Discovery** - Kalco finds all available resources
2. **Directory Creation** - Organized structure is created
3. **Git Setup** - Repository is initialized (if new)
4. **Resource Export** - YAML files are created
5. **Report Generation** - Change analysis report is created
6. **Git Commit** - Changes are committed with timestamp

### Output Structure

Your export creates this directory structure:

```
my-cluster-exports/
├── default/
│   ├── pods/
│   │   ├── nginx-pod.yaml
│   │   └── mysql-pod.yaml
│   ├── services/
│   │   ├── nginx-service.yaml
│   │   └── mysql-service.yaml
│   └── deployments/
│       ├── nginx-deployment.yaml
│       └── mysql-deployment.yaml
├── kube-system/
│   ├── pods/
│   └── services/
├── _cluster/
│   ├── nodes/
│   └── namespaces/
├── kalco-reports/
│   └── 20240819-145542-Initial-cluster-snapshot.md
└── kalco-config.json
```

### Git Integration

Kalco automatically handles Git operations:

1. **Repository Initialization** - Creates Git repo if needed
2. **File Addition** - Adds all exported resources
3. **Commit Creation** - Commits with your message or timestamp
4. **Remote Push** - Pushes to origin if `--git-push` is used

### Generated Reports

Kalco creates comprehensive reports in `kalco-reports/`:

- **Change Summary** - Overview of modifications
- **Resource Details** - Specific changes with diffs
- **Validation Results** - Cross-reference and orphaned resource checks
- **Git Information** - Commit details and history

## Common Scenarios

### Development Cluster

```bash
# Set up development context
kalco context set dev-cluster \
  --kubeconfig ~/.kube/dev-config \
  --output ./dev-exports \
  --description "Local development cluster" \
  --labels env=dev,team=developers

# Export development resources
kalco context use dev-cluster
kalco export --commit-message "Development snapshot"
```

### Production Cluster

```bash
# Set up production context
kalco context set prod-cluster \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform

# Export with Git push
kalco context use prod-cluster
kalco export --git-push --commit-message "Production backup $(date)"
```

### Multi-Cluster Setup

```bash
# Create contexts for different environments
kalco context set staging \
  --kubeconfig ~/.kube/staging-config \
  --output ./staging-exports \
  --labels env=staging,team=qa

kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --labels env=prod,team=platform

# Export all environments
kalco context use staging
kalco export --commit-message "Staging backup"

kalco context use production
kalco export --commit-message "Production backup"
```

## Best Practices

### Context Management

1. **Use descriptive names** for contexts
2. **Include labels** for better organization
3. **Set meaningful output directories**
4. **Document context purposes** with descriptions

### Export Strategy

1. **Start with basic exports** to understand the process
2. **Use meaningful commit messages** for better Git history
3. **Enable Git push** for team collaboration
4. **Exclude noisy resources** like events and replicasets

### Organization

1. **Group contexts by environment** (dev, staging, prod)
2. **Use consistent naming patterns** across your organization
3. **Include team and region information** in labels
4. **Regularly review and clean up** unused contexts

## Troubleshooting

### Common Issues

#### Context Not Found

```bash
Error: context 'my-cluster' not found
```

**Solution**: Use `kalco context list` to see available contexts.

#### Permission Denied

```bash
Error: failed to create output directory
```

**Solution**: Ensure write permissions for the output directory.

#### Git Not Found

```bash
Error: failed to initialize Git repository
```

**Solution**: Install Git and ensure it's in your PATH.

#### Kubernetes Connection Failed

```bash
Error: failed to create Kubernetes clients
```

**Solution**: Verify kubeconfig and cluster access.

### Getting Help

- **Command help**: `kalco <command> --help`
- **Verbose output**: Use `--verbose` flag for detailed information
- **Context help**: `kalco context --help`
- **Export help**: `kalco export --help`

## Next Steps

After your first successful export:

1. **Explore the output structure** to understand organization
2. **Review the generated report** for insights about your cluster
3. **Set up additional contexts** for other clusters
4. **Configure automated exports** for regular backups
5. **Read the [Commands Reference](../commands/index.md)** for advanced usage

## Examples

### Complete First Run

```bash
# 1. Verify installation
kalco version

# 2. Create context
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-cluster-exports \
  --description "My first Kubernetes cluster" \
  --labels env=dev,team=personal

# 3. Use context
kalco context use my-cluster

# 4. Export cluster
kalco export --git-push --commit-message "Initial snapshot"

# 5. Verify results
ls -la ./my-cluster-exports/
git log --oneline ./my-cluster-exports/
```

### Team Collaboration

```bash
# Load team member's context
kalco context load ~/team-exports/prod-cluster

# Use shared context
kalco context use prod-cluster

# Export with team context
kalco export --git-push --commit-message "Team backup $(date)"
```

---

*For more information about using Kalco, see the [Commands Reference](../commands/index.md) or run `kalco --help`.*
