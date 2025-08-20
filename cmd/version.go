package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// These will be set by the build process
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: formatLongDescription(`
Display version information for kalco including build details,
Go version, and platform information.
`),

	RunE: func(cmd *cobra.Command, args []string) error {
		return runVersion()
	},
}

var versionDetailed bool

func runVersion() error {
	if versionDetailed {
		printHeader("Kalco Version Information")

		fmt.Printf("Version:      %s\n", version)
		fmt.Printf("Git Commit:   %s\n", commit)
		fmt.Printf("Build Date:   %s\n", date)
		fmt.Printf("Go Version:   %s\n", runtime.Version())
		fmt.Printf("OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Compiler:     %s\n", runtime.Compiler)

		return nil
	}

	// Simple version output
	fmt.Printf("kalco version %s\n", version)
	if commit != "unknown" && len(commit) >= 7 {
		fmt.Printf("Git commit: %s\n", commit[:7])
	}
	if date != "unknown" {
		fmt.Printf("Built: %s\n", date)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVar(&versionDetailed, "detailed", false, "show detailed version information")
}
