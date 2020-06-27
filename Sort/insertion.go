//wagp to sort elements using selection sort

package main

import "fmt"

func main(){
var i,j,n,key int
var number[30] int

fmt.Println("Enter number of elements")
fmt.Scanln(&n)

for i=0;i<n;i++{
	fmt.Println("Enter element",i+1)
	fmt.Scanln(&number[i])
}

for i=1;i<n;i++ {  
        key=number[i];  
        j=i-1;  
        if j>=0 && number[j]>key {  
            number[j+1]=number[j];  
            j=j-1;  
        }  
        number[j+1]=key;  
}  

fmt.Println(" ")
fmt.Println("Sorted elements after Selection sort are: ")
for i=0; i<n;i++{
	fmt.Println(number[i])
}
fmt.Println("The number of iterations is", n-1)
}