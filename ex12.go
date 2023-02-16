package main;
import ("fmt");

func main() {
	var nomes = [...]string{"Victor", "Gomes", "Lucas", "Gabriel", "Paula"};

	// Loop for
	for i := 0; i < len(nomes); i++ {
		// All names from array
		fmt.Println(nomes[i])
	}

	for i := 0; i < 100 + 1; i++ {
		// Count 
		fmt.Println(i)
		if i == 5 {
			break
		} else {
			continue
		}
	}

	// Others
	var idades = [5]int{1,2,3,4,5}
	for idx, val := range idades {
		fmt.Printf("Index %v \t value %v\n", idx, val)
	}

	for _, val := range idades {
		fmt.Println(val)
	}

}