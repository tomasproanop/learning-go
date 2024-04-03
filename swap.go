//Multiple results

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swap2(c,d int) (int, int){
return d, c}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	c, d := swap2(1, 10)
	fmt.Println(c,d)
}
