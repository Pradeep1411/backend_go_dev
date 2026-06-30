package main
import ("fmt")

func main(){
	//sytax, var arrayName=[length]datatype{values}
	// var arrayName=[...]datatype{values}
	var cars=[3]string{"tata","maruti","VW"}
	fmt.Println(cars)
	//sytax,  arrayName:=[length]datatype{values}
	//  arrayName:=[...]datatype{values}
	bikes:=[...]string{"pulsar","dukati","dhoom"}
	fmt.Println(bikes)

	//Accessing elements
	fmt.Println("First car is : ",cars[0])
	fmt.Println("First bike is : ",bikes[0])

	prices:=[...]int{1,2,3,4,5}
	fmt.Println(prices)
	prices[0]=-2
	fmt.Println(prices)

	//Initialization : default initialization for int is 0, string ""
	arr1:=[5]int{}
	arr2:=[5]int{1,2}
	arr3:=[5]int{1,2,3,4,5}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//specific initialization
	arr4:=[7]int{1,2:-5,3:5,40}
	fmt.Println(arr4)

	// finding length : len(arrayName)
	fmt.Println(len(arr4))
}
