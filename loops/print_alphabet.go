package main
import "fmt"
func printAlphabet(){
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c\n", char)
	}
}
func main(){
	printAlphabet()
}

/** Output

$ go run print_alphabet.go
a
b
c
d
e
f
g
h
i
j
k
l
m
n
o
p
q
r
s
t
u
v
w
x
y
z

*/