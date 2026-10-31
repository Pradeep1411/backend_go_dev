package main

import "fmt"

func fun(n int) {
	fmt.Println(n)
	fun(n)
}
func main() {
	fun(12)
}
