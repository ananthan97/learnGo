package Variables

import "fmt"

/* This package demonstrates variable declaration and initialization in Go.
   Variables in Go can be declared using the 'var' keyword or using the shorthand ':=' operator.
   Variables can also be initialized with values at the time of declaration.
*/
func VariablesDemo() {
	// Variable Declaration with var keyword
	var a int = 10
	var b string = "Hello, Go!"
	var c float64 = 3.14

	// Print the variables
	fmt.Println("Variable a (int):", a)
	fmt.Println("Variable b (string):", b)
	fmt.Println("Variable c (float64):", c)

	// Shorthand Variable Declaration
	d := 20
	e := "Goodbye, Go!"
	f := 2.71828

	// Print the shorthand variables
	fmt.Println("Variable d (int):", d)
	fmt.Println("Variable e (string):", e)
	fmt.Println("Variable f (float64):", f)

	// Multiple Variable Declaration
	var x, y, z int = 1, 2, 3
	fmt.Println("Variables x, y, z (int):", x, y, z)

	// Multiple Shorthand Variable Declaration
	m, n, o := 4, 5, 6
	fmt.Println("Variables m, n, o (int):", m, n, o)

	// Zero Value Declaration
	var p int
	var q string
	var r float64
	fmt.Println("Zero Value of p (int):", p)     // prints 0
	fmt.Println("Zero Value of q (string):", q)  // prints empty string
	fmt.Println("Zero Value of r (float64):", r) // prints 0.0

	// Constant Declaration
	const pi float64 = 3.14159
	fmt.Println("Constant pi (float64):", pi)

	// Shadowing Example
	var shadow int = 100 // outer variable
	{
		var shadow int = 200 // inner variable shadows outer variable
		fmt.Println("Inner shadow variable:", shadow)
	}
	fmt.Println("Outer shadow variable:", shadow)

	// lvalues and the rvalues in Go
	var lvalue int = 50 // lvalue is the variable 'lvalue', rvalue is 50
	fmt.Println("Lvalue variable:", lvalue)

	// Note: This is a basic demonstration of variable declaration and initialization in Go.
	// More complex usage and features can be explored further.
	// End of VariablesDemo function
}
