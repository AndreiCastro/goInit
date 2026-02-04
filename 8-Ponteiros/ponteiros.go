package main

import "fmt"

func main() {
	//PONTEIRO É UMA REFERENCIA DA MEMÓRIA, NÃO O VALOR DELA EX:

	n1 := 10
	n2 := n1
	fmt.Println(n1, n2)
	//N2 recebeu o valor de n1

	n3 := 20
	var ponteiro *int //aqui vc declara variavel como ponteiro

	ponteiro = &n3             //tem que colocar o & pra atribuir o local da memoria que esta a variavel, nesse caso n3
	fmt.Println(n3, ponteiro)  //ponteiro normal exibe o local da memoria
	fmt.Println(n3, *ponteiro) //ponteiro com * pega o valor que tem nesse local da memoria

	n3 = 50
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	//Repare que o local da memoria é o mesmo, mas o valor muda.
	//Ponteiro serve pra buscar pelo local da memoria, sem precisar ocupar novo espaço, ex criando
	//uma nova variavel ocupando um novo espaço em memoria

}
