---
layout: default
title: kalco export
nav_order: 1
parent: Commands Reference
---

# kalco export

Export cluster resources to organized YAML files.

## Synopsis

The `export` command is Kalco's primary functionality. It discovers all available API resources (including CRDs) and exports them with clean metadata suitable for re-application.

```bash
kalco export [flags]
```

## Description

The export command creates an intuitive directory structure:
- **Namespaced resources**: `<output>/<namespace>/<kind>/<name>.yaml`
- **Cluster resources**: `<output>/_cluster/<kind>/<name>.yaml`

Features:
- Automatic resource discovery (native K8s + CRDs)
- Clean metadata removal for re-application
- Git integration for version control
- Flexible filtering options
- Progress tracking and detailed output

## Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--output` | `-o` | string | Output directory path (default: `./kalco-export-<timestamp>`) |
| `--namespaces` | `-n` | []string | Specific namespaces to export (comma-separated) |
| `--resources` | `-r` | []string | Specific resource types to export (comma-separated) |
| `--exclude` | | []string | Resource types to exclude (comma-separated) |
| `--git-push` | | bool | Automatically push changes to remote origin |
| `--commit-message` | `-m` | string | Custom Git commit message |
| `--dry-run` | | bool | Show what would be exported without writing files |

## Examples

### Basic Export

```bash
# Export entire cluster to timestamped directory
kalco export

# Export to specific directory
kalco export --output ./cluster-backup
```

### Filtered Export

```bash
# Export specific namespaces only
kalco export --namespaces default,kube-system,production

# Export specific resource types
kalco export --resources pods,services,deployments,configmaps

# Exclude noisy resources
kalco export --exclude events,replicasets,endpoints
```

### Git Integration

```bash
# Export with Git commit
kalco export --commit-message "Weekly cluster backup"

# Export and push to remote
kalco export --git-push --commit-message "Production snapshot $(date)"

# Export to existing Git repository
kalco export --output ./existing-repo/cluster-state --git-push
```

### Advanced Usage

```bash
# Dry run to see what would be exported
kalco export --dry-run --verbose

# Export production namespace with Git integration
kalco export \
  --namespaces production \
  --exclude events,replicasets \
  --git-push \
  --commit-message "Production backup - $(date +%Y-%m-%d)"

# Export for disaster recovery
kalco export \
  --output ./disaster-recovery/cluster-$(date +%Y%m%d) \
  --exclude events,endpoints,replicasets \
  --commit-message "DR backup"
```

## Output Structure

```
output-directory/
├── _cluster/                    # Cluster-scoped resources
│   ├── ClusterRole/
│   │   ├── admin.yaml
│   │   └── view.yaml
│   ├── ClusterRoleBinding/
│   └── StorageClass/
├── default/                     # Default namespace
│   ├── ConfigMap/
│   ├── Service/
│   └── ServiceAccount/
├── kube-system/                 # System namespace
│   ├── Deployment/
│   ├── Service/
│   └── ConfigMap/
└── kalco-reports/               # Analysis reports
    ├── Cluster-snapshot-2025-08-16-14-55-34.md
    └── Changes-and-validation-demo--2025-08-16-14-55-34.md
```

## Resource Filtering

### Include Specific Resources

```bash
# Export only core resources
kalco export --resources pods,services,deployments,configmaps

# Export with CRDs
kalco export --resources pods,services,mycustomresource
```

### Exclude Resources

```bash
# Exclude noisy resources
kalco export --exclude events,replicasets,endpoints

# Exclude specific CRDs
kalco export --exclude events,mycustomresource
```

## Git Integration

### Automatic Git Operations

```bash
# Initialize Git and commit
kalco export --git-push

# Custom commit message
kalco export --git-push --commit-message "Weekly backup"

# Push to remote origin
kalco export --git-push --commit-message "Production snapshot"
```

### Git Repository Setup

The export command will:
1. Initialize Git repository if not present
2. Add all exported files
3. Commit with timestamp or custom message
4. Push to remote origin if `--git-push` is specified

## Progress and Output

### Verbose Mode

```bash
kalco export --verbose
```

Shows:
- Resource discovery progress
- File writing operations
- Git operations
- Validation results

### Dry Run

```bash
kalco export --dry-run
```

Shows what would be exported without writing files:
- Resource counts by namespace
- File paths that would be created
- Git operations that would be performed

## Exit Codes

- `0` - Success
- `1` - General error
- `2` - Configuration error
- `3` - Kubernetes connection error
- `4` - File system error
- `5` - Git operation error

## Related Commands

- **[validate]({{ site.baseurl }}/docs/commands/validate)** - Validate exported resources
- **[analyze]({{ site.baseurl }}/docs/commands/analyze)** - Analyze cluster state
- **[config]({{ site.baseurl }}/docs/commands/config)** - Manage configuration
