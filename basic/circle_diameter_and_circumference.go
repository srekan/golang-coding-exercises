package main
import "fmt"
func main(){
  fmt.Println("-----Find diameter and circumference of a circle-----")
  const Pi = 3.14
  var radius float64
  
  fmt.Print("Enter radius of circle (in meters): ")
  fmt.Scan(&radius)

  circumference := 2 * Pi * radius
  diameter := 2 * radius

  fmt.Println("Radius       :", radius)
  fmt.Println("Circumference:", circumference)
  fmt.Println("Diameter     :", diameter)
}

/** Output - circle_diameter_and_circumference.go

-----Find diameter and circumference of a circle-----
Enter radius of circle (in meters): 5.5
Radius       : 5.5
Circumference: 34.54
Diameter     : 11

*/
