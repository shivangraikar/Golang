//wagp to read rating and print feedback

package main
import "fmt"

func main(){
var rating int
fmt.Println("Enter the rating between 1-5 inclusive")
fmt.Scanln(&rating)

switch(rating){

case 5,4: fmt.Println("Excellent")

case 3,2: fmt.Println("Good")

case 1: fmt.Println("Poor")

default: fmt.Println("Invalid")
}
}