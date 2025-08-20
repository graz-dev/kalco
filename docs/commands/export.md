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
- **Professional Reporting** - Comprehensive change analysis and tracking reports

## Syntax

```bash
kalco export [flags]
```

## Flags

### Git Integration

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--git-push` | Automatically push to remote origin | `false` | No |
| `--commit-message, -m` | Custom Git commit message | Timestamp-based | No |

### Execution Control

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--dry-run` | Show what would be exported | `false` | No |

## Basic Usage

### Simple Export

Export all cluster resources using the active context:

```bash
kalco export
```

This exports all resources to the output directory specified in the active context.

### Git Integration

Export with automatic Git commit:

```bash
kalco export --git-push --commit-message "Daily backup"
```

### Custom Commit Message

Use a custom commit message:

```bash
kalco export --commit-message "Weekly maintenance backup"
```

## Context-Based Configuration

The export command automatically uses the active context:

```bash
# Set and use a context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports

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
- **Git Information** - Commit details and history
- **Actionable Insights** - Recommendations and next steps

### Report Types

- **Initial Snapshot** - First export with complete resource inventory
- **Change Reports** - Incremental updates with modification details

## Advanced Usage

### Dry Run Mode

Preview what would be exported without making changes:

```bash
kalco export --dry-run
```

Output shows:
- Resources that would be exported
- Directory structure that would be created
- Git operations that would be performed

## Use Cases

### Production Backups

```bash
# Daily automated backup
kalco export \
  --git-push \
  --commit-message "Daily production backup - $(date)"

# Pre-deployment backup
kalco export \
  --commit-message "Pre-deployment backup v2.1.0"
```

### CI/CD Integration

```bash
#!/bin/bash
# Automated backup script

CLUSTER_NAME=$1

# Export cluster using active context
kalco export \
  --commit-message "Automated backup - $CLUSTER_NAME - $(date)"

# Push to remote if available
if [ -d ".git" ]; then
  if git remote get-url origin >/dev/null 2>&1; then
    git push origin main
  fi
fi

echo "Backup completed"
```

---

*For more information about the export command, run `kalco export --help` or see the [Commands Reference](index.md).*
