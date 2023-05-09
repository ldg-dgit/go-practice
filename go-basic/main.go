package main

import (
	"fmt"
)

func main() {
	//map
	//명시한 자료형 내에서만 작성할 수 있는 object 느낌
	bobby := map[string]string{"name":"bobby", "age":"19"}
	fmt.Println(bobby)

	for key, value := range bobby {
		fmt.Println(key, value)
	}

	//
	for _, value := range bobby {
		fmt.Println(value)
	}
}