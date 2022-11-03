package main

import (
  "fmt"
  "sort"
)

func testArray(){
	var arrayInt64 [3]int64 //声明
	arrayInt64[0],arrayInt64[1],arrayInt64[2] = 0,1,2 //初始化
	fmt.Println(arrayInt64) //[0 1 2]

	arrayString := []string{"hello","go"}
	fmt.Println(arrayString) // ["hello" "go"]

	arrayFloat := [...]float64{1,2,3}
	fmt.Println(arrayFloat)// [1 2 3]

	//二维数组
	matrix := [2][2]int64 {
    {2,2},
    {2,2},
  }
	fmt.Println(matrix) //[[2 2] [2 2]]
}

func testSlice(){
	// int型的slice
	ints := make([]int,3)
	ints[0],ints[1],ints[2] = 0,1,2 //没赋值默认为0值
	fmt.Println(ints) //[0,1,2]
	
	// string类型的slice
	names := []string{"zhang","wang","li","zhao"}
	fmt.Println(names, len(names),cap(names)) //["zhang" "wang" "li" "zhao"] 4 4
	
	// 拷贝names 切片
	names1 := names[0:3]
	fmt.Println(names1) //["zhang" "wang" "li"]

        // 修改names1的第一个元素   names与names1共享底层数组
	names1[0] = "laozhang"
	fmt.Println(names) // [laozhang wang li zhao] 

	// 遍历names slice 
	for _, name := range names{
		fmt.Println(name)
    // laozhao
    // wang
    // li
    // zhao
	}

	value1 := make([]int,0)
	for i:=0;i<3;i++{
		value1 = append(value1, i) 
	}
	fmt.Println(value1) // [0 1 2]

	value2 := []int{3,4,5}
	value3 := append(value1, value2...) // ...解包 连接两个slice
	fmt.Println(value3) // [0 1 2 3 4 5]
	
}

func testArrayOperation() {
  names := [4]string{"zhang", "wang", "li", "zhao"}
  fmt.Printf("names has %d elements\n", len(names))

  fmt.Println(names[1])

  names[3] = "lao zhao"
  fmt.Println(names[3])
}

func testSortSlice() {
  people := []struct {
    Name string
    Age int
  }{
    {"zhang", 10},
    {"li", 12},
    {"zhao", 8},
    {"jia", 15},
    {"yi", 1},
    {"bing", 20},
  }
  // Slice no stable sort
  sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name }) 
  fmt.Println("By name:", people)

  sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age})
  fmt.Println("By Age:", people)

  // Slice stable sort
  sort.SliceStable(people, func(i, j int) bool { return people[i].Name > people[j].Name })
  // fmt.Println("By stable name:", people)

  sort.SliceStable(people, func(i, j int) bool { return people[i].Age > people[j].Age })
  // fmt.Println("By stable age:", people)


}
func main(){
	// testArray()
  // testArrayOperation()
	// testSlice()
  // manyInts := make([]int, 100000)
  //
  // a := make([]int, 0) // bad way
  // for _, val := range manyInts {
  //   a = append(a, val+val) // 扩容a 会重新分配内存
  // }
  // // fmt.Println(a)
  //
  // b := make([]int, len(manyInts)) // good way
  // for i, val := range manyInts{
  //   b[i] = val+val // 这里是赋值，不是append
  // }
  // fmt.Println(b)
  testSortSlice()


}
