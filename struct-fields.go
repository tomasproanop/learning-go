package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z int
}

func main() {
	v := Vertex{1, 2, 3}
	v.X = 4
	fmt.Println(v.X)
	v.Y = 5
	fmt.Println(v.Y)
	fmt.Println(v.Z)
  v.Z = 6
  fmt.Println(v.Z)
}


/*
output:
4
5
3
6
*/
