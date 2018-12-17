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
	fmt.Println("-----Reverse of a number-----")
	fmt.Println(1234, "-->", getReverse(1234))
	fmt.Println(1001, "-->", getReverse(1001))
}

/** Output

$ go run reverse_of_a_number.go
-----Reverse of a number-----
1234 --> 4321
1001 --> 1001

*/