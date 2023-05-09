package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string

	// 여기 메소드를 쓰면 그게 그거다
	// 생성자 constructor는 없다.
}

func main() {
	bobby := person{"bobby", 19, []string{"beaf", "kimchi"}}
	fmt.Println(bobby)
	fmt.Println(bobby.age)

	bobby_2 := person{name: "bobby", age: 19, favFood: []string{"beaf", "kimchi"}}
	fmt.Println(bobby_2)
}