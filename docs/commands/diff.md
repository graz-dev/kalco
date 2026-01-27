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
| `--namespace` | Override namespace for local resources | - | No |
| `--ignore-base` | Path to base manifest (from destination branch) for Smart Drift Detection | - | No |
| `--ignore-extra-fields` | Ignore K8s default fields missing from local manifest | `false` | No |

## Advanced Features

### Smart Drift Detection (`--ignore-base`)

When validating a Pull Request, you often have changes that are *intentional*. 
- **Snapshot**: Old State
- **Base Branch**: Current State
- **PR Branch**: New State

If you just diff **PR vs Snapshot**, you see *all* differences.
By providing `--ignore-base <path-to-file-in-base-branch>`, Kalco calculates:
`(Base -> PR)` changes are **INTENTIONAL**.
It then masks these changes when comparing **(PR vs Snapshot)**.
Result: You *only* see drift (unexpected changes), not your PR changes.

### Ignore Defaults (`--ignore-extra-fields`)
Kubernetes often adds default fields (e.g., `dnsPolicy`, `strategy`, `revisionHistoryLimit`).
If your local manifest doesn't specify them, a strict diff shows them as specific values on the cluster vs missing locally.
Use `--ignore-extra-fields` to ensure Kalco only validates fields *you explicitly defined*.

## Output Explained

- **🟢 NEW RESOURCE**: The resource exists locally but not in the snapshot.
- **✅ No changes detected**: The local resource matches the snapshot (ignoring K8s system fields).
- **🔴 Drift Detected**: Differences found. A diff will be shown highlighting added (`+`) or removed (`-`) lines.
- **⚠️ Snapshot is stale**: If not using `--refresh`, this warning appears if the snapshot is older than 24 hours.
