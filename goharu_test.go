package goharu_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gehtsoft-usa/goharu"
)

func TestVersion(t *testing.T) {
	fmt.Printf("%s\n", goharu.Version())
}

func TestDebug(t *testing.T) {
	a := goharu.Create(true)
	defer a.Free()
	a.NewDoc()
	defer a.FreeDoc()

	a.Save("test/test.pdf")
	c := a.Content()
	f, _ := os.Create("test/test1.pdf")
	defer f.Close()
	f.Write(c)

	var i goharu.Image

	i = a.LoadPngImageFromFile("test/logo.png")
	fmt.Printf("%d,%d\n", i.Width(), i.Height())

	b, _ := ioutil.ReadFile("test/logo.png")
	i = a.LoadPngImageFromMem(b)
	fmt.Printf("%d,%d\n", i.Width(), i.Height())

}
