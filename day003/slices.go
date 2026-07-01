package main
import "fmt"

func main(){
	//Syntax Using the []datatype{values} format
	var slice1=[]int{}
	slice2:=[]string{"go","slices","are","powerful"}
	fmt.Println(slice1)
	fmt.Println(slice2)

	// len(slice-name) :  returns the length of the slice (the number of elements in the slice)
	// cap(slice-name) : returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	// cap is starts from the slice starting index till end of the array

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	//  Syntax using an array 
	// var array=[length]datatype{values}
	// slice:=array[start:end]

	arr1:=[...]int{10,11,12,13,14,15}
	myslice:=arr1[2:4]

	fmt.Printf("my slice = %v \n",myslice)
	fmt.Printf("length = %d \n",len(myslice))
	fmt.Printf("capacity = %d \n",cap(myslice))

	// with make fun, slice_name := make([]type, length, capacity) 
	// Note: If the capacity parameter is not defined, it will be equal to length.



	fun:=make([]string,3,5)
	cars:=make([]int,3)
	fmt.Printf("Fun length = %d \n",len(fun))
	fmt.Printf("Fun capacity = %d \n",cap(fun))
	fmt.Printf("Cars length = %d \n",len(cars))
	fmt.Printf("Cars capacity = %d \n",cap(cars))


}
