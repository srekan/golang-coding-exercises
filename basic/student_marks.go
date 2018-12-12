package main
import "fmt"
func main(){
  fmt.Println("-----Student Marks-----")
  var marks1 int
  var marks2 int
  var marks3 int
  var marks4 int
  var marks5 int
  const maxMarksInASubject int = 100

  fmt.Print("Enter Marks in subject 1 of 5: ")
  fmt.Scan(&marks1)
  fmt.Print("Enter Marks in subject 2 of 5: ")
  fmt.Scan(&marks2)
  fmt.Print("Enter Marks in subject 3 of 5: ")
  fmt.Scan(&marks3)
  fmt.Print("Enter Marks in subject 4 of 5: ")
  fmt.Scan(&marks4)
  fmt.Print("Enter Marks in subject 5 of 5: ")
  fmt.Scan(&marks5)

  total := marks1 + marks2 + marks3 + marks4 + marks5
  average := float64(total) / 5
  percentage := 100 * float64(total) / (float64(maxMarksInASubject) * 5)

  fmt.Println("Total    :", total)
  fmt.Printf("Average   : %.2f\n", average)
  fmt.Printf("Percentage: %.2f\n", percentage)
}

/** Output

$ go run student_marks.go
-----Student Marks-----
Enter Marks in subject 1 of 5: 95
Enter Marks in subject 2 of 5: 97
Enter Marks in subject 3 of 5: 98
Enter Marks in subject 4 of 5: 96
Enter Marks in subject 5 of 5: 93
Total    : 479
Average   : 95.80
Percentage: 95.80

*/
