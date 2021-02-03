// v0.1.4
// Author: DIEHL E.
// (©) Sony Pictures Entertainment, Feb 2021

package main

import (
	"fmt"
	"os"

	"github.com/wunderbarb/mygot/internal/common"
	"github.com/wunderbarb/mypkg/toolbox"
)

var fh *os.File

func main() {
	fmt.Println("v0.1-02-02")
	if len(os.Args) < 3 {
		fmt.Println("help testgo fn file")
		os.Exit(1)
	}

	alreadyExisted := true
	// the file name (args 2) may be with or without .go
	sourceName := toolbox.SetExtension(os.Args[2], "go")
	fileName := toolbox.SetExtension(toolbox.Strip(os.Args[2], "go")+"_test", "go")

	if !toolbox.IsExist(fileName) {
		alreadyExisted = false
		err := common.CreateHeader(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}

	var err error
	fh, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer fh.Close()

	_, p, err := analyzeFunction(os.Args[1], sourceName)
	fmt.Println(p)

	if !alreadyExisted {
		buildImport() // only for the first time
	}

	buildTestCase(os.Args[1])
}

func buildImport() {
	write(0, "import (")
	write(1, "\"testing\"")
	emptyLine()
	write(1, "\"github.com/wunderbarb/test\"")
	write(0, ")")
	emptyLine()
}

func buildTestCase(fn string) {
	write(0, "func Test_"+fn+"(t *testing.T){")
	write(1, "require, assert := test.Describe(t)")
	emptyLine()
	write(1, "tests := []struct {")
	write(2, "expRes xxx")
	write(2, "expSuccess bool")

	write(1, "}{")
	write(2, "{ xx, true },")
	write(2, "{ xx, false },")
	write(1, "}")
	write(1, "for i, tt := range tests {")
	write(2, "res, err := "+fn+"()")
	write(2, "require.Equal(tt.expSuccess, err==nil, \"sample %d\", i+1)")
	write(2, "if err == nil {")
	write(3, "assert.Equal(tt.expRes, res)")
	write(2, "}")
	write(1, "}")
	write(0, "}")
	emptyLine()
}

func write(tabs int, text string) {
	s := ""
	for i := 0; i < tabs; i++ {
		s += "\t"
	}
	_, err := fh.WriteString(s + text + "\n")
	if err != nil {
		panic(err)
	}
}

func emptyLine() {
	fh.WriteString("\n")
}

func analyzeFunction(fn string, fileName string) (string, []string, error) {
	p, _, err := common.FindFunctionDeclaration(fn, fileName)
	if err != nil {
		return "", nil, err
	}
	return "", p, nil
}
