---
layout: default
title: Use Cases
---

# Use Cases

Real-world scenarios and examples for using Kalco effectively.

## Disaster Recovery

### Complete Cluster Backup

Create comprehensive backups for disaster recovery scenarios.

```bash
# Daily backup with Git versioning
kalco export \
  --output ./backups/cluster-$(date +%Y%m%d) \
  --exclude events,endpoints,replicasets \
  --git-push \
  --commit-message "Daily backup - $(date)"

# Automated backup script
#!/bin/bash
BACKUP_DIR="/backups/kubernetes/$(date +%Y/%m)"
mkdir -p "$BACKUP_DIR"

kalco export \
  --output "$BACKUP_DIR/cluster-$(date +%Y%m%d-%H%M)" \
  --exclude events,endpoints \
  --commit-message "Automated backup $(date)"
```

### Environment Replication

Replicate production environments for testing or staging.

```bash
# Export production configuration
kalco export \
  --namespaces production \
  --exclude events,pods,replicasets \
  --output ./production-config

# Apply to staging (after review and modification)
kubectl apply -R -f ./production-config/production/
```

## Compliance & Auditing

### Security Auditing

Regular security assessments and compliance reporting.

```bash
# Generate security report
kalco analyze security --output json > security-audit-$(date +%Y%m%d).json

# Validate cluster configuration
kalco validate --output yaml > validation-report-$(date +%Y%m%d).yaml

# Export for compliance review
kalco export \
  --resources "roles,rolebindings,clusterroles,clusterrolebindings,networkpolicies,podsecuritypolicies" \
  --output ./compliance-review
```

### Change Tracking

Track and document cluster changes over time.

```bash
# Weekly change tracking
kalco export \
  --output ./cluster-history \
  --git-push \
  --commit-message "Weekly snapshot - $(date +%Y-W%U)"

# Generate change report
kalco report \
  --types changes \
  --since 7d \
  --output-file weekly-changes.html
```

## Development & Testing

### Local Development Setup

Set up local development environments that mirror production.

```bash
# Export development namespace
kalco export \
  --namespaces development \
  --exclude events,pods \
  --output ./dev-config

# Create KIND cluster with exported config
kind create cluster --name dev-cluster
kubectl apply -R -f ./dev-config/development/
```

### CI/CD Integration

Integrate Kalco into CI/CD pipelines for automated testing and deployment.

```yaml
# GitHub Actions example
name: Cluster Validation
on:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install Kalco
        run: |
          curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
      
      - name: Export and Validate Cluster
        run: |
          kalco export --output ./cluster-state
          kalco validate --output json > validation-results.json
          
      - name: Upload Results
        uses: actions/upload-artifact@v3
        with:
          name: cluster-validation
          path: |
            ./cluster-state
            validation-results.json
```

## Operations & Maintenance

### Resource Cleanup

Identify and clean up unused resources to optimize cluster efficiency.

```bash
# Find orphaned resources
kalco analyze orphaned --detailed --output json > orphaned-resources.json

# Generate cleanup script
cat orphaned-resources.json | jq -r '.orphanedResources[] | 
  "kubectl delete \(.type) \(.name) -n \(.namespace)"' > cleanup-script.sh

# Review and execute cleanup (after careful review!)
chmod +x cleanup-script.sh
# ./cleanup-script.sh  # Execute after review
```

### Capacity Planning

Analyze resource usage for capacity planning decisions.

```bash
# Analyze resource usage
kalco analyze usage --by-namespace --output json > usage-analysis.json

# Generate capacity report
kalco report \
  --types resources,usage \
  --output-file capacity-report.html
```

### Migration Planning

Plan and execute cluster migrations between environments or providers.

```bash
# Export source cluster
kalco export \
  --exclude events,pods,replicasets,endpoints \
  --output ./migration-source

# Validate configuration before migration
kalco validate --output json > pre-migration-validation.json

# Apply to target cluster (after modification)
# kubectl apply -R -f ./migration-source/
```

## Monitoring & Alerting

### Health Monitoring

Regular cluster health checks and alerting.

```bash
#!/bin/bash
# health-check.sh - Run daily health checks

# Validate cluster
kalco validate --output json > /tmp/validation.json

# Check for broken references
BROKEN_REFS=$(jq '.summary.brokenReferences' /tmp/validation.json)

if [ "$BROKEN_REFS" -gt 0 ]; then
  echo "ALERT: $BROKEN_REFS broken references found!"
  # Send alert (email, Slack, etc.)
fi

# Check for orphaned resources
kalco analyze orphaned --output json > /tmp/orphaned.json
ORPHANED_COUNT=$(jq '.summary.totalOrphanedResources' /tmp/orphaned.json)

if [ "$ORPHANED_COUNT" -gt 10 ]; then
  echo "WARNING: $ORPHANED_COUNT orphaned resources found"
  # Send notification
fi
```

### Prometheus Integration

Export metrics for Prometheus monitoring.

```bash
# Generate metrics
kalco analyze usage --output json | \
  jq -r '.metrics[] | "kalco_resource_count{\(.labels)} \(.value)"' > \
  /var/lib/prometheus/node-exporter/kalco-metrics.prom
```

## Documentation & Knowledge Sharing

### Cluster Documentation

Generate comprehensive cluster documentation.

```bash
# Generate complete cluster documentation
kalco report \
  --types summary,resources,security \
  --output-file cluster-documentation.html

# Export configuration for documentation
kalco export \
  --resources "configmaps,secrets" \
  --output ./config-documentation

# Generate resource inventory
kalco resources list --detailed --output json > resource-inventory.json
```

### Team Onboarding

Help new team members understand cluster structure.

```bash
# Create onboarding package
mkdir -p ./onboarding/cluster-overview

# Export key resources
kalco export \
  --resources "deployments,services,ingresses,configmaps" \
  --output ./onboarding/cluster-overview

# Generate overview report
kalco report \
  --types summary,resources \
  --output-file ./onboarding/cluster-overview.html

# List available resources
kalco resources list --output json > ./onboarding/available-resources.json
```

## Advanced Scenarios

### Multi-Cluster Management

Manage multiple clusters with centralized tooling.

```bash
#!/bin/bash
# multi-cluster-export.sh

CLUSTERS=("production" "staging" "development")
BASE_DIR="./multi-cluster-backup/$(date +%Y%m%d)"

for cluster in "${CLUSTERS[@]}"; do
  echo "Exporting cluster: $cluster"
  
  kalco export \
    --kubeconfig ~/.kube/config-$cluster \
    --output "$BASE_DIR/$cluster" \
    --exclude events,endpoints \
    --commit-message "Multi-cluster backup: $cluster"
done

# Generate cross-cluster report
kalco report \
  --input-dirs "$BASE_DIR/*" \
  --output-file "$BASE_DIR/multi-cluster-report.html"
```

### GitOps Integration

Integrate with GitOps workflows for declarative cluster management.

```bash
# Export to GitOps repository
kalco export \
  --output ./gitops-repo/clusters/production \
  --exclude events,pods,replicasets \
  --git-push \
  --commit-message "Production state update"

# Validate GitOps configuration
kalco validate \
  --input ./gitops-repo/clusters/production \
  --output json > gitops-validation.json
```

---

[← Configuration](configuration.md) | [API Reference →](api-reference.md)