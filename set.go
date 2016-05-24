package gocme

import (
	"os/exec"
	"strings"
)


// SetBody replaces the contents of a given window.	
// TODO: make this cleaner and document what's going on.
func SetBody(id string, body string) error {
	winaddr := exec.Command("9p", "write", "acme/"+id+"/addr")
	winaddr.Stdin = strings.NewReader(",")
	err := winaddr.Run()
	if err != nil {
		return err
	}

	winctl := exec.Command("9p", "write", "acme/"+id+"/ctl")
	winctl.Stdin = strings.NewReader("dot=addr")
	err = winctl.Run()
	if err != nil {
		return err
	}

	winwrite := exec.Command("9p", "write", "acme/"+id+"/wrsel")
	winwrite.Stdin = strings.NewReader(body)
	err = winwrite.Run()
	if err != nil {
		return err
	}

	winaddr2 := exec.Command("9p", "write", "acme/"+id+"/addr")
	winaddr2.Stdin = strings.NewReader("0")
	err = winaddr2.Run()
	if err != nil {
		return err
	}

	winctl2 := exec.Command("9p", "write", "acme/"+id+"/ctl")
	winctl2.Stdin = strings.NewReader("dot=addr")
	err = winctl2.Run()
	if err != nil {
		return err
	}

	// Put the new contents to disk.
	err = Put(id)
	if err != nil {
		return err
	}

	// Mark the window as clean.
	err = Clean(id)
	if err != nil {
		return err 
	}

	return nil
}