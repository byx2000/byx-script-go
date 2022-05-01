package parser

import "testing"

func TestRange(t *testing.T) {
	p := Range('0', '9')

	r, err := p.ParseToEnd("0")
	if err != nil || r != '0' {
		t.Errorf("TestRange failed")
	}

	r, err = p.ParseToEnd("5")
	if err != nil || r != '5' {
		t.Errorf("TestRange failed")
	}

	r, err = p.ParseToEnd("9")
	if err != nil || r != '9' {
		t.Errorf("TestRange failed")
	}

	r, err = p.ParseToEnd("*")
	if err == nil {
		t.Errorf("TestRange failed")
	}

	r, err = p.ParseToEnd("a")
	if err == nil {
		t.Errorf("TestRange failed")
	}

	r, err = p.ParseToEnd("")
	if err == nil {
		t.Errorf("TestRange failed")
	}
}

func TestString(t *testing.T) {
	p := String("abc")

	r, err := p.ParseToEnd("abc")
	if err != nil || r != "abc" {
		t.Errorf("TestString failed")
	}

	r, err = p.ParseToEnd("abx")
	if err == nil {
		t.Errorf("TestString failed")
	}

	r, err = p.ParseToEnd("ab")
	if err == nil {
		t.Errorf("TestString failed")
	}

	r, err = p.ParseToEnd("axy")
	if err == nil {
		t.Errorf("TestString failed")
	}

	r, err = p.ParseToEnd("a")
	if err == nil {
		t.Errorf("TestString failed")
	}

	r, err = p.ParseToEnd("")
	if err == nil {
		t.Errorf("TestString failed")
	}
}
