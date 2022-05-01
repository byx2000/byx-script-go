package common

import "io/ioutil"

func ReadFileAsString(filename string) (string, error) {
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		return "", e
	}
	return string(bytes), nil
}
