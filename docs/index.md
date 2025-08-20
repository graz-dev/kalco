---
layout: home
title: Home
nav_order: 1
---

# Kalco

Welcome to the Kalco documentation. Kalco is a CLI tool for Kubernetes cluster dumping and audit.

## What is Kalco?

Kalco transforms your Kubernetes cluster management experience by providing a comprehensive, automated, and intelligent approach to cluster lifecycle control. Whether you're managing production workloads, ensuring compliance, or planning migrations, Kalco has you covered.

## Key Features

- **Context Management** - Manage multiple Kubernetes clusters through unified contexts
- **Resource Export** - Export cluster resources with professional organization
- **Git Integration** - Automatic version control with commit history and change tracking
- **Report Generation** - Professional change analysis and tracking reports

## Quick Start

1. **Install Kalco:**
   ```bash
   curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
   ```

   or use a package manager:

   ```bash
   brew tap graz-dev/kalco
   brew install kalco
   ```

2. **Create a Context:**
   ```bash
   kalco context set my-cluster \
     --kubeconfig ~/.kube/config \
     --output ./my-cluster-dump
   ```

3. **Export Cluster Resources:**
   ```bash
   kalco export --commit-message "My first dump"
   ```

## Documentation Sections

### Getting Started

- **[Install Kalco]({{ site.baseurl }}/getting-started/installation)**
- **[First Run]({{ site.baseurl }}/getting-started/first-run)** - Export your first cluster

### Commands Reference
- **[Command Overview]({{ site.baseurl }}/commands/index.md)** - Complete command reference
- **[Context Management]({{ site.baseurl }}/commands/context.md)** - Manage cluster contexts
- **[Export Command]({{ site.baseurl }}/commands/export.md)** - Export cluster resources

## Support

- **GitHub Issues**: [Report bugs and request features](https://github.com/graz-dev/kalco/issues)
- **Documentation**: This site and [GitHub repository](https://github.com/graz-dev/kalco)

## Contributing

We welcome contributions to improve Kalco and its documentation. See our [contributing guidelines](https://github.com/graz-dev/kalco/blob/master/CONTRIBUTING.md) for details.

---

*Kalco - Built with ❤️ for the Kubernetes community*
