package main

import (
	"fmt"
)

func main()  {
	fmt.Print(splite(17))
}

func splite(sum int)(x,y int)  {
	x = sum*4/9
	y = sum - x
	return
}