---
layout: default
title: First Run
nav_order: 2
parent: Getting Started
---

# First Run

This guide walks you through your first Kalco export, from setting up a context to generating your first cluster snapshot.

## Prerequisites

Before starting, ensure you have:

- **Kalco installed** - See [Installation]({{ site.baseurl }}/getting-started/installation) if needed
- **Kubernetes access** - Valid kubeconfig or in-cluster access
- **Git installed** (optional) - For version control functionality

## Quick Start

### 1. Verify Installation

First, confirm Kalco is working:

```bash
kalco version
```

You should see version information and build details.

### 2. Check Available Commands

Explore what Kalco can do:

```bash
kalco --help
```

This shows the available commands: `context`, `export` and `version`.

### 3. Set Up Your First Context

Create a context for your cluster:

```bash
kalco context set my-cluster \
  --kubeconfig ~/.kube/config \
  --output ./my-cluster-exports \
  --description "My first Kalco dump" \
  --labels env=dev,team=personal
```

### 4. Use the Context

Activate your context:

```bash
kalco context use my-cluster
```

### 5. Export Your Cluster

Perform your first export:

```bash
kalco export --commit-message "Initial cluster snapshot"
```

---

*For more information about using Kalco, see the [Commands Reference](../commands/index.md) or run `kalco --help`.*