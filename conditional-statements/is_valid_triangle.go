package main
import "fmt"
func isValidTriangle(angle1, angle2, angle3 float64)bool{
	return ((angle1 + angle2 + angle3) == 180) 
}
func main(){
	var angle1, angle2, angle3 float64
	fmt.Print("Enter first angle: ")
	fmt.Scan(&angle1)
	fmt.Print("Enter second angle: ")
	fmt.Scan(&angle2)
	fmt.Print("Enter third angle: ")
	fmt.Scan(&angle3)

	if(isValidTriangle(angle1, angle2, angle3)){
		fmt.Println("The give angles CAN form a triangle")
	} else {
		fmt.Println("The give angles CAN_NOT form a triangle")
	}
}

/** Output

$ go run is_valid_triangle.go
Enter first angle: 30
Enter second angle: 60
Enter third angle: 90
The give angles CAN form a triangle

$ go run is_valid_triangle.go
Enter first angle: 90
Enter second angle: 90
Enter third angle: 90
The give angles CAN_NOT form a triangle

*/