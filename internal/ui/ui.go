package ui

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	// RepoAnt palette: danger-forward reds with amber highlights.
	Primary   = color.New(color.FgHiRed, color.Bold)
	Secondary = color.New(color.FgHiWhite)
	Accent    = color.New(color.FgHiYellow, color.Bold)
	Highlight = color.New(color.FgHiWhite, color.Bold)

	// Status colors
	Success = color.New(color.FgGreen, color.Bold)
	Warning = color.New(color.FgYellow, color.Bold)
	Error   = color.New(color.FgRed, color.Bold)
	Info    = color.New(color.FgCyan)

	// Text colors
	Muted   = color.New(color.FgHiBlack)
	Bold    = color.New(color.Bold)
	Dim     = color.New(color.Faint)
	Cyan    = color.New(color.FgHiCyan)
	Magenta = color.New(color.FgMagenta)
	White   = color.New(color.FgHiWhite)
)

const logo = `
██████╗ ███████╗██████╗  ██████╗  █████╗ ███╗   ██╗████████╗
██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔══██╗████╗  ██║╚══██╔══╝
██████╔╝█████╗  ██████╔╝██║   ██║███████║██╔██╗ ██║   ██║
██╔══██╗██╔══╝  ██╔═══╝ ██║   ██║██╔══██║██║╚██╗██║   ██║
██║  ██║███████╗██║     ╚██████╔╝██║  ██║██║ ╚████║   ██║
╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝
`

const smallLogo = `
╭─────────────────────────────────────────────────╮
│                                                 │
│  ██████╗ ███████╗██████╗  ██████╗  █████╗ ███╗   ██╗████████╗  │
│  ██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔══██╗████╗  ██║╚══██╔══╝  │
│  ██████╔╝█████╗  ██████╔╝██║   ██║███████║██╔██╗ ██║   ██║     │
│  ██╔══██╗██╔══╝  ██╔═══╝ ██║   ██║██╔══██║██║╚██╗██║   ██║     │
│  ██║  ██║███████╗██║     ╚██████╔╝██║  ██║██║ ╚████║   ██║     │
│  ╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝     │
│                                                 │
│  delete and mass delete github repo hassle free │
│                                                 │
╰─────────────────────────────────────────────────╯
`

func PrintBanner() {
	// Warm gradient with ant theme
	gradient := []color.Attribute{
		color.FgHiRed,
		color.FgRed,
		color.FgHiYellow,
		color.FgRed,
		color.FgHiRed,
		color.FgHiYellow,
	}

	lines := []string{
		"",
		"       \\   /",
		"        \\./        ██████╗ ███████╗██████╗  ██████╗  █████╗ ███╗   ██╗████████╗",
		"       (o o)       ██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔══██╗████╗  ██║╚══██╔══╝",
		"      ╭─┴─┴─╮      ██████╔╝█████╗  ██████╔╝██║   ██║███████║██╔██╗ ██║   ██║",
		"      │ ANT │      ██╔══██╗██╔══╝  ██╔═══╝ ██║   ██║██╔══██║██║╚██╗██║   ██║",
		"      ╰─────╯      ██║  ██║███████╗██║     ╚██████╔╝██║  ██║██║ ╚████║   ██║",
		"       /   \\       ╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝",
		"      /     \\",
		"",
		"  🐜 delete and mass delete github repo hassle free",
		"  by @aasishdairel",
		"",
	}

	fmt.Println()
	for i, line := range lines {
		if i >= 1 && i <= 6 { // Logo lines with ant
			c := color.New(gradient[i-1%len(gradient)], color.Bold)
			c.Println(line)
		} else { // Ant and tagline
			if strings.Contains(line, "🐜") {
				Accent.Println(line)
			} else {
				White.Println(line)
			}
		}
	}
	fmt.Println()
}

func PrintSuccess(format string, a ...interface{}) {
	Success.Print("  ✔ ")
	White.Printf(format+"\n", a...)
}

func PrintError(format string, a ...interface{}) {
	Error.Print("  ✖ ")
	White.Printf(format+"\n", a...)
}

func PrintWarning(format string, a ...interface{}) {
	Warning.Print("  ⚠ ")
	White.Printf(format+"\n", a...)
}

func PrintInfo(format string, a ...interface{}) {
	Info.Print("  ℹ ")
	White.Printf(format+"\n", a...)
}

func PrintStep(step int, format string, a ...interface{}) {
	Secondary.Printf("  [%d] ", step)
	White.Printf(format+"\n", a...)
}

func PrintRepo(name string, isPrivate bool, isProtected bool) {
	if isProtected {
		Muted.Print("  🔒 ")
		Muted.Printf("%s ", name)
		Dim.Println("(protected)")
	} else if isPrivate {
		Accent.Print("  🔐 ")
		White.Println(name)
	} else {
		Cyan.Print("  📂 ")
		White.Println(name)
	}
}

func PrintRepoSimple(name string) {
	Cyan.Print("  📦 ")
	White.Println(name)
}


func PrintDeleting(name string) {
	fmt.Println()
	Warning.Print("  🗑️  Deleting: ")
	Bold.Println(name)
}

func PrintDeleted(name string) {
	fmt.Println()
	Success.Print("  ✔ Successfully deleted: ")
	Bold.Println(name)
}

func PrintHeader(title string) {
	fmt.Println()
	Primary.Printf("  ◆ %s\n", title)
	Muted.Println("  " + repeatString("─", len(title)+4))
}

func PrintCount(count int, singular, plural string) {
	word := plural
	if count == 1 {
		word = singular
	}
	Muted.Printf("  Found ")
	Primary.Printf("%d ", count)
	Muted.Printf("%s\n\n", word)
}

func PrintHint(hint string) {
	fmt.Println()
	Muted.Printf("  💡 %s\n", hint)
}

func PrintDivider() {
	Muted.Println("\n  " + repeatString("─", 40))
}

func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// Spinner creates a simple loading indicator
func PrintLoading(message string) {
	Info.Print("  ⏳ ")
	Muted.Println(message)
}
