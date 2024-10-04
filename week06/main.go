package main

import (
   "fmt"
)
Func main() {
	i := 13
	f := float64 = 12.9 // f:= 12.9
	
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f) // invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %f\n", i, f, i*int*f) // conversion
	fmt.Printf(reflect.TypeOf(f),reflect.TypeOf(i))
}