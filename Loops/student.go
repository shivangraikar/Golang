//wagp to read n marks of student

package main

import "fmt"
func main(){
var n,i int 
var data [10]int
fmt.Println("Enter number of subjects")
fmt.Scanln(&n) 

if n>len(data){
	fmt.Println("Array size exceeds the maximum limit")
} else{
	var ele int
	for i:=1;i<=n;i++{
		fmt.Println("Enter marks of",i,"subject")
		fmt.Scanln(&ele)
		data[i]=ele
	}
}
for i=1;i<=n;i++{
	fmt.Println("Marks of subject ",i," is: ", data[i])
}
}