package basicSyntax

import "fmt"

/*
Tokens in Go
In Go programming language, tokens are the smallest units of a program that have meaning.
They are the building blocks of the code and can be classified into several categories:

1. Keywords: These are reserved words that have a special meaning in Go.
Examples include "package", "import", "func", "var", "const", "if", "else", "for", "return", etc.

2. Identifiers: These are names used to identify variables, functions, types, and other user-defined entities.
Identifiers must start with a letter (A-Z or a-z) or an underscore (_) followed by letters, digits (0-9), or underscores.

3. Literals: These are the actual values assigned to variables.
They can be of various types such as
integer literals (e.g., 42),
floating-point literals (e.g., 3.14),
string literals (e.g., "Hello"),
and boolean literals (true, false).

4. Operators: These are symbols that perform operations on variables and values.
Examples include arithmetic operators (+, -, *, /),
relational operators (==, !=, <, >),
logical operators (&&, ||, !), and
assignment operators (=, +=, -=).

5. Punctuation: These are symbols that help structure the code.
Examples include commas (,), semicolons (;), parentheses (()), curly braces ({}), and square brackets ([]).
*/
func tokens() {
	fmt.Println("Hello, World!")
}

/*
	The individual tokens are −
	1. package
	2. tokens
	3. import
	4. "fmt"
	5. func
	6. tokens
	7. (
	8. )
	9. {
	10. fmt
	11. .
	12. Println
	13. (
	14. "Hello, World!"
	15. )
	16. }
*/
