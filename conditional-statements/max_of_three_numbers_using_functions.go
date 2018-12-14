package main
import "fmt"
func getMax(a, b int)(int){
	if(a > b){
		return a
	} else {
		return b
	}
}
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

	max := getMax(getMax(a, b), c)
	fmt.Println("Max Number =", max)
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