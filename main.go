package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: universalbatch (filepath)")
		return
	}

	filepath := args[0]

	if !verify(filepath) {
		fmt.Println("path specified must be a '.cmd' or '.bat' file")
		return
	}
	if verify(filepath) {

		lines := read(filepath)
		var i int = 0
		for i < len(lines) {
			parse(lines[i], lines, false)
			i++
		}
	}
}
