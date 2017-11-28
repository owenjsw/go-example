package main

import "fmt"

func main()  {

	a := []byte{1,2,3,4,5}
	b := []byte{'a','b'}
	AppendSlice(b,a)
}

func AppendSlice(slice []byte,data []byte){
	m :=len(slice)
	n :=m+len(data)
	if n > cap(slice){
		newSlice := make([]byte,(n+1)*2)
		copy(newSlice,slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n],data)
	fmt.Printf("slice",slice)
}