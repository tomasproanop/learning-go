package main

import "fmt"

func main() {
	i, j, k := 42, 2701, 250

	p := &i         // point to i
	fmt.Println("this is i through the pointer:", *p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println("this is the new value of j:", j) // see the new value of j
	
	p = &k
	*p = *p / 2
	fmt.Println("this is the new value of k:", k)
}


/*
output:
this is i through the pointer: 42
21
this is the new value of j: 73
this is the new value of k: 125
*/
