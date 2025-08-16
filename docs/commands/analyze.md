---
layout: default
title: kalco analyze
parent: Commands Reference
nav_order: 3
description: "Analyze cluster resources for optimization opportunities"
---

# kalco analyze

Analyze cluster resources for optimization opportunities.

## Synopsis

The `analyze` command provides various analysis capabilities to help you optimize your Kubernetes cluster, identify cleanup opportunities, and improve security posture.

```bash
kalco analyze <subcommand> [flags]
```

## Subcommands

### orphaned

Find orphaned resources no longer managed by controllers.

```bash
kalco analyze orphaned [flags]
```

**Description:**
Identifies resources that are no longer managed by higher-level controllers and may be safe to clean up:

- Pods not owned by ReplicaSets, Deployments, or Jobs
- ReplicaSets not owned by Deployments  
- ConfigMaps and Secrets not referenced by any resources
- Services without matching endpoints
- PersistentVolumes not bound to claims

**Flags:**
- `--namespaces, -n` - Specific namespaces to analyze
- `--output, -o` - Output format (table, json, yaml)
- `--detailed` - Include detailed analysis information

**Examples:**
```bash
# Find all orphaned resources
kalco analyze orphaned

# Analyze specific namespaces
kalco analyze orphaned --namespaces production,staging

# Detailed analysis with JSON output
kalco analyze orphaned --detailed --output json
```

### usage

Analyze resource usage and capacity.

```bash
kalco analyze usage [flags]
```

**Description:**
Analyzes cluster resource usage, capacity, and efficiency metrics including CPU, memory, and storage utilization across nodes, namespaces, and workloads.

**Examples:**
```bash
# Analyze overall cluster usage
kalco analyze usage

# Analyze usage by namespace
kalco analyze usage --by-namespace

# Analyze node capacity
kalco analyze usage --nodes
```

### security

Analyze cluster security posture.

```bash
kalco analyze security [flags]
```

**Description:**
Analyzes your cluster's security configuration and identifies potential security issues or improvements. Checks for common security misconfigurations and compliance with security best practices.

**Examples:**
```bash
# Run security analysis
kalco analyze security

# Check specific security policies
kalco analyze security --policies rbac,network,pod-security
```

## Global Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--output` | `-o` | string | Output format (table, json, yaml) |
| `--namespaces` | `-n` | []string | Specific namespaces to analyze |
| `--detailed` | | bool | Include detailed analysis information |

## Output Examples

### Orphaned Resources (Table)

```
TYPE                    NAME              NAMESPACE    REASON              
Pod                     old-pod-123       default      No Controller Owner
ConfigMap               unused-config     production   No References      
Secret                  old-secret        staging      No References      
Service                 orphaned-svc      default      No References      
```

### Orphaned Resources (JSON)

```json
{
  "summary": {
    "totalOrphanedResources": 15,
    "byType": {
      "Pod": 5,
      "ConfigMap": 4,
      "Secret": 3,
      "Service": 2,
      "PVC": 1
    }
  },
  "orphanedResources": [
    {
      "type": "Pod",
      "name": "old-pod-123",
      "namespace": "default",
      "reason": "No Controller Owner",
      "details": "This Pod has no owner references and may be orphaned"
    }
  ]
}
```

## Use Cases

### Cleanup Operations

```bash
# Find cleanup opportunities
kalco analyze orphaned --output json > cleanup-candidates.json

# Review and create cleanup script
cat cleanup-candidates.json | jq -r '.orphanedResources[] | 
  "kubectl delete \(.type) \(.name) -n \(.namespace)"' > cleanup.sh
```

### Capacity Planning

```bash
# Analyze resource usage for capacity planning
kalco analyze usage --output json > usage-report.json

# Generate capacity planning report
kalco analyze usage --detailed --output html > capacity-report.html
```

### Security Auditing

```bash
# Security assessment
kalco analyze security --output json > security-audit.json

# Check specific security areas
kalco analyze security --policies rbac,network --detailed
```

## Related Commands

- [`kalco validate`](validate.md) - Validate cluster resources
- [`kalco export`](export.md) - Export resources for analysis
- [`kalco report`](report.md) - Generate comprehensive reports

---

[← Validate Command](validate.md) | [Report Command →](report.md)