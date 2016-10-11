package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("--> %s \n", runtime.Version())
	fmt.Printf("Hellooo Word!!! \n")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	/*
		var a int
		var b int32
		a = 15
		b = a + a // compiler error
		b = b + 5 // ok: 5 is a constant
	*/
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v", c1)
}

/*
package main
func main() {
println("Hello", "world")
}
*/
