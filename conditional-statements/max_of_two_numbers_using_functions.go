package main
import "fmt"
func printMax(a int, b int){
	if(a > b){
		fmt.Println("Max number is: ", a)
	}

	if(b > a){
		fmt.Println("Max number is: ", b)
	}

	if(a == b){
		fmt.Println("Both numbers are equal")
	}

	fmt.Println(".............")
}

func main(){
	fmt.Print("20, 10 :> ")
	printMax(20, 10)
	fmt.Print("10, 20 :> ")
	printMax(10, 20)
	fmt.Print("20, 20 :> ")
	printMax(20, 20)
}

/** Output

$ go run max_of_two_numbers_using_functions.go
20, 10 :> Max number is:  20
.............
10, 20 :> Max number is:  20
.............
20, 20 :> Both numbers are equal
.............

*/