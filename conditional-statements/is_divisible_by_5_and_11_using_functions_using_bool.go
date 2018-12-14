package main
import "fmt"
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if(n % 5 == 0 && n % 11 == 0){
		fmt.Println("The given number is divisible by both 5 and 11")
	}

	if(n % 5 == 0 && n % 11 != 0){
		fmt.Println("The given number is divisible by 5, but not 11")
	}
	if(n % 5 != 0 && n % 11 == 0){
		fmt.Println("The given number is divisible by 11, but not 5")
	}
	
	if(n % 5 != 0 && n % 11 != 0){
		fmt.Println("The given number is not divisible by 5 and 11")
	}
}

/** Outputs

$ go run is_divisible_by_5_and_11.go
Enter a number: 55
The given number is divisible by both 5 and 11

$ go run is_divisible_by_5_and_11.go
Enter a number: 22
The given number is divisible by 11, but not 5

$ go run is_divisible_by_5_and_11.go
Enter a number: 25
The given number is divisible by 5, but not 11

$ go run is_divisible_by_5_and_11.go
Enter a number: 13
The given number is not divisible by 5 and 11

*/
