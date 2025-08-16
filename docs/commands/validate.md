---
layout: default
title: kalco validate
parent: Commands Reference
nav_order: 2
description: "Validate cluster resources for issues and broken references"
---

# kalco validate

Validate cluster resources for issues and broken references.

## Synopsis

The `validate` command performs comprehensive validation of your Kubernetes cluster resources, checking for broken references, configuration issues, and potential problems that could cause issues during re-application.

```bash
kalco validate [flags]
```

## Description

The validate command performs several types of validation:

- **Cross-reference validation** - Checks for broken references between resources
- **Configuration validation** - Validates resource configurations against schemas
- **Dependency analysis** - Analyzes resource dependencies and relationships
- **Security validation** - Checks for common security misconfigurations

## Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--output` | `-o` | string | Output format (table, json, yaml) (default: "table") |
| `--namespaces` | `-n` | []string | Specific namespaces to validate |
| `--resources` | `-r` | []string | Specific resource types to validate |
| `--fix` | | bool | Attempt to fix validation issues where possible |

## Examples

### Basic Validation

```bash
# Validate entire cluster
kalco validate

# Validate specific namespaces
kalco validate --namespaces production,staging

# Validate specific resource types
kalco validate --resources deployments,services,configmaps
```

### Output Formats

```bash
# Human-readable table (default)
kalco validate

# JSON output for automation
kalco validate --output json

# YAML output for configuration
kalco validate --output yaml
```

### Advanced Usage

```bash
# Validate and attempt fixes
kalco validate --fix

# Validate production namespace only
kalco validate --namespaces production --output json > validation-report.json

# Validate core workload resources
kalco validate --resources "deployments,services,configmaps,secrets"
```

## Validation Types

### Cross-Reference Validation

Checks for broken references between resources:

- Service selectors pointing to non-existent pods
- ConfigMap/Secret references from pods and deployments
- PVC references from pods
- ServiceAccount references in role bindings

### Configuration Validation

Validates resource configurations:

- Required fields and proper syntax
- Resource limits and requests
- Security contexts and policies
- Network policies and ingress rules

### Dependency Analysis

Analyzes resource relationships:

- Orphaned resources without owners
- Circular dependencies
- Missing dependencies
- Unused resources

## Output Format

### Table Output (Default)

```
RESOURCE TYPE    NAME           NAMESPACE    STATUS    ISSUES
Deployment       app-deploy     default      ✓ Valid   0
Service          app-service    default      ✗ Error   1
ConfigMap        app-config     default      ⚠ Warning 1
```

### JSON Output

```json
{
  "summary": {
    "totalReferences": 45,
    "validReferences": 42,
    "brokenReferences": 2,
    "warningReferences": 1
  },
  "brokenReferences": [
    {
      "sourceType": "Service",
      "sourceName": "app-service",
      "targetType": "Pod",
      "targetName": "missing-pod",
      "field": "spec.selector.app"
    }
  ]
}
```

## Exit Codes

- `0` - Validation passed with no issues
- `1` - Validation found warnings but no errors
- `2` - Validation found errors
- `3` - Validation failed to run

## Related Commands

- [`kalco export`](export.md) - Export resources for validation
- [`kalco analyze orphaned`](analyze.md#orphaned) - Find orphaned resources
- [`kalco report`](report.md) - Generate comprehensive reports

---

[← Commands Overview](index.md) | [Analyze Command →](analyze.md)