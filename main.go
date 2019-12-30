package main

import "os"

// all commands supported
const (
	cmdListURL = "listurl"
)

func main() {
	println("pgodevl")
	wd, _ := os.Getwd()
	println(wd)
}
