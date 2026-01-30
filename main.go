package main

//1. Package Declaration
/* This line declares the package name for the Go file.
In Go, every file belongs to a package, and the 'main' package is special because it defines a standalone executable program.
When you run a Go program, the Go runtime looks for the 'main' package and executes the 'main' function within it. */
import "fmt"

//2. Import Statement
/* This line imports the "fmt" package, which is part of Go's standard library.
The "fmt" package provides functions for formatted I/O operations, such as printing to the console.
By importing this package, we can use its functions in our code. */
func main() {
	//3. Main Function
	/* This line defines the 'main' function, which is the entry point of a Go program.
	   When you run a Go program, execution starts from the 'main' function.
	   Every standalone Go application must have a 'main' function in the 'main' package. */
	fmt.Println("Hello, World!")
	//4. Print Statement
	/* This line calls the 'Println' function from the 'fmt' package to print the string "Hello, World!" to the console.
	   'Println' stands for "print line" and it outputs the provided arguments followed by a newline character.
	   In this case, it prints the classic greeting message used in programming tutorials.

	   Notice the capital P of Println method.
	   In Go language, a name is exported if it starts with capital letter.
	   Exported means the function or variable/constant is accessible to the importer of the respective package.
	*/

}
