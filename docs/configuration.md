---
layout: default
title: Configuration
---

# Configuration

Kalco supports flexible configuration through YAML files, environment variables, and command-line flags.

## Configuration Hierarchy

Configuration is loaded in the following order (later sources override earlier ones):

1. **Default values** - Built-in defaults
2. **Global configuration** - `~/.kalco/config.yaml`
3. **Project configuration** - `./.kalco.yaml`
4. **Environment variables** - `KALCO_*` prefixed variables
5. **Command-line flags** - Explicit flag values

## Configuration Files

### Global Configuration

Located at `~/.kalco/config.yaml`, this file contains user-wide defaults.

```bash
# Initialize global configuration
kalco config init --global
```

### Project Configuration

Located at `./.kalco.yaml` in your project directory, this file contains project-specific settings.

```bash
# Initialize project configuration
kalco config init

# Initialize with advanced template
kalco config init --template advanced
```

## Configuration Schema

### Basic Configuration

```yaml
# Basic Kalco Configuration
output:
  directory: "./kalco-export-{{.Date}}"
  format: "yaml"

filters:
  exclude: ["events"]

ui:
  colors: true
  verbose: false
```

### Advanced Configuration

```yaml
# Advanced Kalco Configuration
output:
  directory: "./kalco-export-{{.Date}}"
  format: "yaml"
  git:
    enabled: true
    auto_push: false
    commit_message: "Kalco export {{.Date}}"

filters:
  namespaces: []
  resources: []
  exclude: ["events", "replicasets"]

validation:
  enabled: true
  strict: false
  
analysis:
  orphaned_resources: true
  security_scan: false
  
reports:
  enabled: true
  formats: ["html", "json"]
  
ui:
  colors: true
  verbose: false
  progress: true

# Kubernetes connection
kubernetes:
  kubeconfig: ""
  context: ""
  namespace: ""

# Custom resource definitions
crds:
  include_all: true
  specific: []
```

## Configuration Sections

### Output Configuration

Controls where and how resources are exported.

```yaml
output:
  directory: "./exports/{{.Date}}"    # Output directory template
  format: "yaml"                      # Output format (yaml, json)
  compression: false                  # Enable gzip compression
  git:
    enabled: true                     # Enable Git integration
    auto_push: false                  # Automatically push to remote
    commit_message: "Export {{.Date}}" # Commit message template
    remote: "origin"                  # Git remote name
```

**Template Variables:**
- `{{.Date}}` - Current date (YYYY-MM-DD)
- `{{.Time}}` - Current time (HH-MM-SS)
- `{{.Timestamp}}` - Unix timestamp
- `{{.Context}}` - Kubernetes context name

### Filter Configuration

Controls which resources are included or excluded.

```yaml
filters:
  namespaces: ["production", "staging"]  # Include specific namespaces
  exclude_namespaces: ["kube-system"]    # Exclude specific namespaces
  resources: ["pods", "services"]        # Include specific resource types
  exclude: ["events", "endpoints"]       # Exclude specific resource types
  labels:                               # Filter by labels
    app: "myapp"
    environment: "production"
  annotations:                          # Filter by annotations
    "deployment.kubernetes.io/revision": "1"
```

### Validation Configuration

Controls resource validation behavior.

```yaml
validation:
  enabled: true                    # Enable validation
  strict: false                   # Strict validation mode
  cross_references: true          # Check cross-references
  schemas: true                   # Validate against schemas
  custom_rules: []                # Custom validation rules
  ignore_warnings: false          # Ignore validation warnings
```

### Analysis Configuration

Controls analysis and reporting features.

```yaml
analysis:
  orphaned_resources: true        # Detect orphaned resources
  security_scan: false           # Run security analysis
  resource_usage: true           # Analyze resource usage
  dependencies: true             # Analyze dependencies
  recommendations: true          # Generate recommendations
```

### UI Configuration

Controls command-line interface appearance and behavior.

```yaml
ui:
  colors: true                   # Enable colored output
  verbose: false                 # Enable verbose output
  progress: true                 # Show progress indicators
  interactive: true              # Enable interactive prompts
  pager: "less"                  # Pager for long output
```

## Environment Variables

All configuration options can be set via environment variables using the `KALCO_` prefix:

```bash
# Output configuration
export KALCO_OUTPUT_DIRECTORY="./my-exports"
export KALCO_OUTPUT_FORMAT="json"

# Filter configuration
export KALCO_FILTERS_EXCLUDE="events,endpoints"
export KALCO_FILTERS_NAMESPACES="production,staging"

# UI configuration
export KALCO_UI_COLORS="false"
export KALCO_UI_VERBOSE="true"

# Kubernetes configuration
export KALCO_KUBERNETES_KUBECONFIG="/path/to/kubeconfig"
export KALCO_KUBERNETES_CONTEXT="production"
```

## Command-Line Configuration

### Setting Configuration Values

```bash
# Set output directory
kalco config set output.directory ./backups

# Set global configuration
kalco config set --global ui.colors false

# Set filter exclusions
kalco config set filters.exclude "events,endpoints,replicasets"

# Set Git integration
kalco config set output.git.enabled true
kalco config set output.git.auto_push true
```

### Viewing Configuration

```bash
# Show current configuration
kalco config show

# Show configuration in JSON format
kalco config show --output json

# Show specific configuration section
kalco config show --section output
```

## Configuration Templates

### Minimal Template

```yaml
output:
  directory: "./kalco-export"
ui:
  colors: true
```

### Development Template

```yaml
output:
  directory: "./dev-exports/{{.Date}}"
  git:
    enabled: true
    commit_message: "Dev export {{.Date}}"

filters:
  namespaces: ["default", "development"]
  exclude: ["events", "endpoints"]

ui:
  verbose: true
  progress: true
```

### Production Template

```yaml
output:
  directory: "/backups/kubernetes/{{.Date}}"
  compression: true
  git:
    enabled: true
    auto_push: true
    commit_message: "Production backup {{.Date}}"

filters:
  exclude: ["events", "endpoints", "replicasets"]

validation:
  enabled: true
  strict: true

analysis:
  orphaned_resources: true
  security_scan: true

reports:
  enabled: true
  formats: ["html", "json"]
```

### CI/CD Template

```yaml
output:
  directory: "./ci-exports"
  format: "json"

filters:
  exclude: ["events", "endpoints", "pods"]

validation:
  enabled: true
  strict: true

ui:
  colors: false
  verbose: true
  interactive: false
```

## Configuration Validation

Validate your configuration file:

```bash
# Validate current configuration
kalco config validate

# Validate specific configuration file
kalco config validate --file .kalco.yaml

# Show configuration with resolved values
kalco config show --resolved
```

## Best Practices

### Security

- **Sensitive Data**: Don't store sensitive information in configuration files
- **Permissions**: Set appropriate file permissions (600) for configuration files
- **Git**: Add `.kalco.yaml` to `.gitignore` if it contains sensitive data

### Organization

- **Global Defaults**: Use global configuration for user preferences
- **Project Specific**: Use project configuration for project-specific settings
- **Environment Variables**: Use environment variables in CI/CD environments

### Maintenance

- **Version Control**: Track configuration changes in version control
- **Documentation**: Document custom configuration options
- **Validation**: Regularly validate configuration files

## Troubleshooting

### Common Issues

**Configuration Not Loading:**
```bash
# Check configuration file syntax
kalco config validate --file .kalco.yaml

# Show effective configuration
kalco config show --resolved
```

**Permission Errors:**
```bash
# Check file permissions
ls -la ~/.kalco/config.yaml

# Fix permissions
chmod 600 ~/.kalco/config.yaml
```

**Template Errors:**
```bash
# Test template rendering
kalco config show --section output
```

---

[← Commands Reference](commands/index.md) | [Use Cases →](use-cases.md)