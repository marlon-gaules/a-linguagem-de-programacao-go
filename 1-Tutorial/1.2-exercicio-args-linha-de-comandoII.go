// Modifique o programa Echo para exibir também os.Args[0], que é o nome do comando que o chamou
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		s = " "
	}
	fmt.Println(s)
	fmt.Println(sep)
}
