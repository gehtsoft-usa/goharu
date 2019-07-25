package goharu_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gehtsoft-usa/goharu"
)

func TestVersion(t *testing.T) {
	fmt.Printf("%s\n", goharu.Version())
}

func TestDebug(t *testing.T) {
	a := goharu.Create()
	defer a.Free()
	a.NewDoc()
	defer a.FreeDoc()

	a.Save("test/test.pdf")
	c := a.Content()
	f, _ := os.Create("test/test1.pdf")
	defer f.Close()
	f.Write(c)

}
