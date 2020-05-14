//wagp to find factorial of a number using function


package main
import "fmt"

func fact(num int) (f int) {
f = 1
for i:=1;i<=num;i++{
	f=f*i;
} 
return
}

func main(){
var num int
fmt.Println("Enter a number")
fmt.Scanln(&num)
if num<0{
	fmt.Println("Number is negative")
} else if num==0 || num==1{
	fmt.Println("Factorial is 1")	
} else {
	ans:= fact( num )
	fmt.Println("Factorial is", ans)
}
}
