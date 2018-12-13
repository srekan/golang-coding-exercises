package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)

	if(n == 0){
		fmt.Println("Given number is zero")
	} else if(n > 0){
		fmt.Println("Given number is positive")
	} else{ /* if(n < 0) // Default case*/
		fmt.Println("Given number is negative")
	}
}
/** Outputs

$ go run find_number_sign.go
Enter an integer: 10
Given number is positive

$ go run find_number_sign.go
Enter an integer: -10
Given number is negative

$ go run find_number_sign.go
Enter an integer: 0
Given number is zero

*/