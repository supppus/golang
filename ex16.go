package main

import (
	"fmt"
	"reflect"
)
func operationIsValid(operation string) bool {
	if (operation != "1" && operation != "2" && operation != "3" && operation != "4") {
		return false
	}
	return true
}
func getOperator(operator string) string {
	switch(operator){
		case "1":
			return "+"
		case "2":
			return "-"
		case "3":
			return "*"
		case "4":
			return "/"
		default:
			return "ERROR"
	}
}
func getResultByVectors(operation string, x int, y int) int {
	switch (operation) {
	case "1":
		return x + y
	case "2":
		return x - y
	case "3":
		return x * y
	case "4":
		return x / y
	default:
		return 0
	}
}
func getABX(operacao string){
	a, b := 0, 0
	fmt.Print("\nDigite o valor de A -> ")
	fmt.Scan(&a)
	fmt.Print("\nDigite agora o valor de B -> ")
	fmt.Scan(&b)
	if (a == 0 || b == 0 || reflect.TypeOf(a).String() != "int" || reflect.TypeOf(b).String() != "int") {
		fmt.Println("Confira os dados fornecidos!")
		return;
	}
	fmt.Printf("\n %v %v %v = %v\n Deseja fazer outra operacao?\n 1 - Sim\n 2 - Nao\n--> ", a, getOperator(operacao), b, getResultByVectors(operacao, a, b))
	choiced := ""
	fmt.Scan(&choiced)
	if choiced != "1" {
		fmt.Println("Volte mais tarde :)")
		return
	} else {
		main()
	}
}
func main() {
	fmt.Println(" Calculadora by Victor \n")
	operacao := "";
	fmt.Print("Qual vai ser a operacao?\n 1 - Somar\n 2 - Subtrair\n 3 - Multiplicar\n 4 - Dividir\n --> ")
	fmt.Scan(&operacao)
	if ! operationIsValid(operacao) {
		fmt.Println("Confira os dados fornecidos!")
		return
	}
	getABX(operacao)
}