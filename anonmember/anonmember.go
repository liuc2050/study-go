package main

import (
	"fmt"
	"time"
)

type point struct {
	X int
	Y int
}

type circle struct {
	point
	Radius int
}

func (c *circle) ccc() {
	c.Radius = 8
}

type Cccer interface {
	ccc()
}

type Wheel struct {
	circle
	Spokes int
}

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Wheel
	w.ccc()
	fmt.Println(w.Radius)
	fmt.Println(time.Now().Format(time.UnixDate))
}
