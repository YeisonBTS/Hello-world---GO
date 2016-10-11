package main 

import (
	"fmt"	
)

func  main() {

	//slice := make([]string,5)
				//make(type, len, cap)
	slice := make([]int, 5, 10)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice,2)

	fmt.Println(slice)
}