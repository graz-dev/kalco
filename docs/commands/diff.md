---
layout: default
title: kalco diff
nav_order: 3
parent: Commands Reference
---

# Diff Command

The `kalco diff` command compares a local manifest file (INTENTION) with the cluster snapshot or live state (REALITY).

## Overview

The diff command helps you identify "drift" between what you have defined in your local configuration and what is currently deployed or captured in your snapshot.

Key features:
- **Deep Diffing** - Ignores irrelevant metadata (e.g., `managedFields`, `status`, `uid`) to focus on meaningful changes.
- **Snapshot Integration** - Automatically maps local files to their expected location in the snapshot structure.
- **Live Refresh** - Optionally refresh the snapshot from the live cluster before comparing.
- **Output Formats** - Supports both human-readable text and Markdown suitable for Pull Request comments.

## Syntax

```bash
kalco diff [flags]
```

## Flags

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--target` | Path to local manifest file | - | Yes |
| `--snapshot-dir` | Root directory of snapshot repository | - | Yes |
| `--refresh` | Fetch live resource from K8s via export | `false` | No |
| `--format` | Output format (`text`, `markdown`) | `text` | No |

## Usage Examples

### Basic Diff

Compare a local deployment file against the existing snapshot:

```bash
kalco diff \
  --target ./charts/my-app/deployment.yaml \
  --snapshot-dir ./cluster-exports
```

### Diff with Live Refresh

Update the snapshot from the cluster (via export) before performing the diff. This ensures you are comparing against the latest live state.

```bash
kalco diff \
  --target ./charts/my-app/deployment.yaml \
  --snapshot-dir ./cluster-exports \
  --refresh
```

### Markdown Output

Generate a Markdown-formatted diff, useful for automated CI/CD pipelines (e.g., GitHub Actions PR comments):

```bash
kalco diff \
  --target ./charts/my-app/deployment.yaml \
  --snapshot-dir ./cluster-exports \
  --format markdown
```

## Output Explained

- **🟢 NEW RESOURCE**: The resource exists locally but not in the snapshot.
- **✅ No changes detected**: The local resource matches the snapshot (ignoring K8s system fields).
- **🔴 Drift Detected**: Differences found. A diff will be shown highlighting added (`+`) or removed (`-`) lines.
- **⚠️ Snapshot is stale**: If not using `--refresh`, this warning appears if the snapshot is older than 24 hours.
