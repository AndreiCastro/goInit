package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever()
	checkmail.ValidateFormat("teste@teste.com")
}
