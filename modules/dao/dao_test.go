package dao

import (
	"testing"
)

func TestPlus(t *testing.T) {
	x := 1
	y := 2
	if Plus(x, y) != 3 {
		t.Fail()
	}
}