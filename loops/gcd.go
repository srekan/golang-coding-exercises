package main
import "fmt"
func getGCD(n1, n2 int)int{
	for (n1 != n2){
        if(n1 > n2){
			n1 -= n2
		} else {
			n2 -= n1
		}
	}
	return n1
}
func main(){
	fmt.Println("-----GCD-----")
	fmt.Println("GCD of 12 and 54:", getGCD(12, 54))
	fmt.Println("GCD of 121 and 343:", getGCD(121, 143))
}

/** Output

$ go run gcd.go
-----GCD-----
GCD of 12 and 54: 6
GCD of 121 and 343: 11

*/