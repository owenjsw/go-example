package main

import (
	"fmt"
)

type MAP struct {
	a,b float64
}

var m map[string]MAP
func main()  {

	m = make(map[string]MAP)
	m["test"] = MAP{
		40.1,40.2,
	}
	fmt.Println(m)

	m["test"] = MAP{
		42,43,
	}
	fmt.Println(m)

	v,ok := m["test"]

	fmt.Println(v,ok)
}
