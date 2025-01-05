package main

import (
	"fmt"
)

type ErrorValue struct {
	message string
}

func (e ErrorValue) Error() string {
	return e.message
}

type ErrorPointer struct {
	message string
}

func (e *ErrorPointer) Error() string {
	return e.message
}

func main() {
	val := ErrorValue{message: "this is a value receiver"}
	err := useError(val)
	if err != nil {
		fmt.Println("値レシーバー:", err)
	} else {
		fmt.Println("値レシーバー: nil")
	}

	ptr := &ErrorPointer{message: "this is a pointer receiver"}
	err = useError(ptr)
	if err != nil {
		fmt.Println("ポインターレシーバー:", err)
	} else {
		fmt.Println("ポインターレシーバー: nil")
	}
}

func useError(err error) error {
	if err == nil {
		return nil
	}
	return err
}
