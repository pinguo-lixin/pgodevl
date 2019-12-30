package main

import (
	"fmt"
	"os"
)

// all commands supported
const (
	cmdListURL = "listurl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command required, supported commands such as: listurl")
		os.Exit(1)
	}
	cmd := os.Args[1]
	if !validCmd(cmd) {
		fmt.Println("command not supported")
		os.Exit(1)
	}

	pgoConf := defaultPgoConfig()

	if cmd == cmdListURL {
		listURL(pgoConf.ControllerPath)
	}
}

func validCmd(cmd string) bool {
	return cmd == cmdListURL
}
