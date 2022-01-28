package main

import "fmt"

type Enum int

const (
	Init    Enum = 0
	Success Enum = 1
	Faild   Enum = 2

	InitName    = "初始化"
	SuccessName = "成功"
	FaildName   = "失败"
)

func (e Enum) Int() int {
	return int(e)
}

func (e Enum) String() string {
	return []string{
		InitName,
		SuccessName,
		FaildName,
	}[e]
}
func main() {
	// fmt.Println(Init)
	// fmt.Println(Success)

	// status := 0
	// fmt.Println(Init == status)
	// fmt.Println(int(Init) == status)
	// fmt.Println(Init.Int() == status)
	// fmt.Println(Success) // fmt.Println会调用String方法

	type Counter map[string]int
	c := Counter{}
	c["wang"] = 1
	fmt.Println(c)

	type Queue []int
	q := make(Queue, 0)
	q = append(q, 1)
	fmt.Println(q)

}
