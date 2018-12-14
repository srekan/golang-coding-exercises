package main
import "fmt"

func getSum(a, b int)int{
	return (a + b)
}

func main(){
	fmt.Println("10 + 20 =", getSum(10, 20))
	fmt.Println("11 + 22 =", getSum(11, 22))
}

/** Output

$ go run return_value.go
10 + 20 = 30
11 + 22 = 33

*/