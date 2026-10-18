package main

import "fmt"

func printStars(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Print("*", " ")
		}
		fmt.Println("")
	}
}
func printStars1(n int) {
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}
}

func main() {
	printStars1(5)
}
