package main

import "fmt"

func main()  {
	max := 0

	for i := 0; i<10;i++{
		max = i
	}

	fmt.Print(max)

	sum := 1
	for ; sum < 100;{
		sum += sum

		print(sum)
	}
}
