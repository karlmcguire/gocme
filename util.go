package gocme

import (
	"os/exec"
	"strings"
)

// Marks the current window as dirty.
func Dirty(id string) error {
	windirty := exec.Command("9p", "write", "acme/"+id+"/ctl")
	windirty.Stdin = strings.NewReader("dirty")
	return windirty.Run()
}

// Marks the current window as clean.
func Clean(id string) error {
	winclean := exec.Command("9p", "write", "acme/"+id+"/ctl")
	winclean.Stdin = strings.NewReader("clean")
	return winclean.Run()
}

// Writes the contents of the current window to disk.
func Put(id string) error {
	winput := exec.Command("9p", "write", "acme/"+id+"/ctl")
	winput.Stdin = strings.NewReader("put")
	return winput.Run()
}