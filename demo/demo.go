package main

import (
	"os"

	"github.com/gehtsoft-usa/goharu/demo/demo"
)

func main() {
	var i int
	for i = 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "fontdemo":
			demo.Fontdemo()
		case "arcdemo":
			demo.Arcdemo()
		case "encryptiondemo":
			demo.Encryptiondemo()
		case "help":
			demo.Help()
		}

	}
}
