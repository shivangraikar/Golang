//wagp to read n marks of students and find avg highest and lowest

package main

import "fmt"
import "sort"
func main(){
var n,i int 
var data [10]int
fmt.Println("Enter number of students")
fmt.Scanln(&n) 

if n>len(data){
	fmt.Println("Array size exceeds the maximum limit")
} else{
	var ele int
	for i:=1;i<=n;i++{
		fmt.Println("Enter marks of",i,"student")
		fmt.Scanln(&ele)
		data[i]=ele
	}
}
for i=1;i<=n;i++{
	fmt.Println("Marks of student ",i," is: ", data[i])
}

sort.Ints(data[:])
fmt.Println("\nSorted marks are: ",data)
}