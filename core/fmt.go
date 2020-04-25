package core

import "fmt"

func init() {
	AddFunction("CROISSANT", Println)
	AddFunction("PAINAUCHOCOLAT", Printf)
	AddGetter("ÉCLAIR", Scanf)
}

func Println(messages ...interface{}) {
	fmt.Println(messages...)
}

func Printf(params ...interface{}) {
	fmt.Printf(fmt.Sprintf("%v", params[0]), params[1:]...)
	fmt.Println("")
}

func Scanf(params ...interface{}) string {
	a := "";
	fmt.Scanf("%s", &a)
	return a
}
