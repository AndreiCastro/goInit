package main

import fmtAlias "fmt"

func somar(n1 int8, n2 int8) int8 {
	//dentro da parentes são os params - fora é o tipo de retorno
	return n1 + n2
}

func subtrair(n1, n2 int8) int8 {
	//se não colocar o tipo no params, assume o próximo tipo declarado, nesse caso int8
	return n1 - n2
}

func calculoMath(n1, n2 int8) (int8, int8) {
	//Com se tem mais de um tipo de retorno (muito usado com erro de retorno do metodo)
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 45)
	fmtAlias.Println("Soma:", soma)

	sub := subtrair(10, 45)
	fmtAlias.Println("Subtração:", sub)

	//declara 2 variaveis, pois o metodo retorno duas variaveis.
	cSoma, cSub := calculoMath(10, 56)
	fmtAlias.Println("O calculo da soma é:", cSoma, "e Subtração:", cSub)

	//Senão quer usar um retorno, add um _
	//SEMPRE QUE DECLARA UMA VARIAVEL O GO EXIGE QUE A USE, COM O _ DA PRA CONTORNAR ESSA OBRIGATORIEDADE
	cSoma2, _ := calculoMath(10, 56)
	fmtAlias.Println("O calculo da soma é:", cSoma2)
}
