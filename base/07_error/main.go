package main

import (
	"errors"
	"fmt"
)

func testDefer() string {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("函数体")
	return "返回值"
}

func testDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero") //返回错误
	}
	return a / b, nil

}

type ArticleError struct {
	Code    int32
	Message string
}

func (e *ArticleError) Error() string {
	return fmt.Sprintf("[ArticleError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewArticleError(code int32, message string) error {
	return &ArticleError{
		Code:    code,
		Message: message,
	}
}

func MustDivide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}

	return a / b
}

func Divide2(a, b int) (res int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	res = MustDivide(a, b)
	return
}
func main() {
	fmt.Println(testDefer())
	// a, b := 1, 0
	// res, err := testDivide(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)

	// fmt.Println(MustDivide(1, 0))
	// fmt.Println("print")

	// res, err := Divide2(10, 0)
	// fmt.Println(res, err)
}
