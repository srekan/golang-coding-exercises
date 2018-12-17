package main
import "fmt"
func printOdds(limit int){
	for i := 1; i <= limit; i += 2 {
		fmt.Println(i)
	}
}
func main(){
	printOdds(100)
}
