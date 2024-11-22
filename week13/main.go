package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
// func test(strs ...string, i int){ 는 에러가 난다
func test(i int, strs ...string) {
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	//fmt.Print(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices, slices[1])
	test(1, "123")
	test(-99, "123", "abc")
	test(55)
	test(0, "123", "abc", "123a")
}
