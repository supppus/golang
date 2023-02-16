package main;
import ("fmt")

func main() {
	// Arrays
	var nomes = [2]string{"victor", "victor"};
	var idades = [5]int{11, 12, 13, 14, 15};
	var carros = [3]string{"audi", "bmw", "volks"};
	fmt.Println(nomes)
	fmt.Println(idades)
	fmt.Println(carros)

	// By index
	fmt.Println(nomes[0])
	fmt.Println(idades[2])
	fmt.Println(carros[2])

	// Change value by index
	nomes[0] = "victor 2";
	idades[2] = 22;
	fmt.Println(nomes[0])
	fmt.Println(idades[2])

	// Arrays where not value
	var array = [2]int{}
	var arrayTwo = [2]int{1}
	fmt.Println(array) // [0 0]
	fmt.Println(arrayTwo) // [1 0]

	// Get length of array where am using a method len()
	var arrayTest = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(len(arrayTest)) // 5

	// Creating a array where not defined lenght
	var arrayTestTwo= [...]string{"victor", "teste"}
	fmt.Println(arrayTestTwo)
}