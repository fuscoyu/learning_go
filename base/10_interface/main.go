package main

import "fmt"

type Sleeper interface {
	Sleep()
}

type Eater interface {
	Eat(foodName string)
}

type LazyAnimal interface {
	Sleeper
	Eater
}

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}
func (d Dog) Eat(foodName string) {
	fmt.Printf("Dog %s is eating %s\n", d.Name, foodName)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}
func (c Cat) Eat(foodName string) {
	fmt.Printf("Cat %s is eating %s\n", c.Name, foodName)
}

func AnimalSleep(s Sleeper) {
	s.Sleep()
}

func test() {
	var s Sleeper
	d := Dog{Name: "小虎"}
	c := Cat{Name: "咪咪"}
	s = d
	AnimalSleep(s)
	s = c
	AnimalSleep(s)

	sleepList := []Sleeper{Dog{Name: "wangwang"}, Cat{Name: "miaomiao"}}
	for _, s := range sleepList {
		s.Sleep()
	}

}
func main() {
	alist := []LazyAnimal{Dog{Name: "dog"}, Cat{Name: "cat"}}
	for _, a := range alist {
		a.Sleep()
		a.Eat("肉")
		//类型断言
		if dog, ok := a.(Dog); ok {
			fmt.Printf("I am a dog, my name is %s", dog.Name)
		}
		if cat, ok := a.(Cat); ok {
			fmt.Printf("I am a cat, my name is %s", cat.Name)
		}
	}

}
