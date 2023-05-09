package main

import (
	"fmt"
)

func main() {
	// array는 배열 크기 지정 필요
	numbers := [5]string{"1", "2", "3"}
	fmt.Println(numbers)
	numbers[3] = "4"
	numbers[4] = "5"
	fmt.Println(numbers)

	// 크기 제한 없는 slice
	names := []string{"bobby"}
	names = append(names, "lee")
	fmt.Println(names)
}