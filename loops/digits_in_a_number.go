package main
import "fmt"
func getNumDigits(num int)int{
	digitCount := 0
	for i := num; i > 0; i = i / 10{
		digitCount++
	}

	return digitCount
}

func main(){
	var num int
	fmt.Print("Entr a number: ")
	fmt.Scan(&num)

	fmt.Println("Number of digits in the given number = ", getNumDigits(num))
}

/** Output

$ go run digits_in_a_number.go
Entr a number: 2345
Number of digits in the given number =  4

$ go run digits_in_a_number.go
Entr a number: 12345
Number of digits in the given number =  5

*/