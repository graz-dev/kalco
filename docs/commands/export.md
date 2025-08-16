---
layout: default
title: kalco export
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

The export creates the following directory structure:

```
cluster-backup/
├── default/                    # Namespace
│   ├── Pod/
│   │   ├── app-pod-1.yaml
│   │   └── app-pod-2.yaml
│   ├── Service/
│   │   └── app-service.yaml
│   └── Deployment/
│       └── app-deployment.yaml
├── kube-system/               # System namespace
│   ├── Pod/
│   ├── Service/
│   └── DaemonSet/
├── _cluster/                  # Cluster-scoped resources
│   ├── Node/
│   │   ├── node-1.yaml
│   │   └── node-2.yaml
│   ├── ClusterRole/
│   └── CustomResourceDefinition/
├── .git/                      # Git repository (if enabled)
└── kalco-reports/            # Generated reports
    ├── export-summary.yaml
    └── change-report.html
```

## Git Integration

When Git integration is enabled:

1. **Repository Initialization**: Creates a Git repository if none exists
2. **Change Tracking**: Commits changes with timestamps and metadata
3. **Remote Push**: Optionally pushes to remote origin
4. **Change Reports**: Generates detailed change reports between exports

### Git Workflow

```bash
# First export - initializes repository
kalco export --output ./cluster-repo --commit-message "Initial export"

# Subsequent exports - tracks changes
kalco export --output ./cluster-repo --commit-message "Weekly update"

# View change history
cd cluster-repo
git log --oneline
git diff HEAD~1 HEAD
```

## Filtering Options

### Namespace Filtering

```bash
# Single namespace
kalco export --namespaces production

# Multiple namespaces
kalco export --namespaces "default,kube-system,monitoring"

# All namespaces except system ones
kalco export --exclude-namespaces "kube-system,kube-public,kube-node-lease"
```

### Resource Type Filtering

```bash
# Core workload resources only
kalco export --resources "pods,services,deployments,configmaps,secrets"

# Exclude noisy resources
kalco export --exclude "events,replicasets,endpoints"

# Custom resources only
kalco export --crds-only
```

## Performance Considerations

- **Large Clusters**: Use namespace filtering to reduce export size
- **Network Resources**: Exclude events and endpoints for faster exports
- **Storage**: Consider using compression for large exports
- **Git History**: Large repositories may need periodic cleanup

## Troubleshooting

### Common Issues

**Permission Errors:**
```bash
# Check cluster access
kubectl auth can-i get pods --all-namespaces

# Use specific kubeconfig
kalco export --kubeconfig ~/.kube/production-config
```

**Large Exports:**
```bash
# Use filtering to reduce size
kalco export --exclude events,endpoints,replicasets

# Export specific namespaces
kalco export --namespaces production,staging
```

**Git Issues:**
```bash
# Check Git configuration
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Manual Git setup
cd output-directory
git init
git remote add origin <repository-url>
```

## Related Commands

- [`kalco validate`](validate.md) - Validate exported resources
- [`kalco analyze orphaned`](analyze.md#orphaned) - Find cleanup opportunities
- [`kalco report`](report.md) - Generate comprehensive reports

---

[← Commands Overview](index.md) | [Validate Command →](validate.md)