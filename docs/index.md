---
layout: default
title: Kalco Documentation
description: "Kubernetes Analysis & Lifecycle Control - Extract, validate, analyze, and version control your entire cluster."
---

# Kalco Documentation

**Kubernetes Analysis & Lifecycle Control** - Extract, validate, analyze, and version control your entire cluster with comprehensive validation and Git integration.

<div style="margin: 20px 0;">
  <a href="#quick-start" class="btn" style="background-color: #159957; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; margin-right: 10px;">Get Started</a>
  <a href="https://github.com/graz-dev/kalco" class="btn" style="background-color: #606c76; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">View on GitHub</a>
</div>

---

## What is Kalco?

Kalco is a powerful CLI tool that performs comprehensive analysis, validation, and export of all resources from your Kubernetes cluster into organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

## Key Features

- ✅ **Complete Resource Discovery** - Automatically finds ALL available API resources
- ✅ **Comprehensive Coverage** - Includes both native K8s resources and Custom Resources (CRDs)  
- ✅ **Structured Output** - Creates intuitive directory structures for easy navigation
- ✅ **Clean YAML** - Intelligently removes metadata fields that aren't useful for re-application
- ✅ **Lightning Fast** - Optimized for speed and efficiency
- ✅ **Git Integration** - Automatic version control with commit history and change tracking
- ✅ **Smart Reporting** - Generates detailed change reports with before/after comparisons
- ✅ **Cross-Reference Validation** - Analyzes exported resources for broken references
- ✅ **Orphaned Resource Detection** - Identifies resources no longer managed by controllers
- ✅ **Modern CLI Experience** - Intuitive commands with rich styling and helpful output

---

## Quick Start

### Installation

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

### Basic Usage

```bash
# Export your entire cluster
kalco export

# Validate cluster resources
kalco validate

# Find orphaned resources
kalco analyze orphaned

# Generate comprehensive report
kalco report
```

---

## Documentation Navigation

| Section | Description |
|---------|-------------|
| [**Getting Started**](getting-started.md) | Installation and first steps |
| [**Commands Reference**](commands/index.md) | Complete command documentation |
| [**Configuration**](configuration.md) | Configuration options and examples |
| [**Use Cases**](use-cases.md) | Real-world scenarios and examples |
| [**FAQ**](faq.md) | Frequently asked questions |
| [**Contributing**](contributing.md) | How to contribute to Kalco |

---

## Community & Support

- 🐙 **GitHub**: [graz-dev/kalco](https://github.com/graz-dev/kalco)
- 🐛 **Issues**: [Report bugs or request features](https://github.com/graz-dev/kalco/issues)  
- 💬 **Discussions**: [Community discussions](https://github.com/graz-dev/kalco/discussions)

---

*Made with ❤️ for the Kubernetes community*