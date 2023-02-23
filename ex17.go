package main

import (
	"fmt"
	"reflect"
)

var saldo float64 = 50

func isValidChoice(choice string) bool {
	switch(choice){
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
func addSaldo(saldo float64) float64 {
	fmt.Print("\nAgora, digite o valor que queira adicionar ->")
	dinheiro := 0.0;
	fmt.Scan(&dinheiro);
	if reflect.TypeOf(dinheiro).String() != "float64" {
		fmt.Print("Confira o valor!")
		main()
		return saldo
	}
	saldo += dinheiro
	fmt.Printf("\n%v R$ Adicionados ao seu saldo, o total ficou %v R$\n", dinheiro, saldo)
	return saldo
}
func remoSaldo(saldo float64) float64 {
	fmt.Print("\nAgora, digite o valor que queira sacar ->")
	dinheiro := 0.0;
	fmt.Scan(&dinheiro);
	if reflect.TypeOf(dinheiro).String() != "float64" {
		fmt.Print("Confira o valor!")
		main()
		return saldo
	}
	saldo -= dinheiro
	fmt.Printf("\n%v R$ Sacados do seu saldo, o total do seu saldo ficou %v R$\n", dinheiro, saldo)
	return saldo
}
func main(){
	fmt.Print("\n1 - Depositar saldo\n2 - Sacar \n3 - Ver saldo\n--> ")
	choice := "";
	fmt.Scan(&choice)
	if ! isValidChoice(choice) {
		fmt.Println("Confira a opcao...")
		main()
	}
	switch(choice) {
		case "1":
			var action float64 = addSaldo(saldo)
			saldo = action
			main()
		case "2":
			var action float64 = remoSaldo(saldo)
			saldo = action
			main()
		case "3":
			fmt.Printf("\nSeu saldo eh de %v R$\n", saldo)
			main()
	}
}// Sistema simples de banco para inserir moedas...