/*
- Crie uma variável de tipo string utilizando uma raw string literal (onde colocamos uma string entre `` para informar que queremos que ela seja printada exatamente da forma como escrevemos.)
- Demonstre-a.
*/

package main

import (
	"fmt"
)

func main() {

	x:= `Eu 
			sou
				o
					Batman!`
					fmt.Println(x)
	
}