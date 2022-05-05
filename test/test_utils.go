package test

import (
	. "byx-script-go/common"
	. "byx-script-go/interpreter"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func replace(src string, pattern string, replace string) string {
	r, _ := regexp.Compile(pattern)
	return r.ReplaceAllString(src, replace)
}

func replaceBlank(output string) string {
	output = replace(output, "\\r\\n", "\n")  // 将\r\n替换成\n
	output = replace(output, "^\\s+", "")     // 去除开头的空白字符
	output = replace(output, "\\s+\\n", "\n") // 将连续的空格换行替换成单个换行
	output = replace(output, "\\n\\s+", "\n") // 换行后的连续空白字符替换成单个换行
	output = replace(output, "\\s+$", "")     // 去除结尾的空白字符
	return output
}

func getOutput(executable func()) string {
	file, _ := ioutil.TempFile("", "tmp")
	oldStdout := os.Stdout
	os.Stdout = file

	defer func() {
		os.Stdout = oldStdout
	}()

	executable()

	bytes, _ := ioutil.ReadFile(file.Name())
	return string(bytes)
}

func importAndRunScript(importPaths []string, script string) string {
	return getOutput(func() {
		RunScript(script, RunConfig{ImportPaths: importPaths})
	})
}

func verify(t *testing.T, script string, expectedOutput string) {
	importAndVerify(t, []string{}, script, expectedOutput)
}

func importAndVerify(t *testing.T, importPaths []string, script string, expectedOutput string) {
	output := importAndRunScript(importPaths, script)
	assert.Equal(t, replaceBlank(expectedOutput), replaceBlank(output))
}

func readFile(filename string) string {
	content, err := ReadFileAsString(filename)
	if err != nil {
		panic(err)
	}
	return content
}

func runTestCase(t *testing.T, caseDir string) {
	importAndRunTestCase(t, []string{}, caseDir)
}

func importAndRunTestCase(t *testing.T, importPaths []string, caseDir string) {
	fmt.Println("===================== caseDir " + caseDir + " =====================")
	dir, err := ioutil.ReadDir(caseDir)
	if err != nil {
		panic(err)
	}
	for _, fi := range dir {
		name := fi.Name()
		if strings.HasSuffix(name, ".bs") {
			caseName := replace(name, "\\.bs", "")
			fmt.Println("case " + caseName + " running...")
			script := readFile(filepath.Join(caseDir, caseName+".bs"))
			expectedOutput := readFile(filepath.Join(caseDir, caseName+".out"))
			importAndVerify(t, importPaths, script, expectedOutput)
		}
	}
}
