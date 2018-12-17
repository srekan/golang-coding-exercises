package main
import "fmt"
func getProductOfDigits(num int)int{
	p := 1
	for i := num; i > 0; i = i / 10{
		p *= i % 10
	}

	return p
}
func main(){
	fmt.Println("Product of digits of 1234 =", getProductOfDigits(1234))
	fmt.Println("Product of digits of 1001 =", getProductOfDigits(1001))
}

/** Output

$ go run product_of_digits.go
Product of digits of 1234 = 24
Product of digits of 1001 = 0

*/