// Power of a number x^y
package main
import "fmt"
func getPower(x ,y int)int{
	p := 1
	for i := 1; i <= y; i++{
		p *= x
	}

	return p
}

func main(){
	fmt.Println("2^3  = ", getPower(2, 3))
	fmt.Println("5^5  = ", getPower(5, 5))
	fmt.Println("10^3 = ", getPower(10, 3))
}

/** Output

$ go run get_power.go

2^3  =  8
5^5  =  3125
10^3 =  1000

*/