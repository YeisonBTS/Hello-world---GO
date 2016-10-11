package main

import (
	"fmt"
	"bufio"
	"os"
)
	

func main() {
	edad := 22;
	name := "katy"
	love := true;
	decimal := 1.96;

	var old int

	fmt.Println("hi! ",edad+1)
	fmt.Printf("hello! %d \n",edad)
	fmt.Printf("hello! %v \n",name)	//imprime el valor origuianl del dato 
	fmt.Printf("hello! %s \n",name)
	fmt.Printf("hello! %t \n",love)
	fmt.Printf("hello! %f \n",decimal)
	fmt.Printf("hello! %e \n",decimal)
	fmt.Printf("hello! %b \n",decimal)

	fmt.Println("how old are you?..")
	fmt.Scanf("%d\n", &old)
	fmt.Printf("%d years old \n", old)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingres your name:")
	text,err := reader.ReadString('\n')	

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("hi! "+text)
	}
}