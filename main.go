package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/fanchann/coi/internal/generator"
	"github.com/fanchann/coi/internal/parser"
	"github.com/fanchann/coi/utils"
)

var (
	fileName = flag.String("f", "", "file not found!")
)

func main() {
	flag.Parse()

	fun, err := parser.ExtractInterfaceFunctions(*fileName)
	if err != nil {
		panic(err)
	}

	fileOut := fmt.Sprintf("%s/%s_impl.go", utils.GetFileLocation(*fileName), utils.MakePrettyFileName(fun.IName))
	// string builder
	code := generator.GenerateImplementationContract(fun)

	ioutil.WriteFile(fileOut, []byte(code), fs.ModePerm)
}
