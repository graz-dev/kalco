# Kalco Installation Script for Windows
# This script downloads and installs the latest version of kalco

param(
    [string]$InstallDir = "$env:LOCALAPPDATA\kalco\bin",
    [switch]$AddToPath = $true
)

# Configuration
$Repo = "graz-dev/kalco"
$BinaryName = "kalco.exe"

# Helper functions
function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White"
    )
    Write-Host $Message -ForegroundColor $Color
}

function Write-Info {
    param([string]$Message)
    Write-ColorOutput "ℹ️  $Message" "Blue"
}

function Write-Success {
    param([string]$Message)
    Write-ColorOutput "✅ $Message" "Green"
}

function Write-Warning {
    param([string]$Message)
    Write-ColorOutput "⚠️  $Message" "Yellow"
}

function Write-Error {
    param([string]$Message)
    Write-ColorOutput "❌ $Message" "Red"
}

# Detect architecture
function Get-Architecture {
    $arch = $env:PROCESSOR_ARCHITECTURE
    switch ($arch) {
        "AMD64" { return "x86_64" }
        "ARM64" { return "arm64" }
        default {
            Write-Error "Unsupported architecture: $arch"
            exit 1
        }
    }
}

# Get latest release version
function Get-LatestVersion {
    Write-Info "Fetching latest release information..."
    try {
        $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
        return $response.tag_name
    }
    catch {
        Write-Error "Failed to get latest version: $_"
        exit 1
    }
}

# Download and install kalco
function Install-Kalco {
    $version = Get-LatestVersion
    $arch = Get-Architecture
    $platform = "Windows_$arch"
    
    Write-Info "Detected platform: $platform"
    Write-Info "Latest version: $version"
    
    # Construct download URL
    $downloadUrl = "https://github.com/$Repo/releases/download/$version/kalco_$platform.zip"
    Write-Info "Downloading from: $downloadUrl"
    
    # Create temporary directory
    $tempDir = New-TemporaryFile | ForEach-Object { Remove-Item $_; New-Item -ItemType Directory -Path $_ }
    $archivePath = Join-Path $tempDir "kalco.zip"
    
    try {
        # Download the archive
        Invoke-WebRequest -Uri $downloadUrl -OutFile $archivePath
        
        # Extract the archive
        Write-Info "Extracting archive..."
        Expand-Archive -Path $archivePath -DestinationPath $tempDir -Force
        
        # Find the binary
        $binaryPath = Get-ChildItem -Path $tempDir -Name $BinaryName -Recurse | Select-Object -First 1
        if (-not $binaryPath) {
            Write-Error "Binary not found in archive"
            exit 1
        }
        
        $fullBinaryPath = Join-Path $tempDir $binaryPath
        
        # Create install directory
        if (-not (Test-Path $InstallDir)) {
            Write-Info "Creating install directory: $InstallDir"
            New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
        }
        
        # Install the binary
        $targetPath = Join-Path $InstallDir $BinaryName
        Write-Info "Installing kalco to $targetPath..."
        Copy-Item -Path $fullBinaryPath -Destination $targetPath -Force
        
        Write-Success "kalco $version installed successfully!"
        
        # Add to PATH if requested
        if ($AddToPath) {
            Add-ToPath $InstallDir
        }
        
        Write-Info "Run 'kalco --help' to get started"
    }
    finally {
        # Cleanup
        Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
    }
}

# Add directory to PATH
function Add-ToPath {
    param([string]$Directory)
    
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($currentPath -notlike "*$Directory*") {
        Write-Info "Adding $Directory to PATH..."
        $newPath = "$currentPath;$Directory"
        [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
        Write-Success "Added to PATH. Please restart your terminal or run: `$env:PATH += ';$Directory'"
    } else {
        Write-Info "$Directory is already in PATH"
    }
}

# Main installation process
function Main {
    Write-Host "🚀 Kalco Installation Script" -ForegroundColor Cyan
    Write-Host "============================" -ForegroundColor Cyan
    Write-Host ""
    
    # Check PowerShell version
    if ($PSVersionTable.PSVersion.Major -lt 5) {
        Write-Error "PowerShell 5.0 or higher is required"
        exit 1
    }
    
    Install-Kalco
    
    Write-Host ""
    Write-Success "Installation complete! 🎉"
    Write-Host ""
    Write-Host "Next steps:"
    Write-Host "  1. Run 'kalco --help' to see available commands"
    Write-Host "  2. Run 'kalco' to dump your current Kubernetes cluster"
    Write-Host "  3. Check out the documentation at: https://github.com/$Repo"
}

# Run main function
Main