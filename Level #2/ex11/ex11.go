/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

	const (
		_ = 2022 + iota
		x
		y 
		z 
		a
)

func main() {

	fmt.Println(x, y, z, a)

}