package test

import (
	. "byx-script-go/interpreter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImport(t *testing.T) {
	importAndRunTestCase(t, []string{"./test_cases/import/p1", "./test_cases/import/p2"}, "./test_cases/import")

	assert.Panics(t, func() {
		RunScript(`
		import x

		println('main')`, RunConfig{ImportPaths: []string{"./test_cases/import/p3"}})
	})
}
