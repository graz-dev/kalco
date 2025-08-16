---
layout: default
title: Installation
nav_order: 1
parent: Getting Started
---

# Installation

Get Kalco up and running on your system with these simple installation methods.

## 🚀 Quick Install

The fastest way to get started is using Go's built-in installer:

```bash
go install github.com/graz-dev/kalco/cmd/kalco@latest
```

## 📦 Package Managers

### Homebrew (macOS/Linux)

```bash
# Add the tap (when available)
brew tap graz-dev/kalco

# Install Kalco
brew install kalco
```

### Manual Download

Download the latest release for your platform from the [releases page](https://github.com/graz-dev/kalco/releases).

## 🔨 Build from Source

If you prefer to build from source:

```bash
# Clone the repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Build the binary
go build -o kalco ./cmd

# Make it executable
chmod +x kalco

# Move to your PATH
sudo mv kalco /usr/local/bin/
```

## ✅ Verify Installation

Check that Kalco is properly installed:

```bash
kalco --version
```

You should see output similar to:
```
kalco version v0.1.0 (commit: abc1234, date: 2025-08-16)
```

## 🔧 Shell Completion

### Bash

```bash
# Add to your ~/.bashrc
echo 'source <(kalco completion bash)' >> ~/.bashrc
source ~/.bashrc
```

### Zsh

```bash
# Add to your ~/.zshrc
echo 'source <(kalco completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

### Fish

```bash
kalco completion fish > ~/.config/fish/completions/kalco.fish
```

## 🐛 Troubleshooting

### Common Issues

**Command not found**: Ensure Kalco is in your PATH
```bash
echo $PATH
which kalco
```

**Permission denied**: Make the binary executable
```bash
chmod +x kalco
```

**Go version issues**: Ensure you have Go 1.19+ installed
```bash
go version
```

## 📚 Next Steps

Once Kalco is installed, proceed to [First Run]({{ site.baseurl }}/docs/getting-started/first-run) to export your first cluster.
