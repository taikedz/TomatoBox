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
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Run(titles []string) {
	myApp := app.New()
	window := myApp.NewWindow("TomatoBox")

	var ctext []fyne.CanvasObject

	for _, text := range titles {
		ctext = append(ctext, canvas.NewText(text, color.Opaque))
		if strings.Contains(strings.ToLower(text), "break") {
			space := layout.NewSpacer()
			space.Resize(fyne.NewSize(100.0, 100.0)) // no effect
			ctext = append(ctext, space)
		}
	}

	// After briefly studying https://docs.fyne.io/container/box/
	content := container.New(layout.NewGridLayout(3), ctext...)

	window.SetContent(content)
	window.ShowAndRun()

}
