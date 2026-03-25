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

	act "tombox/activity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Run(table act.ActivityTable) {
	myApp := app.New()
	window := myApp.NewWindow("TomatoBox")

	content := gridInGrid(table)

	window.SetContent(content)
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

func fText(text ...string) []fyne.CanvasObject {
	var things []fyne.CanvasObject
	for _, item := range text {
		things = append(things, canvas.NewText(item, color.Opaque))
	}
	return things
}
