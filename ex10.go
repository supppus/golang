package main;
import ("fmt")

func main() {
	// IF ELSE
	var a, b int = 5, 10;
	if a < b {
		fmt.Println("A < B")
	}else if a == b {
		fmt.Println("A == B")
	} else {
		fmt.Println("A > B");
	}

	// nested if
	var value int = 10;
	if (value > 1) {
		fmt.Println("Test 1 passed");
		if (value > 2) {
			fmt.Println("Test 2 passed");
			if (value > 3) {
				fmt.Println("Test 3 passed");
			} else {
				fmt.Println("Test 3 broken");
			}
		} else {
			fmt.Println("Test 2 broken");
		}
	} else {
		fmt.Println("Test 1 broken");
	}
}