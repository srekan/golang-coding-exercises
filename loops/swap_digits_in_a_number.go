/**
Swap the first and last digits of a number
Algorithm:
	start
	step 1: read num
	step 2: numDigits = log10(num)
	step 3: lastDigit = num % 10
	step 4: firstDigit = num / 10^(numDigits - 1)
	
	Find number formed by the digits excluding first and last digits of the number 
	Exclude first digit
		step 5: middleNmber = num % 10^(numDigits - 1)
	Exclude last digit
		step 6: middleNumber = middleNumber % 10
	
	Add in the first digit at the end
		step 7: resultedNumber = middleNumber * 10 + firstDigit
	Add in the last digit at the beginning
		step 8: resultedNumber = lastDigit * 10^(numDigits - 1) + rsultedNumber

	resultedNumber is the result.
	end
Eg:
	start
	step 1: let the number be 1234
	step 2: numDigits = log10(1234) = 4
	step 3: lastDigit = 1234 % 10 = 4
	step 4: firstDigit = 1234 / 10^(4-1) = 1234 / 1000 = 1
	
	step 5: middleNumber = 1234 % 10^(4-1) = 1234 % 1000 = 234
	step 6: middleNumber = 234 % 10 = 23

	step 7: resultedNumber = 23 * 10 + 1 = 231 
	step 8: resultedNumber = 4 * 10^(4-1) + 231 = 4000 + 231 = 4231
	end
*/
package main
import (
	"fmt"
	"math"
)
func getSwappedNum(num int)(int){
	numDigits := int(math.Log10(float64(num))) + 1
	if(numDigits == 1){
		return num
	}

	lastDigit := num % 10
	firstDigit := num / int(math.Pow10(numDigits - 1))

	middleNumber := num % int(math.Pow10(numDigits - 1))
	middleNumber = middleNumber / 10

	resultedNumber := middleNumber * 10 + firstDigit
	resultedNumber = lastDigit * int(math.Pow10(numDigits - 1)) + resultedNumber

	return resultedNumber

}

func main(){
	fmt.Println("1234  -->", getSwappedNum(1234))
	fmt.Println("12345 -->", getSwappedNum(12345))
	fmt.Println("12    -->", getSwappedNum(12))
	fmt.Println("1     -->", getSwappedNum(1))
}

/** Output

$ go run swap_digits_in_a_number.go
1234  --> 4231
12345 --> 52341
12    --> 21
1     --> 1

*/