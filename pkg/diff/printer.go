package diff

import (
	"fmt"
	"strings"
)

// PrintDiff outputs the diff result
func PrintDiff(result *DiffResult, format string) {
	if format == "markdown" {
		printMarkdown(result)
	} else {
		printText(result)
	}
}

func printText(result *DiffResult) {
	if result.IsNew {
		fmt.Println("\033[32m🟢 NEW RESOURCE\033[0m")
		return
	}

	if result.IsStale {
		fmt.Println("\033[33m⚠️  Snapshot is stale (>24h). Run with --refresh for live check.\033[0m")
	}

	if !result.HasDrift {
		fmt.Println("\033[32m✅ No changes detected\033[0m")
		return
	}

	// Print raw diff (go-cmp output is readable)
	// Colorize it slightly
	lines := strings.Split(result.Diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "+") {
			fmt.Printf("\033[32m%s\033[0m\n", line)
		} else if strings.HasPrefix(line, "-") {
			fmt.Printf("\033[31m%s\033[0m\n", line)
		} else {
			fmt.Println(line)
		}
	}
}

func printMarkdown(result *DiffResult) {
	if result.IsNew {
		fmt.Println("### Status: 🟢 NEW RESOURCE")
		return
	}
	
	if result.IsStale {
		fmt.Println("> ⚠️ **Warning**: Snapshot is stale (>24h). Run with `--refresh` for live check.")
		fmt.Println("")
	}

	if !result.HasDrift {
		fmt.Println("### Status: ✅ In Sync")
		return
	}

	fmt.Println("### Status: 🔴 Drift Detected")
	fmt.Println("")
	fmt.Println("```diff")
	fmt.Println(result.Diff)
	fmt.Println("```")
}
