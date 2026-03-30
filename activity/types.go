package activity

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
