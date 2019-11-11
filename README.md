# Quadratic-Equation
Quadratic Equation Calculation Package

An example of use:

```
package main

import (
	QuadraticEquation "github.com/kaatinga/Quadratic-Equation"
	"fmt"
	"log"
)

func main() {

	// Our parameters to read
	abc := make(map[int]float64, 3)

	fmt.Println("Welcome to Kaatinga's Quadratic Equation Calculator")
	fmt.Println("We are about to get unknowns (x) from a*(x^2) + b*x + c = 0")

	// Our parameters to print
	tmp := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	var floatTmp float64
	for key, value := range tmp {

		fmt.Println("Enter a float value", value, ":")
		for {
			_, err := fmt.Scan(&floatTmp)
			if err != nil {
				log.Println("Incorrect input. Repeat, please")
				continue
			}
			break
		}
		abc[key] = floatTmp
	}

	D, x1, x2, conclusion := QuadraticEquation.Calc(abc[1], abc[2], abc[3])

	fmt.Println(conclusion)
	fmt.Printf("The Quadratic Equation is %6f*(%6f^2) + %6f*%6f + %6f = 0\n", abc[1], x1, abc[2], x2, abc[3])
	fmt.Println("Discriminant:", D)

}
```
