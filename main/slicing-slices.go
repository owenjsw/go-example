package main

import "fmt"

func main()  {
	s := []int{1,2,3,4,5}

	fmt.Println("s ==",s)

	fmt.Println("s[1:4]",s[1:4])

	a := make([]int,4)
	//0值，5容量
	//b := make([]int,0,5);
	printSlice("a",a)
}
//cap 返回切片容量
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


//切片详解  https://blog.go-zh.org/go-slices-usage-and-internals



