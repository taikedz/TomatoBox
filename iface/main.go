package iface

/*
Main interface file

We want a few ones:

	* Main iface for start/stop, test with a thread for a timer that resets and beeps at every reset
	* iface to view queue
	* iface iface to add to queue and remove from queue
*/

import "fmt"

func Run() {
	fmt.Println("Fyne goes here")
}
