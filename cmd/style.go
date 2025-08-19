package cmd

import (
	"fmt"
	"os"
	"strings"
)

// Color constants
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// colorize applies color to text if colors are enabled
func colorize(color, text string) string {
	if noColor {
		return text
	}
	return color + text + ColorReset
}

// printHeader prints a main section header
func printHeader(title string) {
	if noColor {
		fmt.Printf("\n%s\n", title)
		fmt.Println(strings.Repeat("=", len(title)))
	} else {
		fmt.Printf("\n%s %s\n", colorize(ColorCyan+ColorBold, title), colorize(ColorCyan, strings.Repeat("=", len(title))))
	}
}

// printSubHeader prints a subsection header
func printSubHeader(title string) {
	if noColor {
		fmt.Printf("\n%s\n", title)
		fmt.Println(strings.Repeat("-", len(title)))
	} else {
		fmt.Printf("\n%s %s\n", colorize(ColorBlue+ColorBold, title), colorize(ColorBlue, strings.Repeat("-", len(title))))
	}
}

// printCommandHeader prints a command execution header
func printCommandHeader(title, subtitle string) {
	if noColor {
		fmt.Printf("\n%s\n", title)
		if subtitle != "" {
			fmt.Printf("%s\n", subtitle)
		}
		fmt.Println(strings.Repeat("=", len(title)))
	} else {
		fmt.Printf("\n%s %s\n", colorize(ColorPurple+ColorBold, title), colorize(ColorPurple, strings.Repeat("=", len(title))))
		if subtitle != "" {
			fmt.Printf("%s\n", colorize(ColorCyan, subtitle))
		}
	}
}

// printInfo prints informational text
func printInfo(message string) {
	if noColor {
		fmt.Printf("INFO: %s\n", message)
	} else {
		fmt.Printf("%s %s\n", colorize(ColorBlue, "INFO:"), message)
	}
}

// printSuccess prints success text
func printSuccess(message string) {
	if noColor {
		fmt.Printf("SUCCESS: %s\n", message)
	} else {
		fmt.Printf("%s %s\n", colorize(ColorGreen, "SUCCESS:"), message)
	}
}

// printWarning prints warning text
func printWarning(message string) {
	if noColor {
		fmt.Printf("WARNING: %s\n", message)
	} else {
		fmt.Printf("%s %s\n", colorize(ColorYellow, "WARNING:"), message)
	}
}

// printError prints error text
func printError(message string) {
	if noColor {
		fmt.Printf("ERROR: %s\n", message)
	} else {
		fmt.Printf("%s %s\n", colorize(ColorRed, "ERROR:"), message)
	}
}

// printSeparator prints a visual separator
func printSeparator() {
	if noColor {
		fmt.Println("---")
	} else {
		fmt.Println(colorize(ColorCyan, "---"))
	}
}

// printBanner prints the kalco banner
func printBanner() {
	if noColor {
		fmt.Println("Kalco - Kubernetes Analysis & Lifecycle Control")
		fmt.Println("===============================================")
	} else {
		banner := `
██╗  ██╗ █████╗ ██╗      ██████╗ ██████╗ 
██║ ██╔╝██╔══██╗██║     ██╔════╝██╔═══██╗
█████╔╝ ███████║██║     ██║     ██║   ██║
██╔═██╗ ██╔══██║██║     ██║     ██║   ██║
██║  ██╗██║  ██║███████╗╚██████╗╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ 
                                           
Kubernetes Analysis & Lifecycle Control
Extract, validate, analyze, and version control your cluster`
		fmt.Println(colorize(ColorCyan+ColorBold, banner))
		fmt.Println()
	}
}

// printUsage prints usage information
func printUsage(cmd string, description string) {
	if noColor {
		fmt.Printf("Usage: %s\n", cmd)
		fmt.Printf("Description: %s\n", description)
	} else {
		fmt.Printf("Usage: %s\n", colorize(ColorGreen, cmd))
		fmt.Printf("Description: %s\n", colorize(ColorCyan, description))
	}
}

// printTableHeader prints a table header
func printTableHeader(headers ...string) {
	if noColor {
		fmt.Println(strings.Join(headers, " | "))
		fmt.Println(strings.Repeat("-", len(strings.Join(headers, " | "))))
	} else {
		fmt.Println(colorize(ColorBold, strings.Join(headers, " | ")))
		fmt.Println(colorize(ColorCyan, strings.Repeat("-", len(strings.Join(headers, " | ")))))
	}
}

// printTableRow prints a table row
func printTableRow(cells ...string) {
	fmt.Println(strings.Join(cells, " | "))
}

// printProgress prints a progress indicator
func printProgress(current, total int, message string) {
	percentage := float64(current) / float64(total) * 100
	if noColor {
		fmt.Printf("Progress: %d/%d (%.1f%%) - %s\n", current, total, percentage, message)
	} else {
		fmt.Printf("Progress: %s/%s (%.1f%%) - %s\n", 
			colorize(ColorGreen, fmt.Sprintf("%d", current)), 
			colorize(ColorCyan, fmt.Sprintf("%d", total)), 
			percentage, 
			message)
	}
}

// printStatus prints a status message with appropriate color
func printStatus(status, message string) {
	switch strings.ToLower(status) {
	case "success", "ok", "done":
		printSuccess(message)
	case "warning", "warn":
		printWarning(message)
	case "error", "fail":
		printError(message)
	default:
		printInfo(message)
	}
}

// printHelp prints help text
func printHelp(topic, content string) {
	if noColor {
		fmt.Printf("\n%s\n", topic)
		fmt.Println(strings.Repeat("-", len(topic)))
		fmt.Println(content)
	} else {
		fmt.Printf("\n%s\n", colorize(ColorCyan+ColorBold, topic))
		fmt.Println(colorize(ColorCyan, strings.Repeat("-", len(topic))))
		fmt.Println(content)
	}
}

// printVersion prints version information
func printVersion(version, commit, date string) {
	if noColor {
		fmt.Printf("Version: %s\n", version)
		if commit != "" {
			fmt.Printf("Commit: %s\n", commit)
		}
		if date != "" {
			fmt.Printf("Date: %s\n", date)
		}
	} else {
		fmt.Printf("Version: %s\n", colorize(ColorGreen, version))
		if commit != "" {
			fmt.Printf("Commit: %s\n", colorize(ColorCyan, commit))
		}
		if date != "" {
			fmt.Printf("Date: %s\n", colorize(ColorCyan, date))
		}
	}
}

// printConfig prints configuration information
func printConfig(key, value string) {
	if noColor {
		fmt.Printf("%s: %s\n", key, value)
	} else {
		fmt.Printf("%s: %s\n", colorize(ColorBlue, key), colorize(ColorCyan, value))
	}
}

// printResource prints resource information
func printResource(kind, name, namespace string) {
	if noColor {
		if namespace != "" {
			fmt.Printf("%s/%s in %s\n", kind, name, namespace)
		} else {
			fmt.Printf("%s/%s\n", kind, name)
		}
	} else {
		if namespace != "" {
			fmt.Printf("%s %s in %s\n", 
				colorize(ColorGreen, kind), 
				colorize(ColorCyan, name), 
				colorize(ColorYellow, namespace))
		} else {
			fmt.Printf("%s %s\n", 
				colorize(ColorGreen, kind), 
				colorize(ColorCyan, name))
		}
	}
}

// printDiff prints diff information
func printDiff(added, removed, modified int) {
	if noColor {
		fmt.Printf("Changes: +%d -%d ~%d\n", added, removed, modified)
	} else {
		fmt.Printf("Changes: %s %s %s\n", 
			colorize(ColorGreen, fmt.Sprintf("+%d", added)), 
			colorize(ColorRed, fmt.Sprintf("-%d", removed)), 
			colorize(ColorYellow, fmt.Sprintf("~%d", modified)))
	}
}

// printSummary prints a summary with counts
func printSummary(title string, counts map[string]int) {
	if noColor {
		fmt.Printf("\n%s\n", title)
		fmt.Println(strings.Repeat("-", len(title)))
		for key, value := range counts {
			fmt.Printf("%s: %d\n", key, value)
		}
	} else {
		fmt.Printf("\n%s\n", colorize(ColorCyan+ColorBold, title))
		fmt.Println(colorize(ColorCyan, strings.Repeat("-", len(title))))
		for key, value := range counts {
			fmt.Printf("%s: %s\n", colorize(ColorBlue, key), colorize(ColorGreen, fmt.Sprintf("%d", value)))
		}
	}
}

// printFooter prints a footer message
func printFooter(message string) {
	if noColor {
		fmt.Printf("\n%s\n", message)
	} else {
		fmt.Printf("\n%s\n", colorize(ColorCyan, message))
	}
}

// checkTerminalSize checks if the terminal supports colors
func checkTerminalSize() bool {
	// Simple check for non-interactive terminals
	if os.Getenv("TERM") == "" {
		return false
	}
	return true
}

// initStyle initializes the styling system
func initStyle() {
	// Check if colors should be disabled
	if !checkTerminalSize() {
		noColor = true
	}
}