package QuadraticEquation

import (
	"strconv"
	"testing"
)

func TestCalc(t *testing.T) {

	_, x1, x2, _ := Calc(2, 10, 2)

	if strconv.FormatFloat(x1, 'g', 6, 64) != "-4.79129" || strconv.FormatFloat(x2, 'g', 6, 64) != "-0.208712" {
		t.Error("error")
	}
}
