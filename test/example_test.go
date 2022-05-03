package test

import (
	. "byx-script-go/common"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func readFile(filename string) string {
	content, err := ReadFileAsString(filename)
	if err != nil {
		panic(err)
	}
	return content
}

func TestExample(t *testing.T) {
	exampleDir := "example"
	dir, err := ioutil.ReadDir(exampleDir)
	if err != nil {
		panic(err)
	}
	for _, fi := range dir {
		name := fi.Name()
		if strings.HasSuffix(name, ".bs") {
			caseName := replace(name, "\\.bs", "")
			fmt.Println("case " + caseName + " begin")
			script := readFile(filepath.Join(exampleDir, caseName+".bs"))
			expectedOutput := readFile(filepath.Join(exampleDir, caseName+".out"))
			verify(t, script, expectedOutput)
			fmt.Println("case " + caseName + " end")
		}
	}
}
