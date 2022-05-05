package test

import (
	"testing"
)

func TestStdlib(t *testing.T) {
	importAndRunTestCase(t, []string{"./stdlib"}, "./test_cases/stdlib")
}
