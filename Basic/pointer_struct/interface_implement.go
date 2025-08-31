package main

type Mouth interface {
	Speak() string
	Eat() string
	ChangeName(newName string)
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return d.name + " says Woof!"
}

func (d *Dog) ChangeName(newName string) {
	d.name = newName
}

func (d *Dog) Eat() string {
	return d.name + " is eating."
}

func Test() {
	var m1 Mouth = &Dog{
		name: "Buddy",
	}

	println(m1.Speak())
	println(m1.Eat())
	m1.ChangeName("Max")

	println(m1.Speak())
	println(m1.Eat())
}
