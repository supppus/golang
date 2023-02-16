package main;
import ("fmt")

func main() {
	// Somente em uma linha
	fmt.Print(1 + 999)
	// Quebrando linha
	var Test, TestTwo int = 5, 8;
	fmt.Print(Test, "\n");
	fmt.Print(TestTwo, "\n");
	// Mesmo resultado quebrando a linha
	fmt.Print(Test, "\n", TestTwo, "\n");
	var a, b = 1, 2;
	fmt.Print(a,b); // 1 2
}