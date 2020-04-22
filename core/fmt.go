package core

import "fmt"

func init() {
	AddFunction("CROISSANT", Println)
	AddFunction("PAINAUCHOCOLAT", Printf)
}

func Println(messages ...interface{}) {
	fmt.Println(messages...)
}

func Printf(params ...interface{}) {
	fmt.Printf(fmt.Sprintf("%v", params[0]), params[1:]...)
}
