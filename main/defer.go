package main

import "fmt"

func test()(x int){
	x =1
	return x
}

func main()  {
	defer fmt.Print("abc")

	fmt.Print(test())
}