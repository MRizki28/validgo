package main

import (
	"fmt"
	"os"

	"github.com/MRizki28/validgo/generator"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("usage: validgo make:request RegisterUser")
		return
	}

	command := args[1]
	name := args[2]

	switch command {

	case "make:request":
		generator.MakeRequest(name)

	default:
		fmt.Println("command not found")
	}
}