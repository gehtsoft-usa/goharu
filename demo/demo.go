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
			break
		case "arcdemo":
			demo.Arcdemo()
			break
		case "encryptiondemo":
			demo.Encryptiondemo()
			break
		case "encodings":
			demo.Encodings()
			break
		case "cp1251":
			demo.Cp1251Demo()
			break
		case "help":
			demo.Help()
			break
		}

	}
}
