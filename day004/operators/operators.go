package main

import "fmt"

func main() {
	// Arithmetic operator
	a := 10
	b := 5

	fmt.Printf("Addition + : %v, a=%v, b=%v\n", a+b, a, b)
	fmt.Printf("Substraction - : %v, a=%v, b=%v\n", a-b, a, b)
	fmt.Printf("Multiplication * : %v, a=%v, b=%v\n", a*b, a, b)
	fmt.Printf("Division / : %v, a=%v, b=%v\n", a/b, a, b)
	fmt.Printf("Modulas %%: %v, a=%v, b=%v \n", a%b, a, b)
	a++
	fmt.Printf("Increment with ++ : %v \n", (a))
	b++
	fmt.Printf("Decrement with -- : %v \n", (b))
	// fmt.Println(a++)

	// Assignment operator
	var x = 10
	// x+=5
	// x-=3
	// x*=2
	// x/=3
	// x%=3
	x |= 3 // bitwise and & opertor converts both 10 , 3 to binary representation, then do AND
	fmt.Println(x)

	// Comparison operator
	var a1 = 5
	var b1 = 2
	fmt.Println(a1 == b1)
	fmt.Println(a1 != b1)
	fmt.Println(a1 > b1)
	fmt.Println(a1 < b1)
	fmt.Println(a1 >= b1)
	fmt.Println(a1 <= b1)

	// Logical operators
	// x<5 && y<8 // return true if both statement true else false
	// x<5 || y<8 // return true if any of the statement true
	// !( x<5 || y<8) // reverse the result

}
