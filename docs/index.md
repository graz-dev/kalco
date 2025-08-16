# Kalco Documentation

Welcome to the comprehensive documentation for **Kalco** - Kubernetes Analysis & Lifecycle Control.

## What is Kalco?

Kalco is a powerful CLI tool that performs comprehensive analysis, validation, and export of all resources from your Kubernetes cluster into organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

## Quick Navigation

- [Getting Started](getting-started.md) - Installation and first steps
- [Commands Reference](commands/index.md) - Complete command documentation
- [Configuration](configuration.md) - Configuration options and examples
- [Use Cases](use-cases.md) - Real-world scenarios and examples
- [API Reference](api-reference.md) - Programmatic usage
- [Contributing](contributing.md) - How to contribute to Kalco
- [FAQ](faq.md) - Frequently asked questions

## Key Features

- **Complete Resource Discovery** - Automatically finds ALL available API resources
- **Comprehensive Coverage** - Includes both native K8s resources and Custom Resources (CRDs)
- **Structured Output** - Creates intuitive directory structures for easy navigation
- **Clean YAML** - Intelligently removes metadata fields that aren't useful for re-application
- **Lightning Fast** - Optimized for speed and efficiency
- **Git Integration** - Automatic version control with commit history and change tracking
- **Smart Reporting** - Generates detailed change reports with before/after comparisons
- **Cross-Reference Validation** - Analyzes exported resources for broken references
- **Orphaned Resource Detection** - Identifies resources no longer managed by controllers
- **Modern CLI Experience** - Intuitive commands with rich styling and helpful output

## Quick Start

```bash
# Install Kalco
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/main/scripts/install.sh | bash

# Export your entire cluster
kalco export

# Validate cluster resources
kalco validate

# Find orphaned resources
kalco analyze orphaned
```

## Community & Support

- **GitHub**: [graz-dev/kalco](https://github.com/graz-dev/kalco)
- **Issues**: [Report bugs or request features](https://github.com/graz-dev/kalco/issues)

---

Made with ❤️ for the Kubernetes community