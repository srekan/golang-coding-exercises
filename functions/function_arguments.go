package main
import "fmt"

func printMessage(name string){
	fmt.Printf("My Name is %s. I am a GO developer. I love Go.\n", name)
}

func main(){
	// We will invoke the function three times, each time passing a different name
	printMessage("Sravan")
	printMessage("Madhuri")
	printMessage("Vihari")
}

/** Output

$ go run function_arguments.go
My Name is Sravan. I am a GO developer. I love Go.
My Name is Madhuri. I am a GO developer. I love Go.
My Name is Vihari. I am a GO developer. I love Go.

*/