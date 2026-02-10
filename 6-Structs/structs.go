package main

import "fmt"

type user struct {
	//Struct é como se fosse uma class
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logrdouro string
	numero    int8
}

func main() {

	var usuario user
	usuario.nome = "Andrei"
	usuario.idade = 34
	fmt.Println(usuario)
	//Senão popula a prop fica zerado

	enderecoUser := endereco{"Rua 2", 0}
	usuario2 := user{"May", 33, enderecoUser}
	fmt.Println(usuario2)

	//popular so uma prop especifica da struct
	usuario3 := user{idade: 55}
	fmt.Println(usuario3)

}
