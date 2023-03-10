package main

import "fmt"

type IntSet []uint64

func (i *IntSet) test() {
	fmt.Println("jfd")
}
func (i *IntSet) String() string {
	return "yes"
}
func main() {
	s := IntSet{}
	(&s).test()
	s[0] = 9
	s.test()
}
