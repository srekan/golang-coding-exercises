package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if(n % 2 == 0){
		fmt.Println("Given number is even")
	} else {
		fmt.Println("Given number is odd")
	}
}

/** Outputs

$ go run is_even_or_odd.go
Enter a number: 2
Given number is even

$ go run is_even_or_odd.go
Enter a number: 3
Given number is odd

$ go run is_even_or_odd.go
Enter a number: 0
Given number is even

*/
