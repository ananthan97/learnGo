package dataTypes

import "fmt"

// This package demonstrates various data types in Go.
/* Types of Go Data Types:
1. Boolean Types
2. Numeric Types
   a. Integer Types
   b. Floating Point Types
   c. Complex Types
3. String Type
4. Derived Types
   (a) Pointer types
   (b) Array types
   (c) Structure types
   (d) Union types
   (e) Function types
   (f) Slice types
   (g) Interface types
   (h) Map types
   (i) Channel Types
*/
func DataTypesDemo() {

	// Boolean Type
	var b bool = true
	fmt.Println("Boolean Value:", b)

	// Integer Types

	// Signed Integer Types: int, int8, int16, int32, int64
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Println("Integer Values:", i, i8, i16, i32, i64)

	// Unsigned Integer Types: uint, uint8, uint16, uint32, uint64
	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	fmt.Println("Unsigned Integer Values:", ui, ui8, ui16, ui32, ui64)

	// Floating Point Types: float32, float64
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	fmt.Println("Floating Point Values:", f32, f64)

	// Complex Types: complex64, complex128
	var c64 complex64 = complex(1, 2)
	var c128 complex128 = complex(3, 4)
	fmt.Println("Complex Values:", c64, c128)

	//Other Numeric Types: byte, rune, uintptr
	var by byte = 255
	var r rune = 'A'
	var up uintptr = 12345678
	fmt.Println("Other Numeric Values:", by, r, up)

	// String Type
	var s string = "Hello, Go!"
	fmt.Println("String Value:", s)

	// Derived Types

	// Pointer Type
	var p *int = &i
	fmt.Println("Pointer Value:", *p)

	// Array Type
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array Value:", arr)

	// Slice Type - dynamic array
	var slc []int = []int{4, 5, 6}
	fmt.Println("Slice Value:", slc)

	// Map Type -Hash Table
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println("Map Value:", m)

	// Channel Type
	ch := make(chan int) //chan is keyword to declare channel type, make is used to initialize channel
	//make is used to allocate and initialize objects like slices, maps, and channels.
	go func() { ch <- 7 }()
	fmt.Println("Channel Value:", <-ch)

	//Structure Type
	type Person struct {
		Name string
		Age  int
	}
	var person Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Structure Value:", person)

	// Interface Type - a type that specifies a set of method signatures.
	var inter interface{} = "I am an interface"
	fmt.Println("Interface Value:", inter)

	//Union Types are not directly supported in Go, but can be simulated using interfaces.
	//Example:
	type Union interface{} // type that can hold any type, Union is an alias for empty interface
	var u Union = 42
	fmt.Println("Union Value:", u)

	//function types
	var funcType func(int) int = func(x int) int { return x * x }
	fmt.Println("Function Type Value:", funcType(5)) // prints 25

	// Note: This is a basic demonstration of Go data types.
	// More complex usage and features can be explored further.
	// End of DataTypesDemo function
}
