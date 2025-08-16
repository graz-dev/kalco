---
layout: default
title: FAQ
---

# Frequently Asked Questions

Common questions and answers about Kalco.

## General Questions

### What is Kalco?
 
Kalco (Kubernetes Analysis & Lifecycle Control) is a comprehensive CLI tool for Kubernetes cluster management. It provides automated resource export, validation, analysis, and lifecycle management capabilities.

### How is Kalco different from kubectl?

While kubectl is the standard Kubernetes CLI for day-to-day operations, Kalco focuses on cluster-wide analysis, backup, and lifecycle management:

- **kubectl**: Resource manipulation, debugging, cluster interaction
- **Kalco**: Cluster analysis, backup, validation, optimization, reporting

They complement each other - use kubectl for operations and Kalco for analysis and management.

### Is Kalco safe to use in production?

Yes! Kalco is designed with production safety in mind:

- **Read-only operations** by default (export, validate, analyze)
- **No cluster modifications** unless explicitly requested
- **Dry-run mode** available for testing
- **Comprehensive logging** for audit trails

## Installation & Setup

### Which platforms are supported?

Kalco supports all major platforms:

- **Linux**: x86_64, ARM64
- **macOS**: Intel, Apple Silicon (M1/M2)
- **Windows**: x86_64, ARM64

### Do I need special permissions?

Kalco requires read access to Kubernetes resources. The specific permissions depend on what you want to export:

```yaml
# Minimum RBAC for basic export
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: kalco-reader
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["get", "list"]
```

### Can I use Kalco with multiple clusters?

Yes! You can use Kalco with multiple clusters by:

1. **Different kubeconfig files**: `kalco export --kubeconfig ~/.kube/prod-config`
2. **Different contexts**: Switch context with `kubectl config use-context`
3. **Scripted workflows**: Automate multi-cluster operations

## Usage Questions

### How do I export only specific namespaces?

Use the `--namespaces` flag:

```bash
# Single namespace
kalco export --namespaces production

# Multiple namespaces
kalco export --namespaces "production,staging,development"
```

### Can I exclude certain resource types?

Yes, use the `--exclude` flag:

```bash
# Exclude noisy resources
kalco export --exclude "events,endpoints,replicasets"

# Export only specific types
kalco export --resources "deployments,services,configmaps"
```

### How do I handle large clusters?

For large clusters, consider:

1. **Namespace filtering**: Export specific namespaces
2. **Resource filtering**: Exclude unnecessary resource types
3. **Parallel processing**: Use multiple Kalco instances for different namespaces
4. **Incremental exports**: Regular small exports vs. large full exports

```bash
# Optimized for large clusters
kalco export \
  --namespaces production \
  --exclude "events,endpoints,replicasets" \
  --output ./production-backup
```

### How does Git integration work?

Kalco can automatically manage Git repositories:

1. **Initialize repository** if none exists
2. **Commit changes** with timestamps and metadata
3. **Track changes** between exports
4. **Push to remote** if configured

```bash
# Enable Git integration
kalco export --git-push --commit-message "Weekly backup"
```

## Troubleshooting

### "Permission denied" errors

This usually indicates insufficient RBAC permissions:

```bash
# Check your permissions
kubectl auth can-i get pods --all-namespaces
kubectl auth can-i list customresourcedefinitions

# Use a different kubeconfig
kalco export --kubeconfig ~/.kube/admin-config
```

### "Connection refused" errors

Check your Kubernetes connection:

```bash
# Test basic connectivity
kubectl cluster-info

# Check kubeconfig
kubectl config current-context
kubectl config view

# Use specific kubeconfig
kalco export --kubeconfig /path/to/kubeconfig
```

### Large export files

If exports are too large:

```bash
# Use filtering
kalco export --exclude "events,endpoints,replicasets"

# Export specific namespaces
kalco export --namespaces "production,staging"

# Use compression (if available)
kalco export --compress
```

### Git integration issues

Common Git problems and solutions:

```bash
# Configure Git identity
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Check Git repository status
cd output-directory
git status
git log --oneline

# Manual Git setup
git init
git remote add origin <repository-url>
```

## Advanced Usage

### Can I use Kalco in CI/CD pipelines?

Absolutely! Kalco is designed for automation:

```yaml
# GitHub Actions example
- name: Cluster Backup
  run: |
    kalco export \
      --output ./backup \
      --no-color \
      --commit-message "CI backup $(date)"
```

### How do I integrate with monitoring systems?

Kalco supports multiple output formats for integration:

```bash
# JSON output for programmatic processing
kalco validate --output json | jq '.summary'

# Metrics export for Prometheus
kalco analyze usage --output json | \
  jq -r '.metrics[] | "kalco_\(.name) \(.value)"'
```

### Can I customize the output format?

Yes, Kalco supports multiple output formats:

- **YAML**: Default, human-readable
- **JSON**: Machine-readable, API integration
- **HTML**: Rich reports with styling
- **Table**: Console-friendly tabular output

```bash
kalco validate --output json
kalco report --output html
kalco resources list --output table
```

### How do I handle Custom Resource Definitions (CRDs)?

Kalco automatically discovers and exports CRDs:

```bash
# Include all CRDs (default)
kalco export

# List only CRDs
kalco resources list --crds-only

# Export specific CRDs
kalco export --resources "certificates.cert-manager.io,issuers.cert-manager.io"
```

## Performance & Optimization

### How can I speed up exports?

Several optimization strategies:

1. **Parallel processing**: Use multiple Kalco instances
2. **Filtering**: Reduce the scope of exports
3. **Incremental exports**: Export only changed resources
4. **Local caching**: Cache discovery results

```bash
# Optimized export
kalco export \
  --namespaces production \
  --exclude "events,endpoints" \
  --parallel 4
```

### What about memory usage?

Kalco is designed to be memory-efficient:

- **Streaming processing**: Resources are processed as they're discovered
- **Configurable batch sizes**: Control memory usage
- **Garbage collection**: Automatic cleanup of temporary data

### How do I monitor Kalco performance?

Use the verbose flag for detailed timing information:

```bash
kalco export --verbose
```

## Security & Compliance

### Is my cluster data secure?

Kalco follows security best practices:

- **No data transmission**: All processing is local
- **Configurable output**: Control what gets exported
- **Audit logging**: Track all operations
- **RBAC integration**: Respects Kubernetes permissions

### Can I use Kalco for compliance auditing?

Yes! Kalco provides several compliance features:

```bash
# Security analysis
kalco analyze security --output json

# Validation reporting
kalco validate --output html

# Resource inventory
kalco resources list --detailed
```

### How do I handle sensitive data?

Best practices for sensitive data:

1. **Exclude secrets**: Use `--exclude secrets`
2. **Namespace filtering**: Export only non-sensitive namespaces
3. **Post-processing**: Clean sensitive data from exports
4. **Access control**: Secure export directories

```bash
# Safe export excluding sensitive resources
kalco export \
  --exclude "secrets,serviceaccounts" \
  --namespaces "production,staging"
```

## Getting Help

### Where can I get support?

- **Documentation**: [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
- **GitHub Issues**: [Report bugs or request features](https://github.com/graz-dev/kalco/issues)
- **Discussions**: [Community discussions](https://github.com/graz-dev/kalco/discussions)
- **CLI Help**: `kalco [command] --help`

### How do I report a bug?

1. **Check existing issues**: Search [GitHub Issues](https://github.com/graz-dev/kalco/issues)
2. **Gather information**: Version, OS, Kubernetes version, error messages
3. **Create detailed report**: Include steps to reproduce
4. **Provide context**: Cluster size, configuration, use case

### How can I contribute?

We welcome contributions! See our [Contributing Guide](contributing.md) for details.

---

[← Use Cases](use-cases.md) | [Contributing →](contributing.md)