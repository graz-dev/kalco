---
layout: default
title: Home
nav_order: 1
description: "Kalco - Kubernetes Analysis & Lifecycle Control. Extract, validate, analyze, and version control your entire cluster."
permalink: /
---

# Kalco Documentation
{: .fs-9 }

Kubernetes Analysis & Lifecycle Control - Extract, validate, analyze, and version control your entire cluster with comprehensive validation and Git integration.
{: .fs-6 .fw-300 }

[Get started now](#quick-start){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View it on GitHub](https://github.com/graz-dev/kalco){: .btn .fs-5 .mb-4 .mb-md-0 }

---

## What is Kalco?

Kalco is a powerful CLI tool that performs comprehensive analysis, validation, and export of all resources from your Kubernetes cluster into organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

### Key Features

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

## Navigation

<div class="code-example" markdown="1">

**Getting Started**
: [Installation and first steps]({% link getting-started.md %})

**Commands Reference**  
: [Complete command documentation]({% link commands/index.md %})

**Configuration**
: [Configuration options and examples]({% link configuration.md %})

**Use Cases**
: [Real-world scenarios and examples]({% link use-cases.md %})

**FAQ**
: [Frequently asked questions]({% link faq.md %})

**Contributing**
: [How to contribute to Kalco]({% link contributing.md %})

</div>

---

## Community & Support

- **GitHub**: [graz-dev/kalco](https://github.com/graz-dev/kalco)
- **Issues**: [Report bugs or request features](https://github.com/graz-dev/kalco/issues)  
- **Discussions**: [Community discussions](https://github.com/graz-dev/kalco/discussions)