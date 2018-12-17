package main
import (
	"fmt"
	"math"
)

var names = map[int]string{
	0: "",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
	100: "Hundred",
	1000: "Thousand",
	10000: "Ten Thousand",
	100000: "Million",
	1000000: "Ten Million",
	10000000: "Billion",
	100000000: "Ten Billion",
	1000000000: "Trillion",
}
func isPowerOf10(num int)bool{
	var pow int64 = 1
	for pow < int64(num) {
		pow *= 10
	}
	return pow == int64(num)
} 
func getNameForTwoDigitNumber(num int)string{
	if num < 20 {
		return names[num]
	}
	
	return names[(num / 10) * 10] + " " + names[num % 10]
}
func getNameForThreeDigitNumber(num int)string{
	return names[num/100] + " hundred " + getNameForTwoDigitNumber(num % 100)
}
func getNumberInWords(num int)string{
	// We choose map over array because map is more expressive when we read the code
	

	if num == 0{
		return "zero"
	}

	if isPowerOf10(num) || num < 20{
		return names[num]
	}
	numDigits := int(math.Log10(float64(num))) + 1

	var namedNumber string
	switch(numDigits){
	case 2:
		namedNumber = getNameForTwoDigitNumber(num)
	case 3:
		namedNumber = getNameForThreeDigitNumber(num)
	default:
		powerIn10 := int(math.Pow(10, 3.0))
		namedNumber = getNumberInWords(num / powerIn10) + " " + names[powerIn10] + " " + getNumberInWords(num % powerIn10)		
	}

	return namedNumber
}

func main(){
	fmt.Println("-----Printing numbers in words-----")
	fmt.Println(0, "-------------->", getNumberInWords(0))
	fmt.Println(6, "-------------->", getNumberInWords(6))
	fmt.Println(13, "------------->", getNumberInWords(13))
	fmt.Println(20, "------------->", getNumberInWords(20))
	fmt.Println(54, "------------->", getNumberInWords(54))
	fmt.Println(100, "------------>", getNumberInWords(100))
	fmt.Println(197, "------------>", getNumberInWords(197))
	fmt.Println(999, "------------>", getNumberInWords(999))
	fmt.Println(1000, "----------->", getNumberInWords(1000))
	fmt.Println(1004, "----------->", getNumberInWords(1004))
	fmt.Println(1999, "----------->", getNumberInWords(1999))
	fmt.Println(10000, "---------->", getNumberInWords(10000))
	fmt.Println(99999, "---------->", getNumberInWords(99999))
	fmt.Println(100000, "--------->", getNumberInWords(100000))
	fmt.Println(999999, "--------->", getNumberInWords(999999))
	fmt.Println(10000000, "------->", getNumberInWords(10000000))
}

/** Output 

$ go run print_num_in_words.go
-----Printing numbers in words-----
0 --------------> zero
6 --------------> Six
13 -------------> Thirteen
20 -------------> Twenty
54 -------------> Fifty Four
100 ------------> Hundred
197 ------------> One hundred Ninety Seven
999 ------------> Nine hundred Ninety Nine
1000 -----------> Thousand
1004 -----------> One Thousand Four
1999 -----------> One Thousand Nine hundred Ninety Nine
10000 ----------> Ten Thousand
99999 ----------> Ninety Nine Thousand Nine hundred Ninety Nine
100000 ---------> Million
999999 ---------> Nine hundred Ninety Nine Thousand Nine hundred Ninety Nine
10000000 -------> Billion

*/