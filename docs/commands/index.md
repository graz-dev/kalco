---
layout: default
title: Commands Reference
---

# Commands Reference

Complete reference for all Kalco commands and their options.

## Command Overview

| Command | Description |
|---------|-------------|
| [`export`](export.md) | Export cluster resources to organized YAML files |
| [`validate`](validate.md) | Validate cluster resources for issues and broken references |
| [`analyze`](analyze.md) | Analyze cluster for optimization opportunities |
| [`resources`](resources.md) | List and inspect cluster resources |
| [`report`](report.md) | Generate comprehensive cluster reports |
| [`config`](config.md) | Manage kalco configuration |
| [`completion`](completion.md) | Generate shell completion scripts |
| [`version`](version.md) | Show version information |

## Global Flags

These flags are available for all commands:

| Flag | Description |
|------|-------------|
| `--kubeconfig string` | Path to kubeconfig file |
| `--verbose, -v` | Enable verbose output |
| `--no-color` | Disable colored output |
| `--help, -h` | Show help information |

## Command Categories

### Core Operations
- [`kalco export`](export.md) - Primary functionality for exporting cluster resources
- [`kalco validate`](validate.md) - Validate exported or live cluster resources

### Analysis & Insights
- [`kalco analyze orphaned`](analyze.md#orphaned) - Find orphaned resources
- [`kalco analyze usage`](analyze.md#usage) - Analyze resource usage
- [`kalco analyze security`](analyze.md#security) - Security posture analysis

### Resource Management
- [`kalco resources list`](resources.md#list) - List available resource types
- [`kalco resources describe`](resources.md#describe) - Describe specific resources
- [`kalco resources count`](resources.md#count) - Count resources by type

### Reporting & Documentation
- [`kalco report`](report.md) - Generate comprehensive reports

### Configuration & Setup
- [`kalco config`](config.md) - Manage configuration settings
- [`kalco completion`](completion.md) - Set up shell completion
- [`kalco version`](version.md) - Version and build information

## Quick Examples

### Export Operations
```bash
# Basic export
kalco export

# Export with filtering
kalco export --namespaces production,staging --exclude events

# Export with Git integration
kalco export --git-push --commit-message "Weekly backup"
```

### Validation & Analysis
```bash
# Validate cluster
kalco validate --output json

# Find orphaned resources
kalco analyze orphaned --detailed

# Generate security report
kalco analyze security --output yaml
```

### Resource Discovery
```bash
# List all resources
kalco resources list

# List only CRDs
kalco resources list --crds-only

# Count resources by namespace
kalco resources count --by-namespace
```

## Command Aliases

Many commands have convenient aliases:

- `export` → `dump`, `backup`
- `validate` → `check`, `lint`
- `resources` → `res`, `resource`
- `resources list` → `resources ls`

## Output Formats

Most commands support multiple output formats:

- `table` (default) - Human-readable table format
- `json` - JSON format for programmatic use
- `yaml` - YAML format for configuration
- `html` - HTML format for reports (where applicable)

Use the `--output` or `-o` flag to specify the format:

```bash
kalco validate --output json
kalco analyze orphaned --output yaml
kalco report --output html
```

---

[← Getting Started](../getting-started.md) | [Export Command →](export.md)