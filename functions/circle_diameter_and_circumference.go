package main
import (
	"fmt"
	"math"
)

func getCircleDetails(radius float64)(float64, float64){
	return (2 * math.Pi * radius), (2 * radius)
}
func printCircleDetails(radius float64){
	circumference, diameter := getCircleDetails(radius)
	fmt.Println("Radius          :", radius)
	fmt.Println("Circumference   :", circumference)
	fmt.Println("Diameter        :", diameter)
} 
func main(){
	printCircleDetails(5)
	fmt.Println("....................")
	printCircleDetails(10)
}

/** Output

$ go run circle_diameter_and_circumference.go
Radius          : 5
Circumference   : 31.41592653589793
Diameter        : 10
....................
Radius          : 10
Circumference   : 62.83185307179586
Diameter        : 20

*/
