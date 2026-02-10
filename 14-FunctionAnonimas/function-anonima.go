package main

import "fmt"

func main() {
	//declara uma função sem nome
	func(text string) {
		fmt.Println(text)
	}("Aqui fica o params") //Sempre deve ter () com ou sem params.
}
