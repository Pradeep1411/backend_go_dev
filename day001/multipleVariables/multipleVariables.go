package main

import ("fmt")

func main(){
	//  If you use the type keyword, it is only possible to declare one type of variable per line.

	var a,b,c,d int = 1,3,5,7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)


	// If the type keyword is not specified, you can declare different types of variables on the same line:
	var a,b = 6 ,"Hello"
	c,d :=7,"World!"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Go Variable Declaration in a Block

	var (
		a int
		b int=1
		c string="hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/*
	Go variable naming rules:

		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords
	*/



}