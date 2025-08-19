---
layout: default
title: kalco export
nav_order: 1
parent: Commands Reference
---

# Export Command

The `kalco export` command exports Kubernetes cluster resources to organized YAML files with automatic Git integration and professional report generation.

## Overview

The export command is the core functionality of Kalco, providing:

- **Complete Resource Discovery** - Automatically finds all available API resources including CRDs
- **Structured Organization** - Creates intuitive directory structures by namespace and resource type
- **Clean YAML Output** - Optimizes metadata for re-application
- **Git Integration** - Automatic version control with commit history
- **Professional Reporting** - Comprehensive change analysis and validation reports

## Syntax

```bash
kalco export [flags]
```

## Flags

### Output Configuration

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--output, -o` | Output directory path | `./kalco-export-<timestamp>` | No |
| `--namespaces, -n` | Specific namespaces to export | All namespaces | No |
| `--resources, -r` | Specific resource types to export | All resources | No |
| `--exclude` | Resource types to exclude | None | No |

### Git Integration

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--git-push` | Automatically push to remote origin | `false` | No |
| `--commit-message, -m` | Custom Git commit message | Timestamp-based | No |
| `--no-commit` | Skip Git commit operations | `false` | No |

### Execution Control

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--dry-run` | Show what would be exported | `false` | No |
| `--verbose, -v` | Enable verbose output | `false` | No |

## Basic Usage

### Simple Export

Export all cluster resources to a timestamped directory:

```bash
kalco export
```

This creates a directory like `./kalco-export-20240819-145542/` and exports all resources.

### Custom Output Directory

Specify a custom output directory:

```bash
kalco export --output ./my-cluster-backup
```

### Namespace Filtering

Export resources from specific namespaces:

```bash
# Single namespace
kalco export --namespaces default

# Multiple namespaces
kalco export --namespaces default,kube-system,monitoring

# Exclude system namespaces
kalco export --namespaces default,monitoring --exclude kube-system
```

### Resource Type Filtering

Export specific resource types:

```bash
# Core resources
kalco export --resources pods,services,deployments

# All resources in a category
kalco export --resources "*.apps/v1"

# Exclude noisy resources
kalco export --exclude events,replicasets,endpoints
```

## Git Integration

### Automatic Git Setup

Kalco automatically initializes Git repositories for new export directories:

```bash
kalco export --output ./new-cluster-export
```

The command will:
1. Create the output directory
2. Initialize a Git repository
3. Export cluster resources
4. Create initial commit
5. Generate change report

### Custom Commit Messages

Use meaningful commit messages for better Git history:

```bash
kalco export --commit-message "Production backup - $(date)"
kalco export --commit-message "Weekly maintenance backup"
kalco export --commit-message "Before deployment v2.1.0"
```

### Remote Push

Automatically push changes to remote repositories:

```bash
kalco export --git-push --commit-message "Automated backup"
```

**Note**: This requires the directory to have a remote origin configured.

### Skip Git Operations

Export without Git integration:

```bash
kalco export --no-commit
```

This is useful for:
- One-time exports
- CI/CD pipelines where Git is handled separately
- Testing export functionality

## Output Structure

Kalco creates a professional, organized directory structure:

```
<output_dir>/
├── <namespace>/
│   ├── <resource_kind>/
│   │   ├── <resource_name>.yaml
│   │   └── ...
│   └── ...
├── _cluster/
│   ├── <resource_kind>/
│   │   ├── <resource_name>.yaml
│   │   └── ...
│   └── ...
├── kalco-reports/
│   └── <timestamp>-<commit-message>.md
└── kalco-config.json
```

### Directory Organization

- **Namespaced Resources**: `<output>/<namespace>/<kind>/<name>.yaml`
- **Cluster Resources**: `<output>/_cluster/<kind>/<name>.yaml`
- **Reports**: `<output>/kalco-reports/<timestamp>-<commit-message>.md`
- **Configuration**: `<output>/kalco-config.json`

### File Naming

Resources are saved with descriptive filenames:
- `nginx-deployment.yaml`
- `mysql-service.yaml`
- `redis-configmap.yaml`

## Context Integration

When using contexts, export automatically uses context settings:

```bash
# Set production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports

# Use production context
kalco context use production

# Export uses context settings automatically
kalco export --git-push --commit-message "Production backup"
```

The export command will:
- Connect using the context's kubeconfig
- Save resources to the context's output directory
- Use context metadata in generated reports

## Report Generation

Kalco automatically generates comprehensive reports for each export:

### Report Location

Reports are saved in `kalco-reports/` directory with descriptive names:
- `20240819-145542-Production-backup.md`
- `20240819-160000-Weekly-maintenance.md`

### Report Content

Each report includes:
- **Change Summary** - Overview of modifications
- **Resource Details** - Specific changes with diffs
- **Validation Results** - Cross-reference and orphaned resource checks
- **Git Information** - Commit details and history
- **Actionable Insights** - Recommendations and next steps

### Report Types

- **Initial Snapshot** - First export with complete resource inventory
- **Change Reports** - Incremental updates with modification details
- **Validation Reports** - Cross-reference and orphaned resource analysis

## Advanced Usage

### Dry Run Mode

Preview what would be exported without making changes:

```bash
kalco export --dry-run --output ./preview-export
```

Output shows:
- Resources that would be exported
- Directory structure that would be created
- Git operations that would be performed

### Verbose Output

Enable detailed logging for debugging:

```bash
kalco export --verbose --output ./debug-export
```

Shows:
- Resource discovery progress
- Export operations in detail
- Git operation details
- Validation results

### Resource Filtering Examples

```bash
# Export only application resources
kalco export --resources deployments,services,configmaps,secrets

# Exclude system and temporary resources
kalco export --exclude events,replicasets,endpoints,pods

# Focus on specific application namespace
kalco export --namespaces myapp --resources deployments,services
```

## Use Cases

### Production Backups

```bash
# Daily automated backup
kalco export \
  --output ./prod-backups/$(date +%Y%m%d) \
  --git-push \
  --commit-message "Daily production backup - $(date)"

# Pre-deployment backup
kalco export \
  --output ./prod-backups/pre-deploy-v2.1.0 \
  --commit-message "Pre-deployment backup v2.1.0"
```

### CI/CD Integration

```bash
#!/bin/bash
# Automated backup script

CLUSTER_NAME=$1
BACKUP_DIR="./backups/${CLUSTER_NAME}/$(date +%Y%m%d-%H%M%S)"

# Export cluster
kalco export \
  --output "$BACKUP_DIR" \
  --exclude events,replicasets,pods \
  --commit-message "Automated backup - $CLUSTER_NAME - $(date)"

# Push to remote if available
if [ -d "$BACKUP_DIR/.git" ]; then
  cd "$BACKUP_DIR"
  if git remote get-url origin >/dev/null 2>&1; then
    git push origin main
  fi
fi

echo "Backup completed: $BACKUP_DIR"
```

---

*For more information about the export command, run `kalco export --help` or see the [Commands Reference](index.md).*
