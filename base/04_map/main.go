package main

import (
	"fmt"
	"strings"
)

func testMap(){
	//var m map[string]int 
	// var m map[string]int{}
	// var m map[string]int{"hello":1, "go":2}
	m := make(map[string]int)
	m["zhang"] = 1
	m["wang"] = 2
	m["li"] = 3
	fmt.Println(m)
	// 删除map中key  key不存在也不会报错
	delete(m,"li") 
	fmt.Println(m)

	// 判断key是否存在 map 查不到key也不会报错
	if v,ok := m["wang"]; ok {
		fmt.Printf("m[%s] is %d \n", "wang", v)
	}

	m["lu"] = 4
	m["liu"] = 7
	m["jia"] = 6

	// 遍历map map是没有顺序的
	for k,v := range m {
		fmt.Printf("m[%s]: %d\n",k,v)
	}

	// 取map中所有key
	for k := range m{
		fmt.Println(k)
	}
	// 取map中所有的vaule
	for _,v:= range m{
		fmt.Println(v)
	}


}

func wordVectors(){
	str:="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
 	count := make(map[string]int)
        	
	//分词
	substr:= strings.Split(str," ")
	fmt.Println(substr)

	for _,value := range substr{
		count[value]++
	}
	for k,v:=range count{
		fmt.Printf("%s:%d\n",k,v)
	}



}
func main(){
	// testMap()
	wordVectors()
}
