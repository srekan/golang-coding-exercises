package main
import "fmt"
func main(){
  fmt.Println("----Finding the third angle of a triangle-----")
  var firstAngle float32
  var secondAngle float32

  fmt.Print("Enter first angle: ")
  fmt.Scan(&firstAngle)
  fmt.Print("Enter second angle: ")
  fmt.Scan(&secondAngle)

  thirdAngle := 180 - firstAngle - secondAngle
  fmt.Println("Third Angle:", thirdAngle)
}

/** Output

$ go run third_angle_of_triangle.go
----Finding the third angle of a triangle-----
Enter first angle: 60
Enter second angle: 30
Third Angle: 90

*/