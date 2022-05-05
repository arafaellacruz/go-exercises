/*
- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

func main() {
	x := 200
	y := x == 100
	z := x != 100
	a := x <= 100
	b := x < 100
	c := x >= 100
	d := x > 100

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", y, z, a, b, c, d)

}