//wagp to read integer and find if number is even or odd using switch

package main
import "fmt"

func main() {
var even int
var num int

fmt.Println("Enter a number")
fmt.Scanln(&num)
even= num%2

switch(even){
case 0: fmt.Println("Number is even")

case 1: fmt.Println("Number is odd")

default: fmt.Println("Invalid number")
}
}