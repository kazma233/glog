package main

import (
	"os"

	"github.com/lu4p/binclude/bincludegen"
)

// at root folder: go generate .\config\
func main() {
	os.Exit(bincludegen.Main1())
}
