package main

import "fmt"

func changeStr(s string) {
	s = "wang"
}

func changeMap(m map[string]string) {
	m["wang"] = "li"
}

func changeStrByPtr(s *string) {
	*s = "new wang"
}

//高级函数
func testFuncType() {
	myPrint := func(s string) { fmt.Println(s) }
	myPrint("nihao")
}

func testMapFunc() {
	funcMap := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}

	fmt.Println(funcMap["add"](1, 2))
	fmt.Println(funcMap["sub"](5, 2))
}
func filterIntSlice(intVals []int, predicate func(i int) bool) []int {
	res := make([]int, 0)
	for _, val := range intVals {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

func main() {
	str := "rong"
	changeStr(str) //值传递
	fmt.Println(str)

	m := map[string]string{"name": "jia"} // 值传递  map的底层实现带有指针所以可以修改
	changeMap(m)
	fmt.Println(m)

	name := "li"
	changeStrByPtr(&name)
	fmt.Println(name)

	ints := []int{1, 2, 3, 4, 5}
	isEven := func(i int) bool { return i%2 == 0 }
	fmt.Println(filterIntSlice(ints, isEven))

	testFuncType()

	testMapFunc()

  fmt.Println(fib(10))

}
