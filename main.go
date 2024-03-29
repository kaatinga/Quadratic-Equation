package QuadraticEquation

import (
	"math"
	"strconv"
	"strings"
)

const less0 = "There are no roots"

func Calc(a, b, c float64) (D, x1, x2 float64, conclusion string) {

	// Discriminant
	D = b*b - 4*a*c

	switch {
	case D < 0:
		conclusion = less0
	case D == 0:
		x1 = -b / 2 / a
		x2 = x1
		conclusion = strings.Join([]string{"x1 = x2 = ", strconv.FormatFloat(x1, 'g', 6, 64)}, "")
	case D > 0:
		sqrtD := math.Sqrt(D)
		x1 = (-b - sqrtD) / 2 / a
		x2 = (-b + sqrtD) / 2 / a
		conclusion = strings.Join([]string{"x1 = ", strconv.FormatFloat(x1, 'g', 6, 64), ", x2 = ", strconv.FormatFloat(x2, 'g', 6, 64)}, "")
	}

	return
}
