package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var vars []string
var vals []string

var workingDirectory string = ""
var funcFound = false

func parse(line string, lines []string, isRecursive bool) {

	if len(line) < 1 {
		return
	}

	line = strings.TrimSpace(line)
	terms := strings.Split(line, " ")
	command := terms[0]
	command = strings.ToLower(command)

	if startsWith(line, ":") {
		funcFound = true
	}

	if funcFound {
		if !isRecursive {
			return
		}
	}

	//prints text:
	if command == "echo" {
		args := strings.ReplaceAll(line, "echo", "")
		echo(args)
		return
	}

	//changes current working directory:
	if command == "cd" {
		workingDirectory = terms[1]
		return
	}

	//handles input:
	if command == "set" {
		if terms[1] == "/p" {
			vars = append(vars, terms[2])
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(strings.Split(line, "=")[1])
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			vals = append(vals, value)
		}
		return
	}

	//handles function calling:
	if command == "goto" {
		funcname := terms[1]
		funcname = ":" + funcname
		var i int = 0
		for i < len(lines) {
			currentline := lines[i]
			currentline = strings.TrimSpace(currentline)
			currentterms := strings.Split(currentline, " ")
			if currentterms[0] == funcname {
				j := i + 1
				for j < len(lines) {
					parse(lines[j], lines, true)
					j++
				}
			}
			i++
		}
	}
}

func echo(args string) {
	args = strings.TrimSpace(args)
	if strings.Contains(args, "%") {
		fmt.Println(findVarValue(args))
	} else {
		fmt.Println(args)
	}
}

func findVarValue(args string) string {

	arr := strings.Split(args, "")
	var i int = 0
	var start bool = false
	var name string = ""

	for i < len(arr) {

		if start {
			if arr[i] != "%" {
				name = name + arr[i]
			}
		}
		if arr[i] == "%" {
			if start == true {
				break
			}
			start = true
		}

		i++
	}

	var j int = 0
	for j < len(vals) {
		if vars[j] == name {
			return vals[j]
		}
		j++
	}

	return ""
}

func startsWith(line string, substring string) bool {

	line = strings.TrimSpace(line)
	arr := strings.Split(line, "")
	if arr[0] == ":" {
		return true
	}

	return false
}
