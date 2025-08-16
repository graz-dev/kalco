---
layout: default
title: Commands Reference
nav_order: 3
has_children: true
---

# Commands Reference

Complete reference for all Kalco commands, options, and examples.

## Quick Navigation

- **[Export Commands]({{ site.baseurl }}/docs/commands/export)** - Export cluster resources
- **[Validation Commands]({{ site.baseurl }}/docs/commands/validation)** - Validate resources and references
- **[Analysis Commands]({{ site.baseurl }}/docs/commands/analysis)** - Analyze cluster health and resources
- **[Report Commands]({{ site.baseurl }}/docs/commands/reports)** - Generate comprehensive reports
- **[Utility Commands]({{ site.baseurl }}/docs/commands/utilities)** - Helper commands and utilities

## Command Structure

All Kalco commands follow this pattern:

```bash
kalco <command> [subcommand] [options] [arguments]
```

### Global Options

These options are available for all commands:

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--help` | `-h` | Show help information | |
| `--version` | `-v` | Show version information | |
| `--verbose` | | Enable verbose output | `false` |
| `--quiet` | | Suppress output | `false` |
| `--kubeconfig` | | Path to kubeconfig file | Auto-detected |
| `--context` | | Kubernetes context to use | Current context |

## Common Patterns

### Basic Command

```bash
kalco export
```

### With Options

```bash
kalco export --output-dir ./my-cluster --verbose
```

### With Subcommands

```bash
kalco validate --output json
```

### With Arguments

```bash
kalco export --namespaces default,kube-system
```

## Output Formats

Most commands support multiple output formats:

- **Default** - Human-readable text output
- **JSON** - Machine-readable JSON format
- **YAML** - YAML format for configuration
- **HTML** - Rich HTML reports
- **Table** - Tabular data output

Example:
```bash
kalco validate --output json
kalco analyze orphaned --output html
kalco report --output yaml
```

## Examples

### Export Cluster Resources

```bash
# Export all resources
kalco export

# Export specific namespaces
kalco export --namespaces default,production

# Export with Git integration
kalco export --git-push --commit-message "Daily backup"
```

### Validate Resources

```bash
# Basic validation
kalco validate

# Detailed validation report
kalco validate --detailed --output html

# Validate specific resource types
kalco validate --resources deployments,services
```

### Analyze Cluster

```bash
# Find orphaned resources
kalco analyze orphaned

# Security analysis
kalco analyze security --output html

# Resource usage analysis
kalco analyze usage --detailed
```

## Next Steps

Explore the specific command categories to learn more about each command's options and usage patterns.
