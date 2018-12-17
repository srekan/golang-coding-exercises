package main
import "fmt"
func getSumOfDigits(num int)int{
	sum := 0
	for i := num; i > 0; i = i / 10{
		sum += i % 10
	}

	return sum
}
func main(){
	fmt.Println("Sum of digits of 1234 =", getSumOfDigits(1234))
	fmt.Println("Sum of digits of 1001 =", getSumOfDigits(1001))
}

/** Output

$ go run sum_of_digits.go
Sum of digits of 1234 = 10
Sum of digits of 1001 = 2

*/