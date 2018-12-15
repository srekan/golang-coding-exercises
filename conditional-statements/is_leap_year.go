/*
Refer: Leap year: https://en.wikipedia.org/wiki/Leap_year

Algorithm:
if (year is not divisible by 4) then (it is a common year)
else if (year is not divisible by 100) then (it is a leap year)
else if (year is not divisible by 400) then (it is a common year)
else (it is a leap year
*/
package main
import "fmt"
func isLeapYear(year int)bool{
	if(year % 4 != 0 ){
		return false
	} else if(year % 100 != 0){
		return true
	} else if(year % 400 != 0){
		return false
	} else {
		return true
	}

	/*
	Another method: compact yet complex
	if(((year % 4 == 0) and (year % 100 !=0)) or (year % 400==0))
	{
		LEAP YEAR
	}
	else
	{
		COMMON YEAR;
	}
	*/
}
func main(){
	var year int
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if(isLeapYear(year)){
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Common Year")
	}

}

/** Outputs

$ go run is_leap_year.go
Enter year: 2018
Common Year

$ go run is_leap_year.go
Enter year: 2020
Leap Year

$ go run is_leap_year.go
1Enter year: 900
Common Year

$ go run is_leap_year.go
Enter year: 2024
Leap Year

*/