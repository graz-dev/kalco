---
layout: default
title: Configuration
nav_order: 3
parent: Getting Started
---

# Configuration

This guide covers configuring Kalco for your environment, including context management, output customization, and advanced settings.

## Overview

Kalco uses a hierarchical configuration system:

- **Global Configuration** - Application-wide settings
- **Context Configuration** - Cluster-specific settings
- **Environment Variables** - System-level overrides
- **Command Line Flags** - Runtime overrides

## Configuration Directory

Kalco stores configuration in `~/.kalco/`:

```
~/.kalco/
├── contexts.yaml      # Context configurations
├── current-context    # Currently active context
└── config.json        # Global configuration
```

### Initial Setup

The configuration directory is created automatically on first run:

```bash
# First command creates ~/.kalco/
kalco context list
```

## Context Configuration

### Context Structure

Each context stores cluster-specific information:

```yaml
production:
  name: production
  kubeconfig: ~/.kube/prod-config
  output_dir: ./prod-exports
  description: Production cluster for customer workloads
  labels:
    env: prod
    team: platform
    region: eu-west
  created_at: 2024-01-15T10:30:00Z
  updated_at: 2024-01-15T14:45:00Z

staging:
  name: staging
  kubeconfig: ~/.kube/staging-config
  output_dir: ./staging-exports
  description: Staging cluster for testing
  labels:
    env: staging
    team: qa
    region: eu-west
  created_at: 2024-01-10T09:15:00Z
  updated_at: 2024-01-10T09:15:00Z
```

### Context Fields

| Field | Description | Required | Default |
|-------|-------------|----------|---------|
| `name` | Unique context identifier | Yes | - |
| `kubeconfig` | Path to kubeconfig file | No | Current kubeconfig |
| `output_dir` | Export output directory | No | None |
| `description` | Human-readable description | No | Empty |
| `labels` | Key-value pairs for organization | No | Empty |
| `created_at` | Creation timestamp | Auto | Current time |
| `updated_at` | Last modification timestamp | Auto | Current time |

### Managing Contexts

#### Create Context

```bash
# Basic context
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-exports

# With metadata
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform,region=eu-west
```

#### Update Context

```bash
# Update description
kalco context set production \
  --description "Updated production cluster description"

# Update labels
kalco context set production \
  --labels env=prod,team=platform,region=eu-west,customer=enterprise

# Update output directory
kalco context set production \
  --output ./new-prod-exports
```

#### Delete Context

```bash
# Delete context (must not be current)
kalco context delete staging
```

### Context Best Practices

#### Naming Conventions

- **Environment-based**: `prod`, `staging`, `dev`
- **Region-based**: `prod-eu-west`, `prod-us-east`
- **Team-based**: `prod-platform`, `prod-data`
- **Customer-based**: `prod-enterprise`, `prod-startup`

#### Label Organization

```yaml
# Environment labels
env: prod|staging|dev|testing

# Team labels
team: platform|qa|developers|data

# Region labels
region: eu-west|us-east|ap-southeast

# Customer labels
customer: enterprise|startup|internal

# Project labels
project: website|api|analytics
```

#### Output Directory Strategy

```bash
# Environment-based
./exports/prod/
./exports/staging/
./exports/dev/

# Date-based
./exports/prod/2024-01-15/
./exports/prod/2024-01-16/

# Project-based
./exports/prod/website/
./exports/prod/api/
```

## Global Configuration

### Global Settings

Global configuration in `~/.kalco/config.json`:

```json
{
  "default_kubeconfig": "~/.kube/config",
  "default_output_dir": "./kalco-exports",
  "git_auto_push": false,
  "git_auto_commit": true,
  "report_format": "markdown",
  "exclude_resources": ["events", "replicasets"],
  "include_resources": [],
  "verbose_output": false,
  "color_output": true
}
```

### Configuration Options

| Option | Description | Default | Type |
|--------|-------------|---------|------|
| `default_kubeconfig` | Default kubeconfig path | `~/.kube/config` | string |
| `default_output_dir` | Default output directory | `./kalco-exports` | string |
| `git_auto_push` | Automatically push Git changes | `false` | boolean |
| `git_auto_commit` | Automatically commit changes | `true` | boolean |
| `report_format` | Report output format | `markdown` | string |
| `exclude_resources` | Resources to exclude by default | `["events"]` | array |
| `include_resources` | Resources to include by default | `[]` | array |
| `verbose_output` | Enable verbose output by default | `false` | boolean |
| `color_output` | Enable colored output by default | `true` | boolean |

## Environment Variables

### Supported Variables

Kalco respects these environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `KUBECONFIG` | Path to kubeconfig file | `~/.kube/config` |
| `KALCO_CONFIG_DIR` | Configuration directory | `~/.kalco` |
| `KALCO_DEFAULT_OUTPUT` | Default output directory | `./kalco-exports` |
| `KALCO_GIT_AUTO_PUSH` | Enable auto Git push | `false` |
| `KALCO_VERBOSE` | Enable verbose output | `false` |
| `NO_COLOR` | Disable colored output | `false` |

### Environment Variable Usage

```bash
# Set environment variables
export KUBECONFIG=~/.kube/prod-config
export KALCO_DEFAULT_OUTPUT=./prod-exports
export KALCO_GIT_AUTO_PUSH=true

# Run kalco with environment settings
kalco export
```

## Export Configuration

### Default Export Settings

Configure default export behavior:

```bash
# Set default exclude resources
kalco context set production \
  --exclude events,replicasets,endpoints

# Set default include resources
kalco context set production \
  --resources deployments,services,configmaps,secrets
```

### Export Flags

Override configuration with command-line flags:

```bash
# Override output directory
kalco export --output ./custom-backup

# Override namespace filtering
kalco export --namespaces default,kube-system

# Override resource filtering
kalco export --resources pods,services

# Override exclusions
kalco export --exclude events,replicasets

# Override Git behavior
kalco export --no-commit
kalco export --git-push
```

## Git Configuration

### Repository Settings

Kalco automatically configures Git repositories:

```bash
# Initialize Git repository
kalco export --output ./new-export

# Configure remote origin
cd ./new-export
git remote add origin <your-repo-url>

# Enable auto-push
kalco export --git-push
```

### Git Integration Options

| Option | Description | Default |
|--------|-------------|---------|
| `--git-push` | Automatically push to remote | `false` |
| `--commit-message` | Custom commit message | Timestamp-based |
| `--no-commit` | Skip Git operations | `false` |

### Git Best Practices

1. **Use meaningful commit messages**
2. **Configure remote origins for collaboration**
3. **Use branches for different environments**
4. **Regularly clean up old export directories**

## Output Configuration

### Directory Structure

Customize export output organization:

```bash
# Default structure
<output_dir>/
├── <namespace>/
│   ├── <resource_kind>/
│   │   └── <resource_name>.yaml
│   └── ...
├── _cluster/
│   ├── <resource_kind>/
│   │   └── <resource_name>.yaml
│   └── ...
├── kalco-reports/
│   └── <timestamp>-<commit-message>.md
└── kalco-config.json
```

### Resource Organization

- **Namespaced Resources**: `<namespace>/<kind>/<name>.yaml`
- **Cluster Resources**: `_cluster/<kind>/<name>.yaml`
- **Reports**: `kalco-reports/<timestamp>-<commit-message>.md`
- **Configuration**: `kalco-config.json`

## Advanced Configuration

### Resource Filtering

Configure resource inclusion and exclusion:

```bash
# Exclude noisy resources
kalco context set production \
  --exclude events,replicasets,endpoints,pods

# Include only specific resources
kalco context set production \
  --resources deployments,services,configmaps,secrets

# Combine filters
kalco export \
  --namespaces default,monitoring \
  --resources deployments,services \
  --exclude events
```

### Namespace Filtering

```bash
# Export specific namespaces
kalco export --namespaces default,kube-system

# Exclude system namespaces
kalco export --namespaces default,monitoring,applications

# Export all except system
kalco export --exclude-namespaces kube-system,kube-public
```

### Custom Output Formats

Kalco supports multiple output formats:

```bash
# Markdown reports (default)
kalco export --report-format markdown

# JSON reports
kalco export --report-format json

# HTML reports
kalco export --report-format html
```

## Configuration Examples

### Development Environment

```bash
# Development context
kalco context set dev \
  --kubeconfig ~/.kube/dev-config \
  --output ./dev-exports \
  --description "Development cluster for testing" \
  --labels env=dev,team=developers

# Development export settings
kalco export \
  --namespaces default,dev-apps \
  --exclude events,replicasets \
  --commit-message "Development snapshot"
```

### Production Environment

```bash
# Production context
kalco context set prod \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform,region=eu-west

# Production export settings
kalco export \
  --namespaces production,monitoring \
  --exclude events,replicasets,pods \
  --git-push \
  --commit-message "Production backup $(date)"
```

### Multi-Cluster Setup

```bash
# Staging context
kalco context set staging \
  --kubeconfig ~/.kube/staging-config \
  --output ./staging-exports \
  --labels env=staging,team=qa

# Production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --labels env=prod,team=platform

# Export all environments
for env in staging production; do
  kalco context use $env
  kalco export --git-push --commit-message "$env backup $(date)"
done
```

## Troubleshooting

### Configuration Issues

#### Context Not Found

```bash
Error: context 'production' not found
```

**Solution**: Use `kalco context list` to see available contexts.

#### Invalid Configuration

```bash
Error: invalid context configuration
```

**Solution**: Check context file format and syntax.

#### Permission Denied

```bash
Error: failed to create configuration directory
```

**Solution**: Ensure write permissions for `~/.kalco/`.

### Configuration Validation

Validate your configuration:

```bash
# Check context configuration
kalco context show production

# Verify current context
kalco context current

# Test export with current configuration
kalco export --dry-run
```

### Getting Help

- **Configuration help**: `kalco context --help`
- **Export help**: `kalco export --help`
- **Verbose output**: Use `--verbose` flag
- **Dry run**: Use `--dry-run` to preview configuration

## Next Steps

After configuring Kalco:

1. **Test your configuration** with `--dry-run`
2. **Set up multiple contexts** for different environments
3. **Configure automated exports** for regular backups
4. **Customize resource filtering** for your needs
5. **Read the [Commands Reference](../commands/index.md)** for advanced usage

---

*For more configuration help, see the [Commands Reference](../commands/index.md) or run `kalco --help`.*
