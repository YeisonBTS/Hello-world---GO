package main 

import (
	"fmt"	
)

func  main() {

	//slice := make([]string,5)
				//make(type, len, cap)
	slice := []int{1,2,3,4,5}
	copia := make([]int,len(slice),len(slice)*2)

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
	
}