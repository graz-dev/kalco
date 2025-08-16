---
layout: default
title: Contributing
description: "How to contribute to Kalco development"
---

# Contributing to Kalco

We welcome contributions to Kalco! This guide will help you get started with contributing to the project.

## 🤝 Ways to Contribute

- 🐛 **Report Bugs** - Help us identify and fix issues
- 💡 **Request Features** - Suggest new functionality
- 📖 **Improve Documentation** - Help make our docs better
- 🔧 **Submit Code** - Contribute bug fixes and new features
- 🧪 **Write Tests** - Help improve our test coverage
- 🎨 **Improve UX** - Enhance the user experience

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **Git** - For version control
- **Make** - For build automation
- **Kubernetes cluster** - For testing (KIND, minikube, or real cluster)

### Development Setup

1. **Fork the repository** on GitHub

2. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/kalco.git
   cd kalco
   ```

3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/graz-dev/kalco.git
   ```

4. **Install dependencies**:
   ```bash
   go mod tidy
   ```

5. **Build the project**:
   ```bash
   make build
   ```

6. **Run tests**:
   ```bash
   make test
   ```

7. **Verify everything works**:
   ```bash
   ./kalco --help
   ```

## 🔧 Development Workflow

### Creating a Feature Branch

```bash
# Sync with upstream
git fetch upstream
git checkout master
git merge upstream/master

# Create feature branch
git checkout -b feature/your-feature-name
```

### Making Changes

1. **Write code** following our coding standards
2. **Add tests** for new functionality
3. **Update documentation** if needed
4. **Test your changes** thoroughly

### Testing Your Changes

```bash
# Run all tests
make test

# Run specific tests
go test ./pkg/dumper/...

# Run with coverage
make test-coverage

# Build and test locally
make build
./kalco export --dry-run
```

### Committing Changes

We follow conventional commit messages:

```bash
# Format: type(scope): description
git commit -m "feat(export): add namespace filtering support"
git commit -m "fix(validate): handle missing configmap references"
git commit -m "docs(readme): update installation instructions"
```

**Commit Types:**
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Test additions or changes
- `chore`: Build process or auxiliary tool changes

### Submitting a Pull Request

1. **Push your branch**:
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create a Pull Request** on GitHub with:
   - Clear title and description
   - Reference to related issues
   - Screenshots (if UI changes)
   - Testing instructions

3. **Address review feedback** if needed

4. **Celebrate** when your PR is merged! 🎉

## 📋 Coding Standards

### Go Code Style

We follow standard Go conventions:

- **gofmt** - Format code with `go fmt`
- **golint** - Follow linting recommendations
- **go vet** - Check for common mistakes
- **Effective Go** - Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines

```bash
# Format code
go fmt ./...

# Run linter
make lint

# Check for issues
go vet ./...
```

### Code Organization

```
kalco/
├── cmd/                 # CLI commands
│   ├── root.go         # Root command
│   ├── export.go       # Export command
│   └── ...
├── pkg/                # Core packages
│   ├── dumper/         # Resource dumping logic
│   ├── validation/     # Validation logic
│   └── ...
├── docs/               # Documentation
├── examples/           # Example scripts
└── scripts/            # Build and install scripts
```

### Naming Conventions

- **Packages**: lowercase, single word when possible
- **Functions**: camelCase, exported functions start with uppercase
- **Variables**: camelCase, descriptive names
- **Constants**: UPPER_CASE or camelCase for unexported

### Error Handling

```go
// Good: Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to export resources: %w", err)
}

// Good: Use specific error types when appropriate
type ValidationError struct {
    Resource string
    Message  string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Resource, e.Message)
}
```

### Testing

- **Unit tests** for all public functions
- **Integration tests** for complex workflows
- **Table-driven tests** for multiple test cases
- **Mocks** for external dependencies

```go
func TestExportResources(t *testing.T) {
    tests := []struct {
        name     string
        input    ExportConfig
        expected int
        wantErr  bool
    }{
        {
            name:     "basic export",
            input:    ExportConfig{OutputDir: "/tmp/test"},
            expected: 10,
            wantErr:  false,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ExportResources(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ExportResources() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if result != tt.expected {
                t.Errorf("ExportResources() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## 📖 Documentation Standards

### Code Documentation

- **Package comments** for all packages
- **Function comments** for exported functions
- **Example code** in documentation
- **Godoc format** for API documentation

```go
// Package dumper provides functionality for exporting Kubernetes resources
// to organized YAML files with clean metadata suitable for re-application.
package dumper

// ExportResources exports all discoverable resources from the cluster
// to the specified output directory. It returns the number of resources
// exported and any error encountered.
//
// Example:
//   count, err := ExportResources(ExportConfig{
//       OutputDir: "./backup",
//       Namespaces: []string{"default", "production"},
//   })
func ExportResources(config ExportConfig) (int, error) {
    // Implementation...
}
```

### User Documentation

- **Clear examples** for all features
- **Step-by-step guides** for common tasks
- **Troubleshooting sections** for known issues
- **Screenshots** for UI features

## 🧪 Testing Guidelines

### Test Categories

1. **Unit Tests** - Test individual functions
2. **Integration Tests** - Test component interactions
3. **End-to-End Tests** - Test complete workflows
4. **Performance Tests** - Test scalability and performance

### Running Tests

```bash
# All tests
make test

# Specific package
go test ./pkg/dumper/

# With coverage
make test-coverage

# Verbose output
go test -v ./...

# Race condition detection
go test -race ./...
```

### Test Data

- Use **testdata** directories for test files
- Create **minimal test cases** that cover edge cases
- Use **table-driven tests** for multiple scenarios
- **Mock external dependencies** (Kubernetes API, Git, etc.)

## 🐛 Bug Reports

### Before Reporting

1. **Search existing issues** to avoid duplicates
2. **Test with latest version** to ensure bug still exists
3. **Gather relevant information** (version, OS, Kubernetes version)

### Bug Report Template

```markdown
**Bug Description**
A clear description of what the bug is.

**Steps to Reproduce**
1. Run command: `kalco export --namespaces production`
2. See error: ...

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**
- Kalco version: v1.0.0
- OS: macOS 13.0
- Kubernetes version: v1.28.0
- Go version: 1.21.0

**Additional Context**
Any other context about the problem.
```

## 💡 Feature Requests

### Before Requesting

1. **Check existing issues** and discussions
2. **Consider the scope** - does it fit Kalco's mission?
3. **Think about implementation** - how would it work?

### Feature Request Template

```markdown
**Feature Description**
A clear description of the feature you'd like to see.

**Use Case**
Describe the problem this feature would solve.

**Proposed Solution**
How you envision this feature working.

**Alternatives Considered**
Other approaches you've considered.

**Additional Context**
Any other context or screenshots.
```

## 🏗️ Architecture Guidelines

### Design Principles

1. **Modularity** - Keep components loosely coupled
2. **Testability** - Design for easy testing
3. **Performance** - Optimize for large clusters
4. **Usability** - Prioritize user experience
5. **Reliability** - Handle errors gracefully

### Adding New Commands

1. **Create command file** in `cmd/` directory
2. **Implement command logic** in appropriate `pkg/` package
3. **Add comprehensive tests**
4. **Update documentation**
5. **Add examples** and use cases

### Adding New Features

1. **Design the API** - How will users interact with it?
2. **Plan the implementation** - Which packages are affected?
3. **Consider backwards compatibility** - Will it break existing usage?
4. **Write tests first** - Test-driven development
5. **Document thoroughly** - Code comments and user docs

## 📦 Release Process

### Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** (v2.0.0) - Breaking changes
- **MINOR** (v1.1.0) - New features, backwards compatible
- **PATCH** (v1.0.1) - Bug fixes, backwards compatible

### Release Checklist

1. **Update version** in relevant files
2. **Update CHANGELOG** with new features and fixes
3. **Run full test suite**
4. **Update documentation**
5. **Create release tag**
6. **Publish release** with release notes

## 🎯 Areas for Contribution

### High Priority

- 🔧 **Performance optimization** for large clusters
- 🧪 **Test coverage improvement**
- 📖 **Documentation enhancement**
- 🐛 **Bug fixes** and stability improvements

### Medium Priority

- 🎨 **CLI UX improvements**
- 📊 **New analysis features**
- 🔌 **Integration with other tools**
- 🌐 **Internationalization**

### Ideas for New Contributors

- 📝 **Fix typos** in documentation
- 🧪 **Add test cases** for existing functionality
- 📖 **Improve examples** and tutorials
- 🐛 **Reproduce and fix** reported bugs
- 💡 **Implement small feature requests**

## 📞 Getting Help

### Development Questions

- **GitHub Discussions** - [Ask questions](https://github.com/graz-dev/kalco/discussions)
- **Code Review** - Request feedback on your approach
- **Architecture Decisions** - Discuss design choices

### Communication Channels

- **GitHub Issues** - Bug reports and feature requests
- **GitHub Discussions** - General questions and ideas
- **Pull Request Comments** - Code-specific discussions

## 🙏 Recognition

We appreciate all contributions! Contributors are recognized in:

- **CONTRIBUTORS.md** file
- **Release notes** for significant contributions
- **GitHub contributors** page
- **Special mentions** in documentation

## 📄 License

By contributing to Kalco, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Kalco! Together, we're making Kubernetes cluster management better for everyone. 🚀

[← FAQ](faq.md) | [Documentation Home →](index.md)