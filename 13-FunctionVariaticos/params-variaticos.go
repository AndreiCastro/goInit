package main

import "fmt"

func main() {
	escrevaSoma("HelloWorld!", 1, 4, 8, 10)
	escrevaSoma("Sem numeros", 0)

}

func escrevaSoma(texto string, numeros ...int) {
	//params variatico vc pode declarar quandos valores quiser, ou declarer nenhum
	//só devem ser todos do mesmo tipo e ser o ultimo parametro

	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
