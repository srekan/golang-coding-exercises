package main
import (
	"fmt"
	"math"
)

func getCircleDetails(radius float64)(circumference, diameter float64){
	circumference = (2 * math.Pi * radius)
	diameter = (2 * radius)
	return
}
 
func main(){
	var radius float64 = 5
	circumference, diameter := getCircleDetails(radius)
	fmt.Println("Radius          :", radius)
	fmt.Println("Circumference   :", circumference)
	fmt.Println("Diameter        :", diameter)
}

/** Output

$ go run named_returns.go
Radius          : 5
Circumference   : 31.41592653589793
Diameter        : 10

*/
