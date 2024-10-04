package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9 // f:= 12.9
	c1 := '2'
	c2 := '김'

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f) // invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %f\n", i, f, i*int(f)) // conversion
	fmt.Println(c1, c2)
	fmt.Printf("%x\n", c2)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
