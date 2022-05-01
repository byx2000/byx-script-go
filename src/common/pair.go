package common

import "fmt"

type Pair struct {
	First  any
	Second any
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}
