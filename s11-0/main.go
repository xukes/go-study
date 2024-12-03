package main

import (
	"fmt"
)

type Player interface {
	play() string
}
type Mover interface {
	move()
}
type DDer interface {
	Player
	Mover
}

type Pianist struct {
	name string
	age  int
	p    *Player
}

func (p *Pianist) play() string {
	return "pianist"
}
func (p *Pianist) move() {
	fmt.Println("Pianist move")
}

type Drummer struct {
	name string
}

func (p *Drummer) play() string {
	return "drummer"
}

func main() {
	dr := &Pianist{}

	dd := new(Pianist)
	dd.name = "xxx"
	fmt.Println(dd.name)
	dd1 := &Pianist{name: "xsd"}
	vs := (dd1).name
	fmt.Printf(vs)
	n := 2

	var arr1 [4]int

	slice13 := make([]int, n)
	fmt.Println(slice13)
	fmt.Println(arr1)
	fmt.Println(getName(dr))
	justifyType(dr)
}
func getName(p DDer) string {
	p.move()
	return p.play()
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case Pianist:
		fmt.Println("the name is Pianist %", v)
	}
}
