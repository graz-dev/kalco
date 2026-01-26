# Kalco - Kubernetes Analysis & Lifecycle Control

**Professional CLI tool for Kubernetes cluster management, analysis, and lifecycle control**

[![Release](https://img.shields.io/github/v/release/graz-dev/kalco)](https://github.com/graz-dev/kalco/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://graz-dev.github.io/kalco)

*Extract, analyze, and version control your entire Kubernetes cluster with Git integration.*

## 🚀 Overview

Kalco acts as a bridge between your Kubernetes cluster **Reality** and your GitOps **Intention**. It allows you to:
- **Snapshot** your cluster state into organized YAML (Deployment, Service, CRDs...).
- **Detect Drift** between Git manifests (Helm/Kustomize/Raw) and the live cluster.
- **Audit** changes over time with automated Git commits.

## 📚 Documentation

Full documentation is available at [**graz-dev.github.io/kalco**](https://graz-dev.github.io/kalco).

## ⚡ Quick Start

### Installation

**Homebrew (macOS/Linux)**
```bash
brew install graz-dev/tap/kalco
```

**Install Script**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

### Basic Usage

1. **Setup Context**:
   ```bash
   kalco context set prod --output ./prod-backup
   ```

2. **Snapshot Cluster**:
   ```bash
   kalco export --git-push
   ```

3. **Verify GitOps Drift** (New!):
   ```bash
   kalco diff --target ./my-app/deploy.yaml --snapshot-dir ./prod-backup --refresh
   ```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
