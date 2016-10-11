package main 

import (
	"fmt"	
)

func  main() {
	//var matriz [] int{}
	mat := [] int {1, 2, 3, 4}

	fmt.Println(mat)

	arreglo := [3]int{10,11,12}

	slice := arreglo[:] //posicion [desde :hasta]

	fmt.Println(slice)
}