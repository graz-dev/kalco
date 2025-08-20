---
layout: default
title: kalco context
nav_order: 2
parent: Commands Reference
---

# kalco context

Manage Kalco contexts for different Kubernetes clusters and configurations.

## Overview

A Kalco context defines a complete configuration for working with a specific Kubernetes cluster:

- **Kubeconfig file** - Path to the cluster configuration file
- **Output directory** - Default directory for exports
- **Dynamic labels** - Key-value pairs for context identification
- **Description** - Human-readable description of the context purpose

Contexts make it easy to switch between different clusters and configurations without specifying the same flags repeatedly.

## 📋 Available Commands

- **[context]({{ site.baseurl }}/docs/commands/context)** - Manage cluster contexts and configurations
- **[export]({{ site.baseurl }}/docs/commands/export)** - Export cluster resources to organized YAML files
- **[load]({{ site.baseurl }}/docs/commands/load)** - Load context configuration from an existing kalco directory

## 🚩 Global Options

All commands support these global options:

```bash
--help, -h          Show help for the command
--version, -v       Show version information
--no-color          Disable colored output
--kubeconfig        Path to kubeconfig file
```

## 🔧 Command Structure

```bash
kalco <command> [subcommand] [flags] [arguments]
```

### Examples

```bash
# Basic export
kalco export

# Export with Git integration
kalco export --git-push --commit-message "Daily backup"

# Manage contexts
kalco context set production --kubeconfig ~/.kube/prod-config --output ./prod-exports
kalco context use production

# Load existing context
kalco context load ./existing-kalco-export

# Show version
kalco version
```

## Commands

### kalco context set

Create or update a context with the specified configuration.

```bash
kalco context set <name> [flags]
```

#### Flags

| Flag | Description | Example |
|------|-------------|---------|
| `--kubeconfig` | Path to kubeconfig file | `--kubeconfig ~/.kube/prod-config` |
| `--output` | Output directory for exports | `--output ./prod-exports` |
| `--description` | Human-readable description | `--description "Production cluster"` |
| `--labels` | Labels in key=value format | `--labels env=prod --labels team=platform` |

#### Examples

```bash
# Create a production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster" \
  --labels env=prod \
  --labels team=platform

# Create a development context
kalco context set dev \
  --kubeconfig ~/.kube/dev-config \
  --output ./dev-exports \
  --description "Development cluster" \
  --labels env=dev \
  --labels team=backend

# Update an existing context
kalco context set production \
  --output ./new-prod-exports \
  --labels env=prod \
  --labels team=platform \
  --labels region=eu-west
```

**Note**: When specifying an `--output` directory, Kalco will:
- Create the directory if it doesn't exist
- Initialize a Git repository
- Create a `kalco-config.json` file with context metadata
- Make an initial commit

### kalco context list

Display all available contexts with their configuration details.

```bash
kalco context list
```

#### Output

```
Available contexts:

* production
  Description: Production cluster
  Kubeconfig: /Users/user/.kube/prod-config
  Output Dir: ./prod-exports
  Labels: env=prod, team=platform
  Created: 2025-08-17 15:30:45
  Updated: 2025-08-17 15:30:45

  dev
  Description: Development cluster
  Kubeconfig: /Users/user/.kube/dev-config
  Output Dir: ./dev-exports
  Labels: env=dev, team=backend
  Created: 2025-08-17 14:20:30
  Updated: 2025-08-17 14:20:30

* = current context
```

### kalco context use

Switch to the specified context. This context will be used for future operations.

```bash
kalco context use <name>
```

#### Examples

```bash
# Switch to production context
kalco context use production

# Switch to development context
kalco context use dev
```

### kalco context current

Display information about the currently active context.

```bash
kalco context current
```

#### Output

```
Current context: production
Description: Production cluster
Kubeconfig: /Users/user/.kube/prod-config
Output Directory: ./prod-exports
Labels:
  env: prod
  team: platform
Created: 2025-08-17 15:30:45
Updated: 2025-08-17 15:30:45
```

### kalco context show

Display detailed information about a specific context.

```bash
kalco context show <name>
```

#### Examples

```bash
# Show production context details
kalco context show production

# Show development context details
kalco context show dev
```

### kalco context delete

Delete the specified context. Can now delete the current context as well.

```bash
kalco context delete <name>
```

#### Examples

```bash
# Delete any context (including current)
kalco context delete old-cluster

# Delete current context (will clear current context)
kalco context delete production
```

### kalco context load

Load a context configuration from an existing kalco directory by reading the `kalco-config.json` file.

```bash
kalco context load <directory>
```

This is useful for importing contexts from existing kalco exports or for team collaboration.

#### Examples

```bash
# Load context from existing kalco export
kalco context load ./my-cluster-export

# Load context from team member's export
kalco context load ./team-prod-export

# Load context from backup directory
kalco context load ./backups/production-2025-08-17
```

#### Requirements

The directory must contain a valid `kalco-config.json` file with the following structure:

```json
{
  "context_name": "production",
  "kubeconfig": "/path/to/kubeconfig",
  "output_dir": "./exports/production",
  "labels": {
    "env": "production",
    "team": "platform"
  },
  "description": "Production cluster",
  "created_at": "2025-08-17T15:30:45Z",
  "version": "1.0"
}
```

## Context Integration

### With Export Command

When you run `kalco export`, the command automatically uses the active context:

```bash
# Set and use a context
kalco context set my-cluster \
  --kubeconfig ~/.kube/my-config \
  --output ./my-exports

kalco context use my-cluster

# Export using context configuration
kalco export  # Uses my-cluster context automatically
```

### Context Priority

Context settings have priority over command-line flags:

```bash
# Context specifies output directory
kalco context set prod --output ./prod-exports
kalco context use prod

# This will export to ./prod-exports, not ./custom
kalco export --output ./custom
```

### Override Context

You can still override context settings with flags:

```bash
# Context specifies output directory
kalco context set prod --output ./prod-exports
kalco context use prod

# Override with flag (higher priority)
kalco export --output ./override-output
```

## Configuration Files

Contexts are stored in `~/.kalco/`:

```
~/.kalco/
├── contexts.yaml      # Context definitions
└── current-context    # Currently active context
```

### contexts.yaml

```yaml
production:
  name: production
  kubeconfig: /Users/user/.kube/prod-config
  output_dir: ./prod-exports
  labels:
    env: prod
    team: platform
  description: Production cluster
  created_at: 2025-08-17T15:30:45+02:00
  updated_at: 2025-08-17T15:30:45+02:00

dev:
  name: dev
  kubeconfig: /Users/user/.kube/dev-config
  output_dir: ./dev-exports
  labels:
    env: dev
    team: backend
  description: Development cluster
  created_at: 2025-08-17T14:20:30+02:00
  updated_at: 2025-08-17T14:20:30+02:00
```

### current-context

```
production
```

## Best Practices

### Naming Conventions

- Use descriptive names: `production`, `staging`, `dev-east`, `prod-eu-west`
- Include environment and region information in names or labels
- Use consistent naming across team members

### Label Strategy

- **Environment**: `env=prod`, `env=staging`, `env=dev`
- **Team**: `team=platform`, `team=backend`, `team=frontend`
- **Region**: `region=eu-west`, `region=us-east`, `region=asia`
- **Purpose**: `purpose=testing`, `purpose=monitoring`, `purpose=backup`

### Organization

```bash
# Production contexts
kalco context set prod-eu-west \
  --kubeconfig ~/.kube/prod-eu-west \
  --output ./exports/prod-eu-west \
  --labels env=prod,region=eu-west,team=platform

kalco context set prod-us-east \
  --kubeconfig ~/.kube/prod-us-east \
  --output ./exports/prod-us-east \
  --labels env=prod,region=us-east,team=platform

# Development contexts
kalco context set dev-backend \
  --kubeconfig ~/.kube/dev-backend \
  --output ./exports/dev-backend \
  --labels env=dev,team=backend,purpose=testing

kalco context set dev-frontend \
  --kubeconfig ~/.kube/dev-frontend \
  --output ./exports/dev-frontend \
  --labels env=dev,team=frontend,purpose=testing
```

## Troubleshooting

### Common Issues

**Context not found**
```bash
Error: context 'production' not found
```
Solution: Use `kalco context list` to see available contexts.

**Cannot delete current context**
```bash
Error: cannot delete current context 'production'. Switch to another context first
```
Solution: Switch to a different context first, then delete.

**Invalid label format**
```bash
Error: invalid label format: env=prod,team=platform (expected key=value)
```
Solution: Use separate `--labels` flags: `--labels env=prod --labels team=platform`

### Context Validation

Contexts are automatically validated:

- Kubeconfig file must exist
- Output directory must be writable
- Labels must be in `key=value` format

### Migration from Flags

If you're currently using flags, you can migrate to contexts:

```bash
# Before (using flags)
kalco export --kubeconfig ~/.kube/prod-config --output ./prod-exports

# After (using context)
kalco context set prod \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports
kalco context use prod
kalco export  # Uses context automatically
```

## Related Commands

- **[kalco export]({{ site.baseurl }}/commands/export)** - Export cluster resources
