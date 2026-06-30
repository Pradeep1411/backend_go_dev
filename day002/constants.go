package main
import ("fmt")

func main(){
	// Syntax const CONSTNAME type = value
	const PI =3.14
	fmt.Println(PI)

	//2 types of const : typed, untyped
	const typed int=5
	const untyped string="UnTyped"

	fmt.Println(typed, untyped)
	//multiple constant declaration
	const(
		a int=1
		b="String"
		c=true
	)
	fmt.Printf(" the value of a is %v, and type is %T",a,a)
	fmt.Printf(" the value of b is %v, and type is %T",b,b)
	fmt.Printf(" the value of c is %v, and type is %T",c,c)

}