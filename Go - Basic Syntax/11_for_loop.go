package basicSyntax

import "fmt"

/*
	for Loop in Go
	The for loop is used to execute a block of code repeatedly for a specified number of times or
	while a certain condition is true.
	The syntax of a for loop in Go is as follows:
	for initialization; condition; post {
	    // code to be executed in each iteration
	}
	- initialization: This step is executed once at the beginning of the loop. It is typically used to initialize a loop counter variable.
	- condition: This is a boolean expression that is evaluated before each iteration of the loop. If the condition is true, the loop body is executed; if false, the loop terminates.
	- post: This step is executed after each iteration of the loop. It is typically used to update the loop counter variable.
	Example:
*/
func forLoop() {
	// Example: Print numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// Example: Sum of first 10 natural numbers
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural numbers:", sum)
	// Example: Using for loop as a while loop
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}
	// Example: Infinite loop (uncomment to run)
	// for {
	//     fmt.Println("This is an infinite loop")
	// }

	//nested for loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
