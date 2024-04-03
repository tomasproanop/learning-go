package main

import "fmt"

func main() {
	pi := 3.1415 
	fmt.Printf("v is of type %T\n", v)
}


/*
"When the right hand side of the declaration is typed, 
the new variable is of that same type.

But when the right hand side contains an untyped numeric constant, 
the new variable may be an int, float64, or complex128 depending on the precision of the constant."
*/
