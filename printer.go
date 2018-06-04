package printer

import (
	"fmt"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
)

var instantiated *spinner.Spinner
var once sync.Once

// Create a singleton to the Spinner
// This ensures we only output to one line
func getPrinter() *spinner.Spinner {
	once.Do(func() {
		instantiated = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

		// instantiated.Writer = os.Stderr
	})

	return instantiated
}

// Print progress message
func Print(message string) {
	spinner := getPrinter()
	spinner.Suffix = fmt.Sprintf("  %s", message)
	spinner.Color("yellow")
	spinner.Start()
}

// Step - Print a final step message to the console
func Step(message string) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Yellow.Color(chalk.Dim.TextStyle("➜")), chalk.Bold.TextStyle(message))
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// SubStep - Print a final step message to the console
func SubStep(message string, indent int, last bool) {
	var indentString string

	for i := 1; i <= indent; i++ {
		indentString = fmt.Sprintf("   %s", indentString)
	}

	icon := ""

	switch indent {
	case 1:
		icon = "└─"
	default:
		icon = "├─"
	}

	if last {
		icon = "└─"
	}

	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s%s %s \n", indentString, chalk.Dim.TextStyle(icon), chalk.Dim.TextStyle(message))
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Finish - Print a final message to the console
func Finish(message string) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Green.Color("✔"), message)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Error - Finish with an error
func Error(message string) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Red.Color("✖"), message)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Fatal - Panic with an error
func Fatal(message string, err error) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Red.Color("✖"), message)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
	panic(err)
}
