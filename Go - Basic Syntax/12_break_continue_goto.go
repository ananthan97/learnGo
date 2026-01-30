package basicSyntax

import "fmt"

/*
	break, continue and goto Statements in Go
	In Go, the break, continue, and goto statements
	are used to control the flow of execution in loops and other control structures.
	1. break Statement:
		The break statement is used to exit a loop or switch statement prematurely.
		When the break statement is encountered,
		the control is transferred to the statement immediately following the loop or switch.
		Example:
*/
func breakContinueGoto() {
	// Example of break statement
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Exit the loop when i equals 5
		}
		fmt.Println(i)
	}
}

/*	2. continue Statement:
	The continue statement is used to skip the current iteration of a loop
	and move to the next iteration.
	When the continue statement is encountered,
	the remaining code in the loop body for that iteration is skipped.
	Example:
*/
func continueStatement() {
	// Example of continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i) // Print only odd numbers
	}
}

/*	3. goto Statement:
	The goto statement is used to transfer control to a labeled statement within the same function.
	It allows for unconditional jumps in the code.
	However, the use of goto is generally discouraged
	as it can lead to unstructured and hard-to-read code.
	Example:
*/
func gotoStatement() {
	// Example of goto statement
	i := 1
start:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto start // Jump back to the start label
	}
}

/* Note:
While break and continue are commonly used in loops,
the use of goto should be limited to specific scenarios
where it enhances code clarity and maintainability.
Overusing goto can lead to "spaghetti code" that is difficult to follow.
*/
