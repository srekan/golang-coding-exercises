package main
import "fmt"
func getReverse(num int)int{
	rev := 0
	for i := num; i > 0; i = i/10{
		rev = rev * 10 + (i % 10)
	}

	return rev
}

func main(){
	fmt.Println("-----Is Pallindrome-----")
	fmt.Println(1234, "-->", getReverse(1234) == 1234)
	fmt.Println(1001, "-->", getReverse(1001) == 1001)
}

/** Output

$ go run is_pallindrome.go
-----Is Pallindrome-----
1234 --> false
1001 --> true

*/