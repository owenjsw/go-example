package main

import "fmt"


type Person struct {
	Name string
	Age int
}

type IPAddr [4]byte
func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


func (ip IPAddr) String() string{
	return fmt.Sprintf("%v .",ip.String())
}
func main(){
	a := Person{"test",42}
	b := Person{"test2",43}

	fmt.Println(a,b)

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}