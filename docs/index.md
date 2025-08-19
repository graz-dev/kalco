---
layout: home
title: Home
nav_order: 1
---

# Kalco Documentation

Welcome to the Kalco documentation. Kalco is a professional CLI tool for Kubernetes cluster analysis, resource extraction, validation, and lifecycle management.

## What is Kalco?

Kalco transforms your Kubernetes cluster management experience by providing a comprehensive, automated, and intelligent approach to cluster analysis and lifecycle control. Whether you're managing production workloads, ensuring compliance, or planning migrations, Kalco has you covered.

## Key Features

- **Context Management** - Manage multiple Kubernetes clusters through unified contexts
- **Resource Export** - Export cluster resources with professional organization
- **Git Integration** - Automatic version control with commit history and change tracking
- **Report Generation** - Professional change analysis and validation reports
- **Enterprise Ready** - Designed for production environments and team collaboration

## Quick Start

1. **Install Kalco:**
   ```bash
   curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
   ```

2. **Create a Context:**
   ```bash
   kalco context set production \
     --kubeconfig ~/.kube/prod-config \
     --output ./prod-exports
   ```

3. **Export Cluster Resources:**
   ```bash
   kalco export --git-push --commit-message "Initial backup"
   ```

## Documentation Sections

### Getting Started
- [Installation](getting-started/installation.md) - Install Kalco on your system
- [First Run](getting-started/first-run.md) - Get started with your first export
- [Configuration](getting-started/configuration.md) - Configure Kalco for your environment

### Commands Reference
- [Command Overview](commands/index.md) - Complete command reference
- [Export Command](commands/export.md) - Export cluster resources
- [Context Management](commands/context.md) - Manage cluster contexts

### Examples
- [Quickstart Demo](../examples/quickstart.sh) - Comprehensive example script
- [Production Workflows](examples/production.md) - Real-world usage examples
- [CI/CD Integration](examples/cicd.md) - Automation examples

## Use Cases

### DevOps Teams
- Automated cluster backups and disaster recovery
- Change tracking and compliance auditing
- Environment replication and configuration management

### Platform Engineers
- Infrastructure as Code and GitOps workflows
- Team collaboration and context sharing
- Migration support and configuration validation

### Security Teams
- Configuration auditing and compliance reporting
- Access control through context management
- Security validation and change monitoring

## Architecture

Kalco is built with a modular architecture designed for enterprise use:

- **Context Manager** - Handles cluster configurations and settings
- **Resource Exporter** - Discovers and exports Kubernetes resources
- **Git Integration** - Manages version control operations
- **Report Generator** - Creates change analysis and validation reports
- **Validation Engine** - Performs cross-reference and orphaned resource checks

## Design Principles

- **Professional Interface** - Clean, emoji-free CLI design
- **Minimal Dependencies** - Focused functionality without bloat
- **Enterprise Ready** - Production-grade reliability and performance
- **Team Collaboration** - Shared configurations and context sharing
- **Automation First** - Designed for CI/CD and automated workflows

## Support

- **GitHub Issues**: [Report bugs and request features](https://github.com/graz-dev/kalco/issues)
- **Discussions**: [Join community discussions](https://github.com/graz-dev/kalco/discussions)
- **Documentation**: This site and [GitHub repository](https://github.com/graz-dev/kalco)

## Contributing

We welcome contributions to improve Kalco and its documentation. See our [contributing guidelines](https://github.com/graz-dev/kalco/blob/master/CONTRIBUTING.md) for details.

---

*Kalco - Built with ❤️ for the Kubernetes community*
