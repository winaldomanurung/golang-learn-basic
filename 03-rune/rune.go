package main

import (
	"fmt"
	"reflect"
)

func main(){
	var x rune = 'x'
	fmt.Println(reflect.TypeOf(x))
	fmt.Println('A')
	fmt.Println('b')
	fmt.Println('\n')
}