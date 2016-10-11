package main 

import (
	"fmt"
)

func main() {

	/*for i:=0; i<10; i++{
		fmt.Println(i)
	}*/

	k := 0
	for{
		if k == 7{
			k++
			continue
		}
		fmt.Println(k)
		k++
		if k > 10{
			break
		}
	}
}