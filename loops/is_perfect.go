/*
Perfect Number: 
	the number == the sum of its positive divisors excluding the number itself
eg: 
	6 :> 
		devisors sum => 1 + 2 + 3 == 6
	28 :> 
		devisors sum => 1 + 2 + 4 + 7 + 14 == 28
	496 :>
		devisors sum => 1 + 2 + 4 + 8 + 16 + 31 + 62 + 124 + 248 == 496
	8128 :>
		devisors sum => 1 + 2 + 4 + 8 + 16 + 32 + 64 + 127 + 254 + 508 + 1016 + 2032 + 4064
*/
package main
import "fmt"
func isPerfect(num int)bool{
	sum := 0
	for i := 1; i < num; i++{
        if num % i == 0 {
            sum += i;
        }
	}
	
	return sum == num
}
func main(){
	fmt.Println("-----Is Perfect Number-----")
	fmt.Println(6, "------->", isPerfect(6))
	fmt.Println(28, "------>", isPerfect(28))
	fmt.Println(90, "------>", isPerfect(90))
	fmt.Println(496, "----->", isPerfect(496))
	fmt.Println(8128, "---->", isPerfect(8128))
}

/** Output

$ go run is_perfect.go
-----Is Perfect Number-----
6 -------> true
28 ------> true
90 ------> false
496 -----> true
8128 ----> true

*/