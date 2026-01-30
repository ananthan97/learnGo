package basicSyntax

import "fmt"

/*
	If Statement in Go
	The if statement is used for conditional execution of code blocks
	based on whether a specified condition evaluates to true or false.
	The syntax of an if statement in Go is as follows:

	if condition {
	    // code to be executed if the condition is true
	}

	You can also include an optional else clause to execute code when the condition is false:
	if condition {
	    // code to be executed if the condition is true
	} else {
	    // code to be executed if the condition is false
	}

	You can also use else if to check multiple conditions:
	if condition1 {
	    // code to be executed if condition1 is true
	} else if condition2 {
	    // code to be executed if condition2 is true
	} else {
	    // code to be executed if both conditions are false
	}

	nested if statements are also allowed, where an if statement is placed inside another if or else block.
	Example:
*/
func ifStatement() {
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}
	// Nested if example
	if num > 0 {
		if num%2 == 0 {
			fmt.Println(num, "is a positive even number.")
		} else {
			fmt.Println(num, "is a positive odd number.")
		}
	} else {
		fmt.Println(num, "is not a positive number.")
	}
}
