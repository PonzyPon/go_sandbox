package main

import (
	"errors"
)

func main() {
	err1 := errors.New("foo")
	err2 := errors.New("foo")
	println(err1 == err1) // true
	println(err1 == err2) // false

	customErr1 := customError{}
	customErr2 := customError{}
	println(customErr1 == customErr1) // true
	println(customErr1 == customErr2) // true
	println(&customErr1)              // 0xc000080f0f
	println(&customErr2)              // 0xc000080f0f (1と同じポインタ)

	customValueError1 := customValueError{}
	customValueError2 := customValueError{}
	println(customValueError1 == customValueError1) // true
	println(customValueError1 == customValueError2) // true
	println(&customValueError1)                     // 0xc00007ef00
	println(&customValueError2)                     // 0xc00007eef0 (1と違うポインタ)
}

// structにfieldを持たない
type customError struct{}

func (c customError) Error() string {
	return "CustomError"
}

// structにfieldを持つ
type customValueError struct {
	value string
}

func (c customValueError) Error() string {
	return "CustomValueError"
}
