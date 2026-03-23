package activity

import "fmt"

type Activity struct {
	Title            string
	Duration, Wrapup PrettyTime
	Auto             bool
}

type RunMode int

const (
	Loop RunMode = iota
	Stop
)

type ActivityTable struct {
	Activities []Activity
	Runmode    RunMode
}

func (at *ActivityTable) Append(act Activity) {
	at.Activities = append(at.Activities, act)
}

func (at *ActivityTable) Insert(idx int, act Activity) {
	front, back := at.Activities[:idx], at.Activities[idx:]
	at.Activities = []Activity{}
	at.Activities = append(at.Activities, front...)
	at.Activities = append(at.Activities, act)
	at.Activities = append(at.Activities, back...)
}

type PrettyTime int

func (pt PrettyTime) String() string {
	secs := pt % 60
	total_mins := (pt - secs) / 60
	return fmt.Sprintf("%02d:%02d", total_mins, secs)

	// mins := total_mins % 60
	// hours := (total_mins - mins) / 60
	// return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
}
