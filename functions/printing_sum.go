package main
import "fmt"

func printSum(a int, b int){
	fmt.Println(a, "+", b, "=", (a+b))
}

func main(){
	printSum(10, 20)
	printSum(111, 222)
}

/** Output

$ go run printing_sum.go
10 + 20 = 30
111 + 222 = 333

*/