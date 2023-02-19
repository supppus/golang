package main

import (
	"fmt"
)

// tabuada
func startCalc(choiced string, de int, ate int){
	fmt.Printf("\n Tabuada do %v ate o %v \n", de, ate)
	for i := 0; i < ate; i++ {
		fmt.Printf("\n %v %v %v = %v", i, getChoicedByVector(choiced), de, getCalcByChoice(choiced, i, de))
	}
}
func getChoicedByVector(choiced string) string {
	switch (choiced) {
		case "1":
			return "+"
		case "2":
			return "-"
		case "3":
			return "*"
		default:
			return "ERRO"
	}
}
func getCalcByChoice(choiced string, a int, b int) int {
	switch (choiced) {
	case "1":
		return a + b
	case "2":
		return a - b
	case "3":
		return a * b
	default:
		return 0
	}
}
func verifyChoice(choiced string) bool {
	switch (choiced){
		case "1":
			return true
		case "2":
			return true
		case "3":
			return true
		default:
			return false
	}
}
func main(){
	fmt.Printf("\n %T Tabuada 1.0 By Victor \n")
	fmt.Print("1 - Somar\n2 - Subtrair\n3 - Multiplicar\n--> ")
	received := "";
	fmt.Scan(&received)
	if ! verifyChoice(received) {
		fmt.Println("Confira a opcao selecionada!")
		return 
	} 
	fmt.Print("\nDigite agora o numero da tabuada\n-> ")
	de := 0;
	fmt.Scan(&de)
	if (de == 0) {
		fmt.Println("Confira o valor fornecido!")
		return
	}
	fmt.Print("\nDigite agora ate onde vai a tabuada\n-> ")
	ate := 0;
	fmt.Scan(&ate)
	if (ate == 0) {
		fmt.Println("Confira o valor fornecido!")
		return
	}
	startCalc(received, de, ate + 1)
}