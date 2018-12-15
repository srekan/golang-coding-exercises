package main
import "fmt"
func main(){
	var a int
	var b int
	var c int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Print("Enter third number: ")
	fmt.Scan(&c)

	if(a > b){
		if(a > c){
			fmt.Println("Max Number = ", a)
		} else {
			fmt.Println("Max Number = ", c)
		}
	} else {
		if(b > c){
			fmt.Println("Max Number = ", b)
		} else {
			fmt.Println("Max Number = ", c)
		}
	}
}

/** Outputs

$ go run max_of_three_numbers.go
Enter first number: 10
Enter second number: 20
Enter third number: 30
Max Number =  30

$ go run max_of_three_numbers.go
Enter first number: 10
Enter second number: 30
Enter third number: 20
Max Number =  30

$ go run max_of_three_numbers.go
Enter first number: 30
Enter second number: 10
Enter third number: 20
Max Number =  30

*/