package main

import "fmt"

func main()  {
	pow := make([]int,10)
	for i:= range pow{
		// << 二进制左移位运算符。左操作数值向左移动由右操作数指定的位数。
		pow[i] = 1 << uint(i)
	}

	for index,_ := range pow{
		fmt.Print(index)
	}
}
