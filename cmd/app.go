package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fanchann/coi/internal/generator"
	"github.com/fanchann/coi/internal/parser"
	"github.com/fanchann/coi/utils"
)

func AppRun(file string) {

	ext := filepath.Ext(file)
	if ext != ".go" {
		utils.ErrorWrapper(errors.New("file extension must go!"))
	}
	fun, err := parser.ExtractInterfaceFunctions(file)
	utils.ErrorWrapper(err)

	fileOut := fmt.Sprintf("%s/%s_impl.go", utils.GetFileLocation(file), utils.MakePrettyFileName(fun.IName))
	code := generator.GenerateImplementationContract(fun)

	errWriteFile := ioutil.WriteFile(fileOut, []byte(code), fs.ModePerm)
	utils.ErrorWrapper(errWriteFile, fileOut)

	log.Printf("[COI] Success generate,saved at %v \n", fileOut)
}
