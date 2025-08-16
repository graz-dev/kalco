---
layout: default
title: Configuration
parent: Getting Started
nav_order: 3
---

# Configuration

Customize Kalco to fit your environment and workflow requirements.

## Overview

Kalco works out of the box with sensible defaults, but you can customize its behavior through:

- **Command-line options** - Per-command configuration
- **Environment variables** - System-wide settings
- **Configuration files** - Persistent settings
- **Project-specific configs** - Per-project customization

## Command-Line Options

### Global Options

These options are available for all commands:

| Option | Description | Default |
|--------|-------------|---------|
| `--verbose` | Enable detailed logging | `false` |
| `--quiet` | Suppress output | `false` |
| `--kubeconfig` | Path to kubeconfig file | Auto-detected |
| `--context` | Kubernetes context to use | Current context |
| `--output-dir` | Output directory path | Timestamped directory |

### Export Options

Customize the export process:

```bash
# Basic export with options
kalco export \
  --output-dir ./my-cluster \
  --verbose \
  --exclude events,pods \
  --include-namespaces default,production
```

| Option | Description | Default |
|--------|-------------|---------|
| `--output-dir` | Output directory | `./kalco-export-<timestamp>` |
| `--exclude` | Resource types to exclude | None |
| `--include-namespaces` | Namespaces to include | All namespaces |
| `--exclude-namespaces` | Namespaces to exclude | None |
| `--git-push` | Enable Git integration | `false` |
| `--commit-message` | Custom commit message | Timestamp-based |

## Environment Variables

Set these environment variables for persistent configuration:

```bash
# Set in your shell profile (.bashrc, .zshrc, etc.)
export KALCO_OUTPUT_DIR="/path/to/exports"
export KALCO_KUBECONFIG="$HOME/.kube/config"
export KALCO_VERBOSE="true"
export KALCO_GIT_PUSH="true"
```

### Available Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `KALCO_OUTPUT_DIR` | Default output directory | Current directory |
| `KALCO_KUBECONFIG` | Path to kubeconfig | Auto-detected |
| `KALCO_VERBOSE` | Enable verbose output | `false` |
| `KALCO_GIT_PUSH` | Enable Git integration | `false` |
| `KALCO_COMMIT_MESSAGE` | Default commit message | Timestamp |

## Configuration Files

### Project Configuration

Create a `.kalco.yml` file in your project directory:

```yaml
# .kalco.yml
output_dir: ./cluster-exports
exclude:
  - events
  - pods
  - endpoints
include_namespaces:
  - default
  - production
exclude_namespaces:
  - kube-system
  - kube-public
git:
  enabled: true
  commit_message: "Cluster export {timestamp}"
  push: false
validation:
  enabled: true
  output_format: html
analysis:
  orphaned_resources: true
  broken_references: true
  security_scan: true
```

### Global Configuration

Create a global configuration file at `~/.kalco/config.yml`:

```yaml
# ~/.kalco/config.yml
defaults:
  output_dir: ~/kalco-exports
  verbose: false
  git_push: false
  
clusters:
  production:
    context: prod-cluster
    output_dir: ~/kalco-exports/production
    git_push: true
    
  staging:
    context: staging-cluster
    output_dir: ~/kalco-exports/staging
    git_push: false
    
  development:
    context: dev-cluster
    output_dir: ~/kalco-exports/dev
    git_push: false
```

## Advanced Configuration

### Resource Filtering

Fine-tune which resources get exported:

```bash
# Export only specific resource types
kalco export --resources deployments,services,configmaps

# Exclude specific resource types
kalco export --exclude events,pods,endpoints

# Export resources with specific labels
kalco export --label-selector "app=myapp,env=production"

# Export resources in specific namespaces
kalco export --namespaces default,production
```

### Git Configuration

Customize Git integration behavior:

```bash
# Initialize Git repository
kalco export --git-init

# Commit changes with custom message
kalco export --commit-message "Production backup $(date)"

# Push to remote repository
kalco export --git-push --remote origin

# Configure Git user
kalco export --git-user "Kalco Bot" --git-email "kalco@company.com"
```

### Output Customization

Control the output format and structure:

```bash
# Custom output directory
kalco export --output-dir /backups/cluster-$(date +%Y%m%d)

# Flatten directory structure
kalco export --flatten-directories

# Include resource metadata
kalco export --include-metadata

# Exclude status fields
kalco export --exclude-status
```

## Configuration Precedence

Kalco follows this order of precedence (highest to lowest):

1. **Command-line options** - Override all other settings
2. **Project configuration** - `.kalco.yml` in current directory
3. **Environment variables** - System-wide settings
4. **Global configuration** - `~/.kalco/config.yml`
5. **Default values** - Built-in defaults

## Examples

### Development Environment

```yaml
# .kalco.yml for development
output_dir: ./dev-cluster
exclude:
  - events
  - pods
  - endpoints
git:
  enabled: true
  commit_message: "Dev cluster export"
  push: false
```

### Production Environment

```yaml
# .kalco.yml for production
output_dir: /backups/production
exclude:
  - events
  - pods
  - endpoints
  - secrets
git:
  enabled: true
  commit_message: "Production backup {timestamp}"
  push: true
validation:
  enabled: true
  output_format: html
```

### CI/CD Pipeline

```bash
# Export in CI/CD pipeline
kalco export \
  --output-dir ./cluster-snapshot \
  --exclude events,pods \
  --git-push \
  --commit-message "CI/CD export $(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  --verbose
```

## Troubleshooting Configuration

### Check Current Configuration

```bash
# Show effective configuration
kalco config show

# Validate configuration file
kalco config validate .kalco.yml

# List configuration sources
kalco config sources
```

### Common Issues

**Configuration not loaded:**
- Check file permissions
- Verify YAML syntax
- Ensure file is in correct location

**Options not working:**
- Check option precedence
- Verify option names
- Check for typos in configuration

**Environment variables ignored:**
- Ensure variables are exported
- Check variable names (case-sensitive)
- Restart shell after setting variables

## Next Steps

With configuration set up, you can:

1. **[Run Advanced Exports]({{ site.baseurl }}/docs/commands/export)** - Use advanced filtering and options
2. **[Automate with Scripts]({{ site.baseurl }}/docs/use-cases/automation)** - Create automated backup scripts
3. **[Integrate with CI/CD]({{ site.baseurl }}/docs/use-cases/ci-cd)** - Add to your deployment pipeline
4. **[Customize Output]({{ site.baseurl }}/docs/commands/export)** - Tailor exports to your needs
