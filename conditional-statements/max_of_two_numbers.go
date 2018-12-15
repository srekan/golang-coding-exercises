package main
import "fmt"
func main(){
	var a int
	var b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	if(a > b){
		fmt.Println("Max number is: ", a)
	}

	if(b > a){
		fmt.Println("Max number is: ", b)
	}

	if(a == b){
		fmt.Println("Both numbers are equal")
	}
}

/** Outputs

$ go run max_of_two_numbers.go
Enter first number: 20
Enter second number: 10
Max number is:  20

$ go run max_of_two_numbers.go
Enter first number: 10
Enter second number: 20
Max number is:  20

$ go run max_of_two_numbers.go
Enter first number: 20
Enter second number: 20
Both numbers are equal

*/