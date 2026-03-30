package iface

/*
Main interface file

We want a few ones:

	* Main iface for start/stop, test with a thread for a timer that resets and beeps at every reset
	* iface to view queue
	* iface iface to add to queue and remove from queue
*/

import (
	"fmt"
	"image/color"
	"time"

	act "tombox/activity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// WIP - following https://www.youtube.com/watch?v=vSDtHyBvnXE
func updateContent(display *widget.Label) {
	display.SetText("new value")
}

func runActivities(display *widget.Label, table act.ActivityTable) {
	for _, activity := range table.Activities {
		countdown := activity.Duration
		for countdown >= 0 {
			display.SetText(fmt.Sprintf("%s : %s", activity.Title, countdown))

			time.Sleep(1 * time.Second)
			countdown -= 1
		}
	}

	println("DONE----")
}

func Run(table act.ActivityTable) {
	myApp := app.New()
	window := myApp.NewWindow("TomatoBox")

	display := widget.NewLabel("")
	window.SetContent(display)

	go runActivities(display, table)

	// content := gridInGrid(table)

	window.ShowAndRun()

}

func gridInGrid(table act.ActivityTable) *fyne.Container {
	// Check https://docs.fyne.io/explore/layouts/
	cont := container.New(layout.NewGridLayout(4), fText("Title", "Duration", "Wrap-up", "Auto run")...)
	for _, item := range table.Activities {
		cont.Add(fText(item.Title)[0])
		cont.Add(fText(item.Duration.String())[0])
		cont.Add(fText(item.Wrapup.String())[0])
		cont.Add(fText(fmt.Sprintf("%v", item.Auto))[0])
	}
	return cont
}

/*
Convert a sequence of strings into fyne Text structs.
*/
func fText(text ...string) []fyne.CanvasObject {
	var things []fyne.CanvasObject
	for _, item := range text {
		things = append(things, canvas.NewText(item, color.Opaque))
	}
	return things
}
