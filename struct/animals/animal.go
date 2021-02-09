package main

import "fmt"

type IFly interface {
	fly() bool
}

type IWalk interface {
	walk() bool
	run(speed int)
}

type NameSetter interface {
	setName(name string)
}

type animal struct {
	Name string
	id int
}

type Dog struct {
	Name string
}

type Bird struct {
	Name string
}

type Car struct {
}

func (a Dog)walk() bool {
	fmt.Println(a.Name + " can walk")
	return true
}

func (a Dog)run(speed int)  {
	fmt.Printf("%s run at speed %d\n", a.Name, speed)
}

func (b Bird)fly() bool {
	fmt.Println(b.Name + " can fly")
	return true
}

func (c Car)run(speed int)  {
	fmt.Printf("car run at spped %d\n", speed)
}

func (a animal)fly() bool {
	fmt.Printf("%s cant fly\n", a.Name)
	return false
}

func (a *animal)setName(name string)  {
	a.Name = name
}

func DoTest()  {
	a := animal{"fish", 1}
	a.setName("gg fish")
	a.fly()

	var na NameSetter
	na = &a
	na.setName("ok fish")
	a.fly()

	var f IFly
	var w IWalk
	f = Bird{"bird"}
	w = Dog{"dog"}

	f.fly()
	w.walk()
	w.run(3)
}

func main()  {
	DoTest()
}