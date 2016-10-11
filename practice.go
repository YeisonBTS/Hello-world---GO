package main

import (
	"fmt"
	"math"
)

func main() {
	radio := 3.5
	circunferencia := 2 * math.Pi * radio
	fmt.Println("Circunferencia = ", circunferencia)
	circulo := math.Pi * radio * radio
	fmt.Println("Circulo = ", circulo)
}
