/**
Strong Number: the sum of factorial of the digits in any number is equal the given number

eg:
1 :>
	1! = 1
2 :> 
	2! = 2
145 :>
	1! + 4! + 5! => 1 + 24 + 120 => 145

*/
package main
import "fmt"
func getFactorial(num int)int64{
	var fact int64 = 1
	for i := num; i > 0; i--{
		fact *= int64(i)
	}

	return fact
}
func isStrong(num int)bool{
	sum := 0
	for i := num; i > 0; i /= 10{
		sum += int(getFactorial(i % 10))
	}

	return sum == num
}
func main(){
	fmt.Println("-----Is Strong Number-----")
	fmt.Println(1, "----->", isStrong(1))
	fmt.Println(2, "----->", isStrong(2))
	fmt.Println(145, "--->", isStrong(145))
	fmt.Println(20, "---->", isStrong(20))
}

/** Output

$ go run is_strong.go
-----Is Strong Number-----
1 -----> true
2 -----> true
145 ---> true
20 ----> false

*/