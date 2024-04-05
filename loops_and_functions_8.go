//newton's method



package main

import (
	"fmt"
)

// first implementation
func Sqrt(x float64) float64 {
	
	z := 1.0
	
	fmt.Printf("Square root approximation of %v:\n", x)
	
	for i := 1; i <= 10; i++{
		z -= (z * z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}


//second implementation

func Sqrtb(x float64) float64{
	z := float64(1)
	var t float64
	
	for {
		z, t = z - (z * z - x) / (2 * z), z
		if math.Abs(t - z) < 1e-6 {
			break
		}
	}
	return z
}


func main() {
	fmt.Println(Sqrt(2))
	
	guess := Sqrtb(2)
    expected := math.Sqrt(2)
    fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess - expected))
} 
