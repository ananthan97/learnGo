package basicSyntax

import "fmt"

/*
	Switch Statement in Go
	The switch statement is used for conditional execution of code blocks based on the value of an expression.
	It provides a way to select one of many code blocks to be executed.
	The syntax of a switch statement in Go is as follows:
	switch expression {
		case value1:
			// code to be executed if expression == value1
		case value2:
			// code to be executed if expression == value2
		...
		default:
			// code to be executed if expression doesn't match any case
	}
	The switch expression is evaluated once, and its value is compared against the values specified in each case.
	If a match is found, the corresponding code block is executed.
	The default case is optional and is executed if none of the cases match the expression's value.
	Go's switch statement automatically breaks after executing a case block, so there's no need for explicit break statements.
	You can also have multiple values in a single case by separating them with commas.
	Example:
*/
func switchStatement() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
	// Example with multiple values in a case
	grade := "B"
	switch grade {
	case "A", "B", "C":
		fmt.Println("Good grade")
	case "D", "E":
		fmt.Println("Average grade")
	default:
		fmt.Println("Invalid grade")
	}
}
