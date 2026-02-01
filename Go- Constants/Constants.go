package Constants

import "fmt"

/*
This package demonstrates the use of constants in Go.
Constants are immutable values that are defined at compile time and cannot be changed during the execution of the program.
They are declared using the 'const' keyword.
*/
func ConstantsDemo() {

	// Constant Declaration
	const pi float64 = 3.14159
	const greeting string = "Hello, Go Constants!"
	const isGoFun bool = true
	fmt.Println("Constant pi:", pi)
	fmt.Println("Constant greeting:", greeting)
	fmt.Println("Constant isGoFun:", isGoFun)

	// Multiple Constant Declaration
	const (
		e        float64 = 2.71828
		language string  = "Go"
		version  int     = 1
	)
	fmt.Println("Constant e:", e)
	fmt.Println("Constant language:", language)
	fmt.Println("Constant version:", version)

	// Untyped Constants
	const untypedInt = 42
	const untypedFloat = 3.14
	const untypedString = "Untyped Constant"
	fmt.Println("Untyped Constant int:", untypedInt)
	fmt.Println("Untyped Constant float:", untypedFloat)
	fmt.Println("Untyped Constant string:", untypedString)

	// Constant Expressions
	const radius = 5
	const area = pi * radius * radius
	fmt.Println("Area of circle with radius", radius, "is:", area)

	// iota Example
	const (
		_              = iota // ignore first value by assigning to blank identifier
		FirstConstant  = iota // 1
		SecondConstant        // 2
		ThirdConstant         // 3
	)
	fmt.Println("FirstConstant:", FirstConstant)
	fmt.Println("SecondConstant:", SecondConstant)
	fmt.Println("ThirdConstant:", ThirdConstant)
}
