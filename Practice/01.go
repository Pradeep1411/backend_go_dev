package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

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
func print12(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			fmt.Print(string('A' + n - j))
		}
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
func print13(n int) {
	x := 1
	for i := 0; i <= 2*(n-1); i++ {
		if i < n {
			for j := 0; j < n-i; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*i; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < n-i; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 0; j <= x; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*(n-x-1); j++ {
				fmt.Print(" ")
			}
			for j := 0; j <= x; j++ {
				fmt.Print("*")
			}
			x++
		}
		fmt.Println()
	}
}
func print14(n int) {
	x := 1
	for i := 0; i <= 2*n-1; i++ {
		if i < n {
			for j := 0; j < n-i; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*i; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < n-i; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 0; j < x; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*(n-x); j++ {
				fmt.Print(" ")
			}
			for j := 0; j < x; j++ {
				fmt.Print("*")
			}
			x++
		}
		fmt.Println()
	}
}
func print15(n int) {
	x := n
	for i := 0; i <= 2*(n-1); i++ {
		if i < n {
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*(n-i-1); j++ {
				fmt.Print(" ")
			}
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j < x; j++ {
				fmt.Print("*")
			}
			for j := 0; j < 2*(n-x+1); j++ {
				fmt.Print(" ")
			}
			for j := 1; j < x; j++ {
				fmt.Print("*")
			}
			x--
		}
		fmt.Println()
	}
}
func print15_1(n int) {
	spaces := 2 * (n - 1)
	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > n {
			stars = 2*n - i
		}
		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}
		if i < n {
			spaces = spaces - 2
		} else {
			spaces = spaces + 2
		}
		fmt.Println()
	}

}
func print16(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func print17(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			top := i
			left := j
			right := 2*(n-1) - j
			bottom := 2*(n-1) - i
			ans := n - min(top, left, right, bottom)
			fmt.Print(ans)
		}
		fmt.Println()
	}
}
func printDigits(n int) {
	// digits := 0
	// for n > 0 {
	// 	digits++
	// 	n = n / 10
	// }
	// fmt.Println("Digits: ", digits)
	digit := math.Floor(math.Log10(float64(n))) + 1
	fmt.Println("New digit:", digit)
}
func reverseNum(n int) {
	rev := 0
	for n > 0 {
		x := n % 10
		rev = rev*10 + x
		n = n / 10
	}
	fmt.Println(rev)
}
func palnidrome(n int) bool {
	sn := strconv.Itoa(n)
	left := 0
	right := len(sn) - 1
	for left < right {
		if sn[left] != sn[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func isArmstrong(n int) {
	n1 := n
	ans := 0
	for n > 0 {
		t := n % 10
		ans = ans + (t * t * t)
		n = n / 10
	}
	if ans == n1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
func printDivisors(n int) {
	divisors := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != (n / i) {
				divisors = append(divisors, n/i)
			}
		}
	}
	fmt.Println(divisors)
	sort.Ints(divisors)
	fmt.Println(divisors)
}
func primeNumber(n int) {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++
			if (n / i) != i {
				count++
			}
		}
	}
	if count == 2 {
		fmt.Printf("%d is a prime number", n)
	} else {
		fmt.Printf("%d is not a prime number", n)
	}
}
func gcdOrhcf(n1, n2 int) {
	gcd := 1
	for i := 2; i <= int(math.Min(float64(n1), float64(n2))); i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}
	fmt.Println(gcd)
}
func eculideanAlgo(n1, n2 int) {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	fmt.Println("GCD is:", n1)
}
func mains() {
	// printStars1(5)
	// reverseNum(10400)
	// fmt.Println(palnidrome(123421))
	// primeNumber(17)
	eculideanAlgo(12, 21)

}
