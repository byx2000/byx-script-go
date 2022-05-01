package test

import (
	. "byx-script-go/src/interpreter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"regexp"
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
