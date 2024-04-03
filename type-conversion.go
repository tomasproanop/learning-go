#Type conversion

package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var x, y int = 3, 4
	var w, v int = 5, 6

	var f float64 = math.Sqrt(float64(x*x + y*y)) 
	var p float32 = (float32(w+v))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(z))
}

#Output
#3 4 5
#11
#float32
#float64
#uint
