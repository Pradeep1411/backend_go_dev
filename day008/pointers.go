package main

import "fmt"

func tenxVal(val int) {
	val *= 10
}
func tenxPointer(pointer *int) {
	*pointer *= 10
}
func mains() {
	x := 5

	fmt.Println("Inital : ", x)
	tenxVal(x)
	fmt.Println("tenxVal : ", x)
	tenxPointer(&x)
	fmt.Println("tenxPointer : ", x)
	p := new(42)
	fmt.Println("value at *p:", *p)
	tenxPointer(p)
	fmt.Println("value at *p:", *p)

}
