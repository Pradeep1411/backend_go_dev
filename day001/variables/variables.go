package main
import "fmt"

func main(){
	// variable  syntax : var variablename type = value

	var name string="Pradeep"
	var age int=27

	fmt.Println(name)
	fmt.Println(age)

	var isMale bool=true
	var salary float32=555.555

	fmt.Println(isMale)
	fmt.Println(salary)

	var collegeName = "Graphic Era Hill Unviersity"
	collegeLocation:="Bhimtal"
	fmt.Println(collegeName)
	fmt.Println(collegeLocation)

	var a string
	var b int
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	/*
		var	
			Can be used inside and outside of functions	
			Variable declaration and value assignment can be done separately
		:=
			Can only be used inside functions
			Variable declaration and value assignment cannot be done separately (must be done in the same line)
	*/

	 
}