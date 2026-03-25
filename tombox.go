package main

import (
	"fmt"
	act "tombox/activity"
	"tombox/iface"
)

const SEC = 1
const MIN = 60 * SEC
const HR = 60 * MIN

func main() {
	default_run_table := defaultRunTable()

	events := []string{}
	for _, item := range default_run_table.Activities {
		events = append(events, fmt.Sprintf("%s for %v (%v wrap up)", item.Title, item.Duration, item.Wrapup))
	}

	iface.Run(events)
}

func defaultRunTable() act.ActivityTable {
	atable := act.ActivityTable{Activities: []act.Activity{}, Runmode: act.Loop}

	atable.Append(act.Activity{Title: "Focus", Duration: 25 * MIN, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 5 * MIN, Wrapup: 10 * SEC, Auto: true})

	atable.Append(act.Activity{Title: "Focus", Duration: 25 * MIN, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 5 * MIN, Wrapup: 10 * SEC, Auto: true})

	atable.Append(act.Activity{Title: "Focus", Duration: 25 * MIN, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 5 * MIN, Wrapup: 10 * SEC, Auto: true})

	atable.Append(act.Activity{Title: "Focus", Duration: 25 * MIN, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Long break", Duration: 1 * HR, Wrapup: 5 * MIN, Auto: true})

	return atable
}
