package demotables

import (
	act "tombox/activity"
)

const SEC = 1
const MIN = 60 * SEC
const HR = 60 * MIN

func doLoop(loop bool) act.RunMode {
	if loop {
		return act.Loop
	} else {
		return act.Stop
	}
}

func Mini(loop bool) act.ActivityTable {
	atable := act.ActivityTable{Activities: []act.Activity{}, Runmode: doLoop(loop)}

	atable.Append(act.Activity{Title: "Focus", Duration: 3 * SEC, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 2 * SEC, Wrapup: 10 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Last", Duration: 1 * SEC, Wrapup: 10 * SEC, Auto: true})

	return atable
}

func Micro(loop bool) act.ActivityTable {
	atable := act.ActivityTable{Activities: []act.Activity{}, Runmode: doLoop(loop)}

	atable.Append(act.Activity{Title: "Focus", Duration: 1 * SEC, Wrapup: 30 * SEC, Auto: true})
	atable.Append(act.Activity{Title: "Break", Duration: 1 * SEC, Wrapup: 10 * SEC, Auto: true})

	return atable
}

func Default(loop bool) act.ActivityTable {
	atable := act.ActivityTable{Activities: []act.Activity{}, Runmode: doLoop(loop)}

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
