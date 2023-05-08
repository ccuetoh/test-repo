package main

import (
	"errors"
	"fmt"
)

func main() {
	errorFunc("This should error and we don't handle it..")
	fmt.Println("Hello world!")
}

func unusedFunc() {
	fmt.Println("unused")
}

func errorFunc(msg string) error {
	fmt.Println(msg)
	return errors.New("this is an error")
}
