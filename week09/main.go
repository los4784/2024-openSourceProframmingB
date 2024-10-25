package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1 ~ 6
	fmt.Println(answer)

	fmt.print("숫자 입력: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString("/n")
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // 줄바꿈, 띄엇쓰기 탭 등 제거 python strip과 유사
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답입니다")
	} else if answer > guess {
		fmt.Println("입력하신 값은 정답보다 작은 수 입니다. LOW")
	} else { // answer < guess
		fmt.Println("입력하신 값은 정답보다 큰 수 입니다.")
	}
}
