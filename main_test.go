package QuadraticEquation

import (
	"testing"
)

func TestCalc(t *testing.T) {

	_, x1, x2, _ := Calc(2, 10, 2)

	if x1 == 0 || x2 == 0 {
		t.Error("error")
	}

}
