package main

import (
	"flag"

	"github.com/fanchann/coi/cmd"
)

var (
	fileName = flag.String("f", "", "input go file")
)

func main() {
	flag.Parse()
	cmd.AppRun(*fileName)
}
