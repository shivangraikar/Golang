//wagp to sort elements using bubble sort

package main

import "fmt"

func main(){
var i,n,j,temp int
var number[30] int

fmt.Println("Enter number of elements")
fmt.Scanln(&n)

for i=0; i<n;i++{
	fmt.Println("Enter element",i+1)
	fmt.Scanln(&number[i])
}

for i=0;i<n-1;i++{
      for j=0;j<n-i-1;j++ {
        if(number[j]>number[j+1]){
           temp=number[j];
           number[j]=number[j+1];
           number[j+1]=temp;
  }
 }
}


fmt.Println(" ")
fmt.Println("Sorted elements after Bubble sort are: ")
for i=0; i<n;i++{
	fmt.Println(number[i])
}
fmt.Println("The number of iterations is",n-1)
}