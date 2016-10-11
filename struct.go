package main 

import (
	"fmt"	
)

type bts struct{
		age int
		name, last_name string 
	}	

func  main() {
	//var kat bts
	beast := bts{7, "a", "b"} //asignar valor a todos los elementos
	inf := bts{age:22,last_name:"liz", name:"katy" } //no es necesario asignar valor a todos los elementos
	kat := bts{name:"katy", last_name:"liz"}
	//fmt.Println(bts {name:"katy", last_name:"lizbeth"})
	fmt.Println(kat)
	fmt.Println(inf)
	fmt.Println(beast)

	yei := new(bts) // devuelve puntero
	fmt.Println(yei)

	fmt.Println(kat.name)//  (*kat).name

}