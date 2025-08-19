# Kalco Examples

This directory contains practical examples demonstrating how to use Kalco for Kubernetes cluster management and analysis.

## Available Examples

### Quickstart Script (`quickstart.sh`)

A comprehensive demonstration of Kalco's core functionality:

- **Context Management**: Creating and managing cluster contexts
- **Resource Export**: Simulating cluster resource export and organization
- **Git Integration**: Demonstrating version control capabilities
- **Report Generation**: Showing how change reports are structured

**Usage:**
```bash
# Run the complete quickstart demo
./examples/quickstart.sh

# Keep the demo directory for inspection
./examples/quickstart.sh --keep
```

**What you'll learn:**
- How to create and manage contexts with `kalco context`
- How contexts store cluster configuration and output directories
- How Kalco organizes exported resources in a structured way
- How Git integration works for version control
- How `kalco-config.json` stores context information

## Prerequisites

Before running the example, ensure you have:

1. **Kalco installed** - Download from [releases](https://github.com/graz-dev/kalco/releases) or build from source
2. **Git available** - For version control functionality
3. **Basic Kubernetes knowledge** - Understanding of clusters, namespaces, and resources

## Running the Example

### Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/graz-dev/kalco.git
   cd kalco
   ```

2. **Install Kalco:**
   ```bash
   # Quick install
   curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
   
   # Or build from source
   go mod tidy
   go build -o kalco
   ```

3. **Run the example:**
   ```bash
   # Make script executable
   chmod +x examples/quickstart.sh
   
   # Run quickstart
   ./examples/quickstart.sh
   ```

### Example Output

The example will show you:

- **Context Creation**: How to set up contexts for different environments
- **Resource Organization**: How Kalco structures exported resources
- **Git Integration**: How version control is automatically handled
- **Configuration Management**: How contexts store cluster settings
- **Report Generation**: How change tracking and validation work

## Customizing the Example

You can modify this example to:

- **Use Real Clusters**: Replace simulated resources with actual Kubernetes clusters
- **Add Custom Resources**: Include your organization's specific resource types
- **Modify Contexts**: Adjust context configurations for your environment
- **Extend Workflows**: Add additional steps specific to your use case

## Next Steps

After running the example:

1. **Try with Real Clusters**: Use `kalco export` with your actual Kubernetes clusters
2. **Explore Reports**: Examine the generated reports in `kalco-reports/` directories
3. **Share Contexts**: Use `kalco context load` to import configurations from team members
4. **Read Documentation**: Visit [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
5. **Join Community**: Participate in [GitHub discussions](https://github.com/graz-dev/kalco/discussions)

## Troubleshooting

### Common Issues

- **Permission Denied**: Ensure script is executable with `chmod +x examples/quickstart.sh`
- **Kalco Not Found**: Verify Kalco is installed and in your PATH
- **Git Not Available**: Install Git for version control functionality
- **Directory Issues**: Ensure you have write permissions in the current directory

### Getting Help

- **Documentation**: [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
- **Issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Discussions**: [GitHub Discussions](https://github.com/graz-dev/kalco/discussions)

## Contributing

We welcome contributions to improve this example:

- **Bug Reports**: Report issues with the example
- **Enhancements**: Suggest improvements
- **Documentation**: Help improve example descriptions and usage
- **New Examples**: Submit examples for additional use cases

For more information, see the [main README](../README.md) and [contributing guidelines](../CONTRIBUTING.md).
