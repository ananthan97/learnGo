package basicSyntax

import "fmt"

func lineSeperator() {
	fmt.Println("Hello, World!")
	fmt.Println("I am in Go Programming World!")
	/*
		Line Separator:
		In a Go program, the line separator key is a statement terminator.
		That is, individual statements don't need a special separator like ; in C.
		The Go compiler internally places ; as the statement terminator to indicate the end of one logical entity.
		However, if you want to write multiple statements on a single line, you can use ; to separate them.
		For example:
		fmt.Println("Hello, World!"); fmt.Println("I am in Go Programming World!")

		In general practice, it's recommended to write each statement on a new line for better readability.
	*/
}
