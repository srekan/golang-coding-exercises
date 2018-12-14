package main
import "fmt"

func printMessage(name, language string){
	fmt.Printf("My Name is %s. I am a %s developer. I love %s.\n", name, language, language)
}

func main(){
	// We will invoke the function three times, each time passing a different name
	printMessage("Sravan", "GO")
	printMessage("Madhuri", "JavaScript")
	printMessage("Vihari", "Python")
}

/** Output

$ go run multiple_arguments.go
My Name is Sravan. I am a GO developer. I love GO.
My Name is Madhuri. I am a JavaScript developer. I love JavaScript.
My Name is Vihari. I am a Python developer. I love Python.

*/