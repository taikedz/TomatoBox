package activity

import (
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {
	var duration PrettyTime = 310
	var s string

	s = fmt.Sprintf("%s", duration)
	if s != "05:10" {
		t.Fail()
	}

	duration = 304
	s = fmt.Sprintf("%s", duration)
	if s != "05:04" {
		t.Fail()
	}

	s = fmt.Sprintf("%s", duration-2)
	if s != "05:02" {
		t.Fail()
	}

	duration = 1*60*60 + 3*60 + 17
	s = fmt.Sprintf("%s", duration)
	if s != "01:03:17" {
		t.Fail()
	}
}
