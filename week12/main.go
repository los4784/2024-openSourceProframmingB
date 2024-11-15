package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time
	//dates[0] = time.Unix(0, 0)
	//dates[2] = time.Unix(1708012345, 1)

	//fmt.Println(dates[0], dates[1], dates[2]) //  unix time zero value, 2024-02-16 ...

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012345, 0),
	}
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println(dates)          // 배열 찍음
	//fmt.Println("%#v\n", dates) // 배열 리터널 literall

	//for i := 0; i <= 2; i++  아래보다 안좋은 코드
	//for i := 0; i < len(dates); i++ {
	//	fmt.Println(i, dates[i])
	for i, date := range dates { // like python for in 안전하게 사용 가능
		fmt.Println(i, date)
	}
	for _, date := range dates { // 인덱스 놉
		fmt.Println(date)
	}

}
