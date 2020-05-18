//wagp to read n marks of student and print highest, and avg mark

package main

import "fmt"
func main(){
var ele,n,i,sum int 
var data [10]int
fmt.Println("Enter number of subjects")
fmt.Scanln(&n) 

if n>len(data){
	fmt.Println("Array size exceeds the maximum limit")
} else{
	for i:=1;i<=n;i++{
		fmt.Println("Enter marks of",i,"subject")
		fmt.Scanln(&ele)
		data[i]=ele
	}
}
for i=1;i<=n;i++{
	fmt.Println("Marks of subject ",i," is: ", data[i])
}

max := data[0]
for i:=1;i<n;i++ {

        if max < data[i] {

            max = data[i]
        }
    }

fmt.Println("\nHighest marks is", max)

for i=1;i<=n;i++{
	fmt.Scan(data[i])
	sum+=data[i]
}
fmt.Println("Sum of Array elements is", sum)
avg:= sum/n
fmt.Println("Average of Array elements is", avg)
}