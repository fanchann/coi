package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/fanchann/coi/internal/generator"
	"github.com/fanchann/coi/internal/parser"
	"github.com/fanchann/coi/utils"
)

func AppRun(file string) {

	ext := filepath.Ext(file)
	if ext != "go" {
		utils.ErrorWrapper(errors.New("file extension must go!"))
	}

	fun, err := parser.ExtractInterfaceFunctions(file)
	utils.ErrorWrapper(err)

	fileOut := fmt.Sprintf("%s/%s_impl.go", utils.GetFileLocation(file), utils.MakePrettyFileName(fun.IName))
	fmt.Printf("fileOut: %v\n", fileOut)
	code := generator.GenerateImplementationContract(fun)

	errWriteFile := ioutil.WriteFile(fileOut, []byte(code), fs.ModePerm)
	utils.ErrorWrapper(errWriteFile, fileOut)
}
