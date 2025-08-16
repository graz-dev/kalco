package cmd

import (
	"fmt"
	"os"
	"strings"
)

// Color codes for terminal output
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
	ColorDim    = "\033[2m"
)

// Icons for different message types
const (
	IconSuccess = "✅"
	IconError   = "❌"
	IconWarning = "⚠️"
	IconInfo    = "ℹ️"
	IconRocket  = "🚀"
	IconFolder  = "📁"
	IconGear    = "⚙️"
	IconChart   = "📊"
	IconLock    = "🔒"
	IconSearch  = "🔍"
	IconPackage = "📦"
	IconGit     = "🔀"
)

// isColorEnabled checks if colored output should be used
func isColorEnabled() bool {
	if noColor {
		return false
	}
	
	// Check if we're in a terminal
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	
	// Check NO_COLOR environment variable
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	
	return true
}

// colorize applies color to text if colors are enabled
func colorize(color, text string) string {
	if !isColorEnabled() {
		return text
	}
	return color + text + ColorReset
}

// printSuccess prints a success message with green color and checkmark
func printSuccess(message string) {
	icon := IconSuccess
	if !isColorEnabled() {
		icon = "[SUCCESS]"
	}
	fmt.Printf("%s %s\n", icon, colorize(ColorGreen, message))
}

// printError prints an error message with red color and X mark
func printError(message string) {
	icon := IconError
	if !isColorEnabled() {
		icon = "[ERROR]"
	}
	fmt.Printf("%s %s\n", icon, colorize(ColorRed, message))
}

// printWarning prints a warning message with yellow color and warning icon
func printWarning(message string) {
	icon := IconWarning
	if !isColorEnabled() {
		icon = "[WARNING]"
	}
	fmt.Printf("%s %s\n", icon, colorize(ColorYellow, message))
}

// printInfo prints an info message with blue color and info icon
func printInfo(message string) {
	icon := IconInfo
	if !isColorEnabled() {
		icon = "[INFO]"
	}
	fmt.Printf("%s %s\n", icon, colorize(ColorBlue, message))
}

// printHeader prints a styled header
func printHeader(title string) {
	if !isColorEnabled() {
		fmt.Printf("\n=== %s ===\n", strings.ToUpper(title))
		return
	}
	
	border := strings.Repeat("━", len(title)+4)
	fmt.Printf("\n%s\n", colorize(ColorCyan, border))
	fmt.Printf("%s %s %s\n", colorize(ColorCyan, "━"), colorize(ColorBold+ColorWhite, title), colorize(ColorCyan, "━"))
	fmt.Printf("%s\n", colorize(ColorCyan, border))
}

// printSubHeader prints a styled sub-header
func printSubHeader(title string) {
	if !isColorEnabled() {
		fmt.Printf("\n--- %s ---\n", title)
		return
	}
	
	fmt.Printf("\n%s %s\n", colorize(ColorPurple, "▶"), colorize(ColorBold, title))
}

// printBanner prints the kalco banner
func printBanner() {
	if noColor {
		fmt.Println("KALCO - Kubernetes Analysis & Lifecycle Control")
		fmt.Println("Extract, validate, analyze, and version control your cluster")
		return
	}
	
	banner := `
██╗  ██╗ █████╗ ██╗      ██████╗ ██████╗ 
██║ ██╔╝██╔══██╗██║     ██╔════╝██╔═══██╗
█████╔╝ ███████║██║     ██║     ██║   ██║
██╔═██╗ ██╔══██║██║     ██║     ██║   ██║
██║  ██╗██║  ██║███████╗╚██████╗╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ 
                                          
🚀 Kubernetes Analysis & Lifecycle Control
   Extract, validate, analyze, and version control your cluster`

	fmt.Println(colorize(ColorCyan+ColorBold, banner))
	fmt.Println()
}

// printCommandHeader prints a styled command header
func printCommandHeader(command, description string) {
	if !isColorEnabled() {
		fmt.Printf("=== %s ===\n%s\n\n", strings.ToUpper(command), description)
		return
	}
	
	fmt.Printf("%s %s %s\n", 
		colorize(ColorBlue+ColorBold, "▶"), 
		colorize(ColorWhite+ColorBold, command), 
		colorize(ColorDim, description))
	fmt.Println()
}

// printFlag prints a styled flag description
func printFlag(flag, description string) {
	if !isColorEnabled() {
		fmt.Printf("  %s: %s\n", flag, description)
		return
	}
	
	fmt.Printf("  %s %s\n", 
		colorize(ColorGreen+ColorBold, flag), 
		colorize(ColorWhite, description))
}

// printExample prints a styled example
func printExample(title, command string) {
	if !isColorEnabled() {
		fmt.Printf("Example - %s:\n  %s\n\n", title, command)
		return
	}
	
	fmt.Printf("%s %s\n", 
		colorize(ColorYellow+ColorBold, "Example:"), 
		colorize(ColorWhite+ColorBold, title))
	fmt.Printf("  %s\n\n", 
		colorize(ColorCyan, command))
}

// printTable prints data in a simple table format
func printTable(headers []string, rows [][]string) {
	if len(rows) == 0 {
		printInfo("No data to display")
		return
	}
	
	// Calculate column widths
	widths := make([]int, len(headers))
	for i, header := range headers {
		widths[i] = len(header)
	}
	
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}
	
	// Print header
	fmt.Print(colorize(ColorBold, ""))
	for i, header := range headers {
		fmt.Printf("%-*s", widths[i]+2, header)
	}
	fmt.Print(colorize(ColorReset, ""))
	fmt.Println()
	
	// Print separator
	for i := range headers {
		fmt.Print(strings.Repeat("-", widths[i]+2))
	}
	fmt.Println()
	
	// Print rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				fmt.Printf("%-*s", widths[i]+2, cell)
			}
		}
		fmt.Println()
	}
}

// printProgress prints a progress indicator
func printProgress(current, total int, message string) {
	if !isColorEnabled() {
		fmt.Printf("[%d/%d] %s\n", current, total, message)
		return
	}
	
	percentage := float64(current) / float64(total) * 100
	fmt.Printf("%s [%d/%d] %.1f%% %s\n", 
		colorize(ColorBlue, IconGear), 
		current, 
		total, 
		percentage, 
		message)
}

// printSeparator prints a visual separator
func printSeparator() {
	if !isColorEnabled() {
		fmt.Println(strings.Repeat("-", 50))
		return
	}
	
	fmt.Println(colorize(ColorDim, strings.Repeat("─", 50)))
}