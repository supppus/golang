package main;
import ("fmt");

func sum(a int, b int) int {
	return a + b;
}

// func nameFuncion(params datatype) datatype to return {logic}

func main() {
	var x, y int = 50, 125;
	fmt.Printf("50 + 125 = %v\n", sum(x, y))
}