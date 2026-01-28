package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	// Brand colors
	Primary   = color.New(color.FgHiCyan, color.Bold)
	Secondary = color.New(color.FgHiMagenta)
	Accent    = color.New(color.FgHiYellow)

	// Status colors
	Success = color.New(color.FgHiGreen, color.Bold)
	Warning = color.New(color.FgHiYellow, color.Bold)
	Error   = color.New(color.FgHiRed, color.Bold)
	Info    = color.New(color.FgHiBlue)

	// Text colors
	Muted   = color.New(color.FgHiBlack)
	Bold    = color.New(color.Bold)
	Dim     = color.New(color.Faint)
	Cyan    = color.New(color.FgCyan)
	Magenta = color.New(color.FgMagenta)
	White   = color.New(color.FgWhite)
)

const logo = `
   ▄████  ██▓▄▄▄█████▓  ██████  ▄▄▄        █████▒▓█████     ██▀███   ███▄ ▄███▓
  ██▒ ▀█▒▓██▒▓  ██▒ ▓▒▒██    ▒ ▒████▄    ▓██   ▒ ▓█   ▀    ▓██ ▒ ██▒▓██▒▀█▀ ██▒
 ▒██░▄▄▄░▒██▒▒ ▓██░ ▒░░ ▓██▄   ▒██  ▀█▄  ▒████ ░ ▒███      ▓██ ░▄█ ▒▓██    ▓██░
 ░▓█  ██▓░██░░ ▓██▓ ░   ▒   ██▒░██▄▄▄▄██ ░▓█▒  ░ ▒▓█  ▄    ▒██▀▀█▄  ▒██    ▒██
 ░▒▓███▀▒░██░  ▒██▒ ░ ▒██████▒▒ ▓█   ▓██▒░▒█░    ░▒████▒   ░██▓ ▒██▒▒██▒   ░██▒
  ░▒   ▒ ░▓    ▒ ░░   ▒ ▒▓▒ ▒ ░ ▒▒   ▓▒█░ ▒ ░    ░░ ▒░ ░   ░ ▒▓ ░▒▓░░ ▒░   ░  ░
   ░   ░  ▒ ░    ░    ░ ░▒  ░ ░  ▒   ▒▒ ░ ░       ░ ░  ░     ░▒ ░ ▒░░  ░      ░
 ░ ░   ░  ▒ ░  ░      ░  ░  ░    ░   ▒    ░ ░       ░        ░░   ░ ░      ░
       ░  ░                 ░        ░  ░           ░  ░      ░            ░
`

const smallLogo = `
  ╭─────────────────────────────────╮
  │  🗑️  gitsafe-rm                 │
  │  Safely delete GitHub repos     │
  ╰─────────────────────────────────╯
`

func PrintBanner() {
	gradient := []color.Attribute{
		color.FgHiCyan,
		color.FgHiBlue,
		color.FgHiMagenta,
		color.FgHiCyan,
	}

	lines := []string{
		"  ╭─────────────────────────────────────────╮",
		"  │                                         │",
		"  │   🗑️  gitsafe-rm                        │",
		"  │   Safely delete GitHub repositories     │",
		"  │                                         │",
		"  ╰─────────────────────────────────────────╯",
	}

	fmt.Println()
	for i, line := range lines {
		c := color.New(gradient[i%len(gradient)], color.Bold)
		c.Println(line)
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
