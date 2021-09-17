package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil

	return err
	//return nil
}

func main() {
	err := Foo()

	fmt.Println(err)                        //nil
	fmt.Println(err == nil)                 //false
	fmt.Println(err.(*os.PathError) == nil) //true
}
