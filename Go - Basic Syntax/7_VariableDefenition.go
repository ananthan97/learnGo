package basicSyntax

import "fmt"

/*
	Variable Definition in Go
	Variables in Go are used to store data values. Before using a variable, it must be defined with a specific type.
	There are several ways to define variables in Go:
	1. Using the var keyword:
		var variableName dataType
	Example:
		var age int
	2. Defining and initializing a variable in one line:
		var variableName dataType = value
	Example:
		var name string = "John"
	3. Using type inference (the compiler infers the type based on the assigned value):
		var variableName = value
	Example:
		var isActive = true
	4. Short variable declaration (only inside functions):
		variableName := value
	Example:
		count := 10
	You can also define multiple variables at once:
		var x, y, z int
		var a, b, c = 1, 2, 3
	Variables can be defined at the package level (global scope) or within functions (local scope).
*/
func variableDefinition() {
	// Example of variable definition and initialization
	var age int = 25
	var name string = "Alice"
	isActive := true
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Active:", isActive)
}
