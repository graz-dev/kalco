# Kalco Examples

This directory contains practical examples and scripts demonstrating how to use Kalco effectively.

## 📁 Available Examples

### 🚀 Quick Start
- **[quickstart.sh](quickstart.sh)** - Complete workflow from cluster creation to export and analysis

### 🌐 Context Management
- **[context-example.sh](context-example.sh)** - Demonstrates context management for multiple clusters

## 🎯 Quick Start Example

The `quickstart.sh` script provides a complete end-to-end example:

```bash
# Make executable and run
chmod +x examples/quickstart.sh
./examples/quickstart.sh
```

This example:
1. Creates a Kind cluster
2. Deploys a sample application
3. Runs Kalco export
4. Makes changes to the cluster
5. Runs Kalco again to see changes
6. Cleans up the cluster

## 🔄 Context Management Example

The `context-example.sh` script demonstrates context management:

```bash
# Make executable and run
chmod +x examples/context-example.sh
./examples/context-example.sh
```

This example:
1. Creates multiple contexts for different environments
2. Shows context switching
3. Demonstrates automatic context usage in export
4. Shows context override with flags
5. Cleans up example contexts

## 🛠️ Prerequisites

Before running the examples, ensure you have:

- **Kalco** installed and in your PATH
- **Docker** running (for Kind clusters)
- **Kind** installed (for local clusters)
- **kubectl** configured

## 📚 Learning Path

1. **Start with quickstart.sh** - Learn basic Kalco operations
2. **Try context-example.sh** - Master context management
3. **Experiment with your own clusters** - Apply concepts to real scenarios

## 🔧 Customization

Feel free to modify these examples:

- Change cluster names and configurations
- Add your own applications and resources
- Modify export parameters and filters
- Customize context labels and descriptions

## 🐛 Troubleshooting

### Common Issues

**Permission denied**
```bash
chmod +x examples/*.sh
```

**Command not found**
```bash
# Ensure Kalco is in your PATH
which kalco
```

**Cluster connection issues**
```bash
# Check cluster status
kubectl cluster-info
kubectl get nodes
```

## 📖 Next Steps

After running the examples:

1. **Read the documentation** - Explore the full command reference
2. **Try different clusters** - Test with your own Kubernetes clusters
3. **Customize contexts** - Create contexts for your specific use cases
4. **Integrate with CI/CD** - Automate exports in your pipelines

## 🤝 Contributing

Have a great example? Feel free to contribute:

1. Create a new script with descriptive name
2. Add clear comments and documentation
3. Include error handling and cleanup
4. Update this README with your example

## 📚 Resources

- [Kalco Documentation](../docs/)
- [Commands Reference](../docs/commands/)
- [Getting Started](../docs/getting-started/)
