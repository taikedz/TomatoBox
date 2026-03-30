package activity

import "fmt"

type PrettyTime int

func (pt PrettyTime) String() string {
	secs := pt % 60
	total_mins := (pt - secs) / 60

	if total_mins < 60 {
		return fmt.Sprintf("%02d:%02d", total_mins, secs)
	}

	mins := total_mins % 60
	hours := (total_mins - mins) / 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
}
