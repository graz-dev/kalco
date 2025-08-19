# Context Management

The `kalco context` command manages cluster configurations and settings, allowing you to work with multiple Kubernetes clusters through a unified interface.

## Overview

Contexts in Kalco store cluster-specific information including:
- **Kubeconfig path** - Connection details for the cluster
- **Output directory** - Where exported resources are stored
- **Description** - Human-readable context description
- **Labels** - Key-value pairs for organization and filtering
- **Metadata** - Creation and modification timestamps

## Subcommands

### `kalco context set`

Create or update a context with the specified configuration.

#### Syntax

```bash
kalco context set <name> [flags]
```

#### Arguments

- **`<name>`** - Unique name for the context (required)

#### Flags

| Flag | Description | Required | Default |
|------|-------------|----------|---------|
| `--kubeconfig` | Path to kubeconfig file | No | Current kubeconfig |
| `--output` | Output directory for exports | No | None |
| `--description` | Human-readable description | No | Empty |
| `--labels` | Labels in key=value format | No | Empty |

#### Examples

```bash
# Create a production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform

# Create a staging context with minimal configuration
kalco context set staging \
  --kubeconfig ~/.kube/staging-config \
  --output ./staging-exports

# Update existing context
kalco context set production \
  --description "Updated production cluster description"
```

### `kalco context list`

Display all available contexts with their configuration details.

#### Syntax

```bash
kalco context list
```

#### Output

The command displays:
- Context names with current context indicator (*)
- Description and configuration details
- Labels and metadata
- Creation and modification timestamps

#### Example Output

```
Available contexts:

* production
  Description: Production cluster for customer workloads
  Kubeconfig: ~/.kube/prod-config
  Output Dir: ./prod-exports
  Labels: env=prod, team=platform
  Created: 2024-01-15 10:30:00
  Updated: 2024-01-15 14:45:00

  staging
  Description: Staging cluster for testing
  Kubeconfig: ~/.kube/staging-config
  Output Dir: ./staging-exports
  Labels: env=staging, team=qa
  Created: 2024-01-10 09:15:00
  Updated: 2024-01-10 09:15:00

* = current context
```

### `kalco context use`

Switch to the specified context. This context will be used for future operations.

#### Syntax

```bash
kalco context use <name>
```

#### Arguments

- **`<name>`** - Name of the context to switch to (required)

#### Examples

```bash
# Switch to production context
kalco context use production

# Switch to staging context
kalco context use staging
```

#### Output

```
Switched to context 'production'
```

### `kalco context show`

Display detailed information about a specific context.

#### Syntax

```bash
kalco context show <name>
```

#### Arguments

- **`<name>`** - Name of the context to display (required)

#### Examples

```bash
# Show production context details
kalco context show production

# Show staging context details
kalco context show staging
```

#### Example Output

```
Context: production
Description: Production cluster for customer workloads
Kubeconfig: ~/.kube/prod-config
Output Directory: ./prod-exports
Labels:
  env: prod
  team: platform
Created: 2024-01-15 10:30:00
Updated: 2024-01-15 14:45:00
```

### `kalco context current`

Display information about the currently active context.

#### Syntax

```bash
kalco context current
```

#### Examples

```bash
# Show current context
kalco context current
```

#### Example Output

```
Current context: production
Description: Production cluster for customer workloads
Kubeconfig: ~/.kube/prod-config
Output Directory: ./prod-exports
Labels:
  env: prod
  team: platform
Created: 2024-01-15 10:30:00
Updated: 2024-01-15 14:45:00
```

### `kalco context delete`

Remove the specified context. Cannot delete the currently active context.

#### Syntax

```bash
kalco context delete <name>
```

#### Arguments

- **`<name>`** - Name of the context to delete (required)

#### Examples

```bash
# Delete staging context
kalco context delete staging

# Delete production context (must switch first)
kalco context use staging
kalco context delete production
```

#### Output

```
Context 'staging' deleted successfully
```

### `kalco context load`

Load a context configuration from an existing kalco directory by reading the `kalco-config.json` file.

#### Syntax

```bash
kalco context load <directory>
```

#### Arguments

- **`<directory>`** - Path to existing kalco export directory (required)

#### Examples

```bash
# Load context from existing export
kalco context load ./existing-kalco-export

# Load context from team member's export
kalco context load ~/team-exports/prod-cluster
```

#### Example Output

```
Context 'prod-cluster' loaded successfully from './existing-kalco-export'
   Kubeconfig: ~/.kube/prod-config
   Output Dir: ./existing-kalco-export
   Description: Production cluster export
   Labels: env=prod, team=platform
```

## Context Configuration

### Context File Structure

Contexts are stored in `~/.kalco/contexts.yaml`:

```yaml
production:
  name: production
  kubeconfig: ~/.kube/prod-config
  output_dir: ./prod-exports
  description: Production cluster for customer workloads
  labels:
    env: prod
    team: platform
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
  created_at: 2024-01-10T09:15:00Z
  updated_at: 2024-01-10T09:15:00Z
```

### Current Context

The currently active context is stored in `~/.kalco/current-context`:

```
production
```

## Best Practices

### Naming Conventions

- **Use descriptive names** that reflect the cluster purpose
- **Include environment information** (e.g., `prod-eu-west`, `staging-us-east`)
- **Use consistent naming patterns** across your organization

### Label Organization

- **Environment labels**: `env=prod`, `env=staging`, `env=dev`
- **Team labels**: `team=platform`, `team=qa`, `team=developers`
- **Region labels**: `region=eu-west`, `region=us-east`
- **Customer labels**: `customer=enterprise`, `customer=internal`

### Output Directory Structure

- **Use meaningful paths** that reflect the cluster purpose
- **Include timestamps** for historical tracking
- **Use consistent naming** across contexts

### Context Management

- **Regularly review** and clean up unused contexts
- **Document context purposes** with clear descriptions
- **Share context configurations** with team members
- **Use context switching** for multi-cluster operations

## Integration with Export

Contexts automatically configure export operations:

```bash
# Set production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports

# Use production context
kalco context use production

# Export automatically uses context settings
kalco export --git-push --commit-message "Production backup"
```

The export command will:
- Connect using the context's kubeconfig
- Save resources to the context's output directory
- Use context metadata in generated reports

## Troubleshooting

### Common Issues

#### Context Not Found

```
Error: context 'production' not found
```

**Solution**: Use `kalco context list` to see available contexts.

#### Cannot Delete Current Context

```
Error: cannot delete the currently active context
```

**Solution**: Switch to another context first, then delete.

#### Invalid Directory for Load

```
Error: directory './invalid-path' is not a valid kalco directory
```

**Solution**: Ensure the directory contains a `kalco-config.json` file.

#### Permission Denied

```
Error: failed to create output directory
```

**Solution**: Ensure write permissions for the output directory.

### Getting Help

- **Context help**: `kalco context --help`
- **Subcommand help**: `kalco context <subcommand> --help`
- **Verbose output**: Use `--verbose` flag for detailed information

## Examples

### Multi-Environment Setup

```bash
# Production environment
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --description "Production cluster for customer workloads" \
  --labels env=prod,team=platform,region=eu-west

# Staging environment
kalco context set staging \
  --kubeconfig ~/.kube/staging-config \
  --output ./staging-exports \
  --description "Staging cluster for testing and validation" \
  --labels env=staging,team=qa,region=eu-west

# Development environment
kalco context set development \
  --kubeconfig ~/.kube/dev-config \
  --output ./dev-exports \
  --description "Development cluster for local development" \
  --labels env=dev,team=developers,region=eu-west
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

### Context Switching Workflow

```bash
# Work with production
kalco context use production
kalco export --commit-message "Production backup"

# Switch to staging
kalco context use staging
kalco export --commit-message "Staging backup"

# Switch back to production
kalco context use production
kalco export --commit-message "Production update"
```

---

*For more information about context management, see the [Commands Reference](index.md) or run `kalco context --help`.*
