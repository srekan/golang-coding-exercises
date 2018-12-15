package main
import "fmt"
func getTriangleType(side1, side2, side3 float64)string{
	if(side1 == side2 && side2 == side3){
		return "EQUILATERAL triangle"
	}

	if(side1 == side2 || side2 == side3 || side3 == side1){
		return "ISOSCELES triangle"
	}

	return "SCALENE triangle" // this is the remaining case after the aboce two cases
}
func main(){
	var side1, side2, side3 float64
	fmt.Print("Enter first side: ")
	fmt.Scan(&side1)
	fmt.Print("Enter second second: ")
	fmt.Scan(&side2)
	fmt.Print("Enter third third: ")
	fmt.Scan(&side3)

	fmt.Println(getTriangleType(side1, side2, side3))
}

/** Output

$ go run get_triangle_type.go
Enter first side: 10
Enter second second: 10
Enter third third: 10
EQUILATERAL triangle

$ go run get_triangle_type.go
Enter first side: 10
Enter second second: 10
Enter third third: 20
ISOSCELES triangle

$ go run get_triangle_type.go
Enter first side: 10
Enter second second: 20
Enter third third: 30
SCALENE triangle

*/