package main

import "fmt"

func main() {
	var teste string = "Teste" //variavel explicita
	var (
		teste2 string = "Teste2"
		teste3 string = "Teste3"
	)

	teste4 := "Teste4" //variavel implicita
	teste5, teste6 := "Teste5", "Teste6"

	fmt.Println(teste)
	fmt.Println(teste2)
	fmt.Println(teste3)
	fmt.Println(teste4)
	fmt.Println(teste5)
	fmt.Println(teste6)
}
