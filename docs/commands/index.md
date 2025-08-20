---
layout: default
title: Commands Reference
nav_order: 3
has_children: true
---

# Commands Reference

This section provides comprehensive documentation for all Kalco commands and their options.

## Command Overview

Kalco provides a focused set of commands designed for professional Kubernetes cluster management:

| Command | Description | Usage |
|---------|-------------|-------|
| `kalco context` | Manage cluster contexts | `kalco context set/list/use/load` |
| `kalco export` | Export cluster resources | `kalco export [flags]` |
| `kalco version` | Version information | `kalco version` |

## Global Flags

All Kalco commands support these global flags:

| Flag | Description | Default |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig file | `~/.kube/config` |
| `--no-color` | Disable colored output | `false` |
| `--help, -h` | Show help information | - |

## Context Management

The `kalco context` command manages cluster configurations and settings.

### Subcommands

- **`set`** - Create or update a context
- **`list`** - List all available contexts
- **`use`** - Switch to a specific context
- **`show`** - Display context details
- **`current`** - Show current active context
- **`delete`** - Remove a context
- **`load`** - Import context from existing directory

### Usage Examples

```bash
# Create a production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform

# List all contexts
kalco context list

# Switch to production context
kalco context use production

# Show current context
kalco context current

# Load context from existing export
kalco context load ./existing-kalco-export
```

## Resource Export

The `kalco export` command exports Kubernetes cluster resources with Git integration.

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--git-push` | Automatically push to remote origin | `false` |
| `--commit-message, -m` | Custom Git commit message | Timestamp-based |
| `--dry-run` | Show what would be exported | `false` |

### Usage Examples

```bash
# Basic export using active context
kalco export

# Export with Git integration
kalco export --git-push --commit-message "Weekly backup"

# Custom commit message
kalco export --commit-message "Pre-deployment backup"

# Dry run to see what would be exported
kalco export --dry-run
```

## Version Information

The `kalco version` command displays version and build information.

```bash
 kalco version
```

Output includes:
- Version number
- Git commit hash
- Build timestamp
- Go version used

## Command Aliases

Some commands provide convenient aliases:

| Command | Aliases |
|---------|---------|
| `kalco export` | `dump`, `backup` |
| `kalco context list` | `ls` |
| `kalco context show` | `info` |

## Error Handling

Kalco provides clear error messages and exit codes:

- **Exit Code 0** - Success
- **Exit Code 1** - General error
- **Exit Code 2** - Configuration error
- **Exit Code 3** - Kubernetes connection error

## Best Practices

### Context Management

1. **Use descriptive names** for contexts (e.g., `prod-eu-west`, `staging-us-east`)
2. **Include labels** for better organization and filtering
3. **Set output directories** that reflect the cluster purpose
4. **Regularly clean up** unused contexts

### Resource Export

1. **Use meaningful commit messages** for better Git history
2. **Enable Git push** for team collaboration
3. **Use dry-run** to verify export configuration
4. **Leverage context-based configuration** for consistent exports

### Automation

1. **Integrate with CI/CD** pipelines for automated backups
2. **Use context switching** for multi-cluster operations
3. **Set up regular exports** for change tracking
4. **Use Git integration** for version control

## Troubleshooting

### Common Issues

- **Permission denied**: Ensure write access to output directory
- **Git not found**: Install Git for version control functionality
- **Kubernetes connection failed**: Verify kubeconfig and cluster access
- **Context not found**: Use `kalco context list` to see available contexts

### Getting Help

- **Command help**: `kalco <command> --help`
- **Global help**: `kalco --help`
- **GitHub issues**: Report bugs and request features

---

*For more detailed information about specific commands, see the individual command documentation pages.*
