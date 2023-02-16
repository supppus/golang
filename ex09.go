package main;
import ("fmt")

func main() {
	// Operators
	var a, b, c int = 50, 150, 1500;
	a--; // Decrement
	a++; // Increment
	b--; b++; c--; c++;

	fmt.Println(a + 1) // Sum
	fmt.Println(a - 1) // Sub
	fmt.Println(a * 1) // Mult
	fmt.Println(a / 2) // Division
	fmt.Println(a % 1) // remainder division
	
	fmt.Println(a + b + c);
	fmt.Println(a + b / c);

	// Increment with number
	a += 5; // Icrement 5
	b += 10; // Increment 10

	// Decrement with number
	a -= 5; // decrement 5
	b -= 10; // decrement 10

	// Multi with number
	a *= 2; // a = a * 2
	b *= 5; // b = b * 5

	// Division with number
	a /= 2; // a = a / 2
	b /= 5; // b = b / 5

	// Remainder division with number
	a %= 2; // a = a % 2
	b %= 5; // b = b % 5

	// comparison operators
	// == igual
	// !=  diferente
	// > maior que
	// < menor que
	// >= maior ou igual a
	// <= menor ou igual a 

	// Logical operators
	// && e 
	// || ou 
	// ! nao

	fmt.Println(a > b);
	fmt.Println(a < b);
	fmt.Println(a >= b);
	fmt.Println(a <= b);
	fmt.Println(a != b);
	fmt.Println(a > b && a < c);


}