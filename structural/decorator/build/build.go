package build

import (
	"fmt"
)

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		fmt.Println("Starting the execution with the integer", n)
		n += 11
		result := fn(n)
		fmt.Println("Execution is completed with the result", result)
		return result
	}
}
