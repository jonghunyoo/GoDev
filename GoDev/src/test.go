package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func TestString(x string) string{

	return x
}

func main(){
	//fmt.Printf("hello, world\n")

	//fmt.Println(add(2, 1))

	fmt.Printf(TestString("aaa"))
}
