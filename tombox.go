package main

import (
	"tombox/activity"
	"tombox/demotables"
	"tombox/iface"

	"github.com/taikedz/goargs/goargs"
)

func main() {
	var default_run_table activity.ActivityTable

	parser := goargs.NewParser("TomatoBox : pomodoro timeboxer")

	tabletype := parser.Choices("tabletype", []string{"default", "mini", "micro"}, "Type of demo table to use")
	loopmode := parser.Bool("loop", false, "Whether to loop")

	parser.ParseCliArgs()

	switch *tabletype {
	case "default":
		default_run_table = demotables.Default(*loopmode)
	case "mini":
		default_run_table = demotables.Mini(*loopmode)
	case "micro":
		default_run_table = demotables.Micro(*loopmode)
	}

	iface.Run(default_run_table)
}
