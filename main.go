package main

import (
	"boilerplate-go/cmd"
	"os"
)

func main() {
	//	initiate config

	//	initiate serve
	args := os.Args[1:]
	cmd.Command(args)
}
