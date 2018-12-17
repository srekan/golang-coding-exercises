package main
import "fmt"
func getFactorial(num int)int64{
	var fact int64 = 1
	for i := num; i > 0; i--{
		fact *= int64(i)
	}

	return fact
}
func main(){
	fmt.Println("-----Finding Factorials-----")
	fmt.Println("5! = ", getFactorial(5))
	fmt.Println("10! = ", getFactorial(10))
}

/** Output

$ go run factorial.go
-----Finding Factorials-----
5! =  120
10! =  3628800

*/