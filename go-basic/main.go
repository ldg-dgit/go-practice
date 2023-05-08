package main

import (
	"fmt"
)

func canIDrink(age int) bool{
	// variable expression
	// if를 사용함과 동시에 변수를 생성할 수 있다. if 밖에서 생성하는 것과는 같다. 가시성을 올린다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(15))
}