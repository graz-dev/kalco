---
layout: default
title: Commands Reference
nav_order: 3
has_children: true
---

# Commands Reference

Complete reference for all Kalco commands and options.

## 📋 Available Commands

- **[context]({{ site.baseurl }}/docs/commands/context)** - Manage cluster contexts and configurations
- **[export]({{ site.baseurl }}/docs/commands/export)** - Export cluster resources to organized YAML files

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

# Show version
kalco version
```

## 📚 Command Categories

### Core Operations
- **context** - Context management for different clusters and configurations
- **export** - Primary functionality for cluster resource extraction

## 🎯 Getting Help

### Command Help

```bash
# General help
kalco --help

# Command-specific help
kalco export --help
kalco context --help
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
