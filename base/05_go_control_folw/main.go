package main

import (
	"fmt"
	"time"
)

func testIf() {
	ok := true
	if ok {
		fmt.Println("ok is true")
	}

	day := "Firday"
	if day == "Firday" {
		fmt.Println("周五啦")
	} else if day == "Monday" {
		fmt.Println("该上班了")
	} else {
		fmt.Println("摸鱼期")
	}

	m := make(map[string]string)
	m["wang"] = "wudi"

	if v, ok := m["wang"]; ok {
		fmt.Printf("%s is exist", v)
	}

}

func testSwitch() {
	day := 1
	switch day {
	case 0, 6:
		fmt.Println("周末啦")
	case 1, 2, 3, 4, 5:
		fmt.Println("摸鱼")
	default:
		fmt.Println("不合法")
	}

	a, b := 1, 2
	switch {
	case a > b:
		fmt.Println("a>b")
	case a < b:
		fmt.Println("a<b")
	}

}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("sleep")
	}

}

func main() {
	// testIf()
	// testSwitch()
	testFor()
}
