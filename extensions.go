package main

import "strings"

var supportedExtensions = []string{"cmd", "bat"}

func verify(path string) bool {

	extension := strings.Split(path, ".")[1]
	extension = strings.ToLower(extension)

	var i int = 0
	for i < len(supportedExtensions) {
		if extension == supportedExtensions[i] {
			return true
		}
		i++
	}

	return false
}
