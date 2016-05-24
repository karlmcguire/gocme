package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/karlmcguire/gocme"
)

func main() {
	// Get the id of the current window.
	id, err := gocme.GetId()
	if err != nil {
		panic(err)
	}

	// Get the absolute filepath of the current window.
	filename, err := gocme.GetFilename(id)
	if err != nil {
		panic(err)
	}

	// Get the contents of the current window.
	body, err := gocme.GetBody(id)
	if err != nil {
		panic(err)
	}

	// Run "gofmt" on the file.
	rawfix, err := exec.Command("gofmt", filename).Output()
	if err != nil {
		// Gofmt returned an error.
		if e, ok := err.(*exec.ExitError); ok {
			// Output the gofmt error to Errors window.
			fmt.Print(string(e.Stderr))
		}
	} else {
		fix := string(rawfix)
		// Check if new contents is any different than present contents.
		if strings.Compare(body, fix) != 0 {
			// Set the body of the current window to the new contents.
			e := gocme.SetBody(id, fix)
			if e != nil {
				panic(e)
			}
		} else {
			// Already formatted, all done.
		}
	}
}
