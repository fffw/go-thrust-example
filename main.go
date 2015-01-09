package main

import (
	"fmt"
	"os"

	"github.com/miketheprogrammer/go-thrust"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/menu"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/commands"
)

func main() {
	var url = os.Args[1]
	fmt.Println(url)
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	// thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()
	thrustWindow := window.NewWindow(url, nil)
	// "http://127.0.0.1:14547/65l7mSNgZKeLDpJdnW6U-Q/index.html"
	thrustWindow.Show()
	thrustWindow.Focus()

	bakeApplicationMenu()
	thrust.LockThread()
}

func bakeApplicationMenu() {
	applicationMenu := thrust.NewMenu()
	applicationMenuRoot := thrust.NewMenu()
	fileMenu := thrust.NewMenu()
	applicationMenuRoot.AddItem(1, "About")
	applicationMenuRoot.RegisterEventHandlerByCommandID(1,
		func(reply commands.CommandResponse, item *menu.MenuItem) {
			fmt.Println("About Handled")
		})
	fileMenu.AddItem(2, "Open")
	fileMenu.RegisterEventHandlerByCommandID(2,
		func(reply commands.CommandResponse, item *menu.MenuItem) {
			fmt.Println("Open Handled")
		})
	fileMenu.AddItem(3, "Edit")
	fileMenu.AddSeparator()
	fileMenu.AddItem(4, "Close")
	fileMenu.RegisterEventHandlerByCommandID(4,
		func(reply commands.CommandResponse, item *menu.MenuItem) {
			fmt.Println("Close Event Handled")
			thrust.Exit()
		})
	applicationMenu.AddSubmenu(5, "Application", applicationMenuRoot)
	applicationMenu.AddSubmenu(6, "File", fileMenu)
	applicationMenu.SetApplicationMenu()
}
