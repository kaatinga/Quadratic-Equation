package QuadraticEquation

import (
	"strconv"
	"testing"
)

func TestCalc1(t *testing.T) {

	_, x1, x2, _ := Calc(2, 10, 2)

	if strconv.FormatFloat(x1, 'g', 6, 64) != "-4.79129" || strconv.FormatFloat(x2, 'g', 6, 64) != "-0.208712" {
		t.Error("error 1")
	}
}
func TestCalc2(t *testing.T) {
	D, _, _, _ := Calc(2, 4, 2)

	if D != 0 {
		t.Error("error 2")
	}

}

func TestCalc3(t *testing.T) {
	_, _, _, conclusion := Calc(2, 2, 2)

	if conclusion != less0 {
		t.Error("error 3")
	}

}
