/*
Arm strong number: 
	Def1: Sum of cubes of digits of the number = the number
	Def2: A number that is the sum of its own digits each raised to the power of the number of digits.
eg:
153 = 1^3 + 5^3 + 3^3
370 = 3^3 + 7^3 + 0^3
371 = 3^3 + 7^3 + 1^3
407 = 4^3 + 0^3 + 7^3
*/
package main
import (
	"fmt"
	"math"
)
func isArmstrong(num int)bool{
	sum := 0
	for i := num; i > 0; i = i / 10{
		sum += int(math.Pow(float64(i % 10), 3))
	}

	return sum == num
}
func main(){
	fmt.Println("-----Is Armstrong Number-----")
	fmt.Println("111 ---->", isArmstrong(100))
	fmt.Println("153 ---->", isArmstrong(153))
	fmt.Println("370 ---->", isArmstrong(370))
	fmt.Println("371 ---->", isArmstrong(371))
	fmt.Println("407 ---->", isArmstrong(407))
}

/** Output

$ go run is_armstrong.go
-----Is Armstrong Number-----
111 ----> false
153 ----> true
370 ----> true
371 ----> true
407 ----> true

*/