package main

import (
	"os"

	"github.com/gehtsoft-usa/goharu/demo/demo"
)

func main() {
	var i int
	for i = 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "font":
			demo.Fontdemo()
			break
		case "arc":
			demo.Arcdemo()
			break
		case "encrypt":
			demo.Encryptiondemo()
			break
		case "encodings":
			demo.Encodings()
			break
		case "cp1251":
			demo.Cp1251Demo()
			break
		case "jp":
			demo.JpDemo()
			break
		case "png":
			demo.Pngdemo()
			break
		case "jpg":
			demo.Jpgdemo()
			break
		case "text":
			demo.Textdemo()
			break
		case "help":
			demo.Help()
			break
		}

	}
}
