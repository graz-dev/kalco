---
layout: default
title: kalco validate
nav_order: 4
parent: Commands Reference
---

# Validate Command

The `kalco validate` command is a high-level wrapper designed specifically for **CI/CD Pipelines**. It automates the entire drift detection workflow for Pulse/Merge Requests.

## Overview

Unlike `kalco diff`, which compares a single file, `kalco validate`:
1.  **Auto-Detects Changes**: Finds which manifests changed between your current branch and the base branch (e.g., `origin/master`).
2.  **Fetches Base Content**: Automatically retrieves the "before" state of files from git to enable [Smart Drift Detection]({{ site.baseurl }}/commands/diff#smart-drift-detection---ignore-base).
3.  **Generates Report**: Produces a consolidated Markdown report of all validations.
4.  **Ignores Defaults**: Automatically enables `--ignore-extra-fields` to reduce noise.

## Syntax

```bash
kalco validate [flags]
```

## Flags

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `--base` | Base branch/ref to compare against | `origin/master` | No |
| `--snapshot-dir` | Directory containing snapshots | - | Yes |
| `--dirs` | Directories to scan for changes | `.` | No |
| `--namespace` | Override namespace for resources | - | No |
| `--report-file` | File path to write Markdown report to | - | No |

## Usage Examples

### Basic CI Integration (GitHub Actions)

In a GitHub Action workflow, use `kalco validate` to check changed manifests:

```bash
kalco validate \
  --base origin/${{ github.base_ref }} \
  --dirs manifests \
  --snapshot-dir tmp-snapshot \
  --report-file report.md
```

This single command replaces complex bash scripting loops.

### Local Validation

You can also run it locally to verify your changes before pushing:

```bash
# Validate changes in 'charts/' against 'main'
kalco validate \
  --base origin/main \
  --dirs charts \
  --snapshot-dir ./cluster-backup
```

## How it works (Under the hood)

1.  **Git Diff**: Runs `git diff --name-only <base> -- <dirs>` to find changed YAML files.
2.  **Loop**: Iterates through each file.
3.  **Load Base**: Runs `git show <base>:<file>` to get the original content.
4.  **Kalco Diff**: Runs `kalco diff` using the local file, the fetched base content (as `--ignore-base`), and the snapshot.
5.  **Report**: Aggregates the results into a single output.

If **Drift** is detected (changes that are *not* in your PR intent), the command exits with code 1.
