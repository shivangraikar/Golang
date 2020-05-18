//wagp to read marks and the result

package main
import "fmt"

func main(){
var marks int
fmt.Println("Enter the marks of a student")
fmt.Scanln(&marks)

if marks<=40 && marks>=0  {
	fmt.Println("Fail, Try nexy time")
} else if marks>40 && marks<61 {
	fmt.Println("Pass")
} else if marks>60 && marks<76{
	fmt.Println("First class")
} else if marks>75 && marks<101{
	fmt.Println("Distinction")
} else {
	fmt.Println("Invalid marks")
}
}