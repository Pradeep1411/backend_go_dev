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
func printStars2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}

		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
func printStars3(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(n-i)-1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
func printStars4(n int) {
	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > n {
			stars = 2*n - i
		}
		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
func print5(n int) {
	for i := 1; i <= n; i++ {
		start := i % 2
		for j := 1; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()

	}
}

func print6(n int) {
	for i := 1; i <= n; i++ {
		space := 2 * (n - i)
		for k := 1; k <= i; k++ {
			fmt.Print(k)
		}
		for k := 1; k <= space; k++ {
			fmt.Print(" ")
		}
		for k := i; k >= 1; k-- {
			fmt.Print(k)
		}
		fmt.Println()
	}
}
func print7(n int) {
	c := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(c, " ")
			c++
		}
		fmt.Println()
	}
}
func print8(n int) {
	for i := 0; i < n; i++ {
		for ch := 'A'; ch <= 'A'+rune(i); ch++ {
			fmt.Print(string(ch), " ")
		}
		fmt.Println()
	}
}
func print9(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(string('A'+rune(j)), " ")
		}
		fmt.Println()
	}
}
func print10(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A'+rune(i)), " ")
		}
		fmt.Println()
	}
}
func print11(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i-1; j++ {
			fmt.Print(" ")
		}
		ch := 'A'
		brk := (2*i + 1) / 2
		for j := 1; j <= 2*i+1; j++ {
			fmt.Print(string(ch))
			if j <= brk {
				ch++
			} else {
				ch--
			}
		}
		for j := 0; j <= n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
func main() {
	// printStars1(5)
	print11(5)
}
