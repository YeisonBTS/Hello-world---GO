package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	sexo := "1010"

	edad_string := strconv.Itoa(edad)
	sexo_int,_ := strconv.Atoi(sexo)
	
	fmt.Println("dasd"+edad_string)
	fmt.Println(sexo_int)
}
