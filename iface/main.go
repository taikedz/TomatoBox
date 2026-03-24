package iface

/*
Main interface file

We want a few ones:

	* Main iface for start/stop, test with a thread for a timer that resets and beeps at every reset
	* iface to view queue
	* iface iface to add to queue and remove from queue
*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Run(strings []string) {
	myApp := app.New()
	window := myApp.NewWindow("TomatoBox")

	var ctext []fyne.CanvasObject

	for _, text := range strings {
		ctext = append(ctext, canvas.NewText(text, color.Opaque))
	}

	// After briefly studying https://docs.fyne.io/container/box/
	content := container.New(layout.NewVBoxLayout(), ctext...)

	window.SetContent(content)
	window.ShowAndRun()

}
