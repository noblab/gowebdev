package main

import (
	"fmt"
)

type square struct {
	height int
	width  int
}

type circle struct {
	radius int
}

func (sq square) area() int {
	return sq.height * sq.width
}

func (c circle) area() int {
	return c.radius * c.radius * 3
}

type shape interface {
	area() int
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	var sq square
	var c circle
	sq.height = 4
	sq.width = 4
	c.radius = 3
	info(sq)
	info(c)
}
