package main 

import (
	"fmt"	
)

func  main() {

	var x, y *int
	ent := 5;
	/*
	- * acceso valor 
	- & acceso direccion
	*/

	x = &ent
	y = &ent
	*x = 10
	fmt.Println(*x)
	fmt.Println(*y)
}