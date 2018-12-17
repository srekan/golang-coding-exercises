package main
import "fmt"
func getDigits(num int)(int, int){
	lastDigit := num % 10
	firstDigit := lastDigit
	for i := num / 10; i > 0; i = i / 10{
		firstDigit = i % 10
	}

	return firstDigit, lastDigit
}

func main(){
	var num int
	fmt.Print("Entr a number: ")
	fmt.Scan(&num)

	f, l := getDigits(num)

	fmt.Println("First Digit =", f)
	fmt.Println("Last Digit  =", l)
}

/** Output

$ go run digits_of_a_number.go
Entr a number: 1234
First Digit = 1
Last Digit  = 4

$ go run digits_of_a_number.go
Entr a number: 8899
First Digit = 8
Last Digit  = 9

*/