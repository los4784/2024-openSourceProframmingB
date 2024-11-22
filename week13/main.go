package main

import (
	"fmt"
)

func main() {

	var emptySlice []bool
	// emptySlice = make([]bool, 5) slice zeor value (nil)
	fmt.Printf("%#v\n", emptySlice)

	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // 배열 초기화 시 중괄호 사용 atty := attay literal
	gpa_slice := gpas[1:4]                       // [4.1, 4.5, 3.9]
	// gpa_slice[1] = 2.71
	gpas[2] = 2.71
	// gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55) // 뒤 3.로 시작하는건 5개가 최대라 늘어나지 않는다 그럴때는 다른공간을 할당 받는다
	fmt.Println(len(gpa_slice), gpa_slice, gpas)
}
