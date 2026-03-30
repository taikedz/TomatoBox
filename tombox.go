package main

import (
	act "tombox/activity"
	"tombox/iface"
)

const SEC = 1
const MIN = 60 * SEC
const HR = 60 * MIN

func main() {
	default_run_table := miniRuntable() //defaultRunTable()

	iface.Run(default_run_table)
}

func miniRuntable() act.ActivityTable {
	atable := act.ActivityTable{Activities: []act.Activity{}, Runmode: act.Loop}

	atable.Append(act.Activity{Title: "Focus", Duration: 3 * SEC, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 2 * SEC, Wrapup: 10 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Last", Duration: 1 * SEC, Wrapup: 10 * SEC, Auto: true})

	return atable
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
