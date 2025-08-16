---
layout: default
title: Commands Reference
nav_order: 3
has_children: true
---

# Commands Reference

Complete reference for all Kalco commands and options.

## 📋 Available Commands

- **[export]({{ site.baseurl }}/docs/commands/export)** - Export cluster resources to organized YAML files
- **[validate]({{ site.baseurl }}/docs/commands/validate)** - Validate cluster resources and cross-references
- **[analyze]({{ site.baseurl }}/docs/commands/analyze)** - Analyze cluster state and generate reports
- **[config]({{ site.baseurl }}/docs/commands/config)** - Manage Kalco configuration

## 🚩 Global Options

All commands support these global options:

```bash
--help, -h          Show help for the command
--version, -v       Show version information
--verbose           Enable verbose output
--quiet             Suppress non-error messages
--config            Path to configuration file
```

## 🔧 Command Structure

```bash
kalco <command> [subcommand] [flags] [arguments]
```

### Examples

```bash
# Basic export
kalco export

# Export with options
kalco export --output ./backup --namespaces default,kube-system

# Validate cluster
kalco validate --cross-references

# Show configuration
kalco config show
```

## 📚 Command Categories

### Core Operations
- **export** - Primary functionality for cluster resource extraction
- **validate** - Resource validation and health checks
- **analyze** - Cluster analysis and reporting

### Configuration & Management
- **config** - Configuration management and validation
- **completion** - Shell completion generation

## 🎯 Getting Help

### Command Help

```bash
# General help
kalco --help

# Command-specific help
kalco export --help
kalco validate --help
```

### Examples

```bash
# Show examples for export command
kalco export --help | grep -A 10 "Examples"

# Show all available flags
kalco export --help | grep -E "^  --"
```

## 🔍 Command Discovery

Explore available commands:

```bash
# List all commands
kalco --help

# Show command tree
kalco --help | grep -E "^  [a-z]"
```

## 📖 Next Steps

Explore the specific command categories to learn more about each command's options and usage patterns.
