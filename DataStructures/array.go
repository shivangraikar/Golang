//wagp read n integer array from user and display on the screen



package main
import "fmt"

func main(){
var data[100] int
var n int
fmt.Println("Enter number of elements")
fmt.Scanln(&n)
if n>len(data){
	fmt.Println("Array size is max", len(data), "elements")
} else{
	for i:=0; i<n; i++{
	var ele int
	fmt.Println("enter element")
	fmt.Scanln(&ele)
	data[i]=ele
}
}
for i:=0;i<n;i++{
	fmt.Print(data[i]," ")
}
}