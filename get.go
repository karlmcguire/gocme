package gocme

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

// GetId returns the id of the current acme window.
func GetId() (string, error) {
	if len(os.Getenv("winid")) == 0 {
		return "", errors.New("no winid environment variable")
	}
	return os.Getenv("winid"), nil
}

// GetTags returns a string slice of all tags in the window given.
func GetTags(id string) ([]string, error) {
	// Get a list of all the tags in the window.
	rawtags, err := exec.Command("9p", "read", "acme/"+id+"/tag").Output()
	if err != nil {
		return nil, err
	}

	// Seperate the byte slice at spaces, leaving us with [][]byte tag names.
	btags := bytes.Split(rawtags, []byte(" "))

	// Convert the byte tags into a string slice we can return.
	tags := make([]string, 0)
	for _, v := range btags {
		tags = append(tags, string(v))
	}

	return tags, nil
}

// GetBody returns the contents of the window given.
func GetBody(id string) (string, error) {
	// Get the byte contents of the window given by id.
	rawbody, err := exec.Command("9p", "read", "acme/"+id+"/body").Output()
	if err != nil {
		return "", err
	}

	return string(rawbody), nil
}

// GetFilename returns the filename of the file being editted in the current window given.
func GetFilename(id string) (string, error) {
	tags, err := GetTags(id)
	if err != nil {
		return "", err
	}

	return tags[0], nil
}
