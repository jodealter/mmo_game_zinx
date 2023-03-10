package main

import "fmt"

type Point struct {
	x int
	y int
}

/*
下边的p Point与p *Point 是接收器
*/
func (p Point) name() {
	fmt.Println("name")
}
func (p *Point) test() {
	fmt.Println("test")
}

func main() {
	cache := struct {
		*Point
	}{&Point{
		x: 0,
		y: 0,
	}}
	p := &Point{
		x: 0,
		y: 0,
	}

	//cache.test 是选择器
	cache.test()
	p.test()
	//这个叫做方法“值”
	var t = cache.test
	t()
	p2 := Point{
		x: 0,
		y: 0,
	}
	p3 := &Point{
		x: 0,
		y: 0,
	}

	//下边这个叫做方法表达式
	s := Point.name // 这样指定了s是这个类型的方法，但是并没有指定具体的接收器，只是指定了接收器的类型
	s(p2)
	s2 := (*Point).test
	s2(p3)
}
