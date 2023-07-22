package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	PetName string
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleep\n", a.Name)
}

func (a *Animal) SetPetName(PetName string) {
	// 下面两种方式都可以
	// a.PetName = PetName //go 提供的语法糖
	(*a).PetName = PetName
}

func (a Animal) GetPetName() string {
	return a.PetName
}

func NewAnimal(name string, age int) *Animal {
	a := Animal{
		Name: name,
		Age:  age,
	}
	return &a
}

type Dog struct {
	Animal //继承Animal的所有方法属性
	color  string
}

func (d Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func main() {
	a := Animal{Name: "cat", Age: 3}
  fmt.Println(a.Name)
	// fmt.Println(a, a.Name, a.Age)
	// a.Sleep()

	aPrt := &Animal{Name: "dog", Age: 4}
  fmt.Println(aPrt.Name)
	// aPrt.SetPetName("little dog")
	// fmt.Println(aPrt.GetPetName())
	//
	// a1 := NewAnimal("pig", 2)
	// fmt.Println(a1)
	//
	// d := Dog{}
	// d.Name = "xiaohu"
	// d.Sleep()

}
