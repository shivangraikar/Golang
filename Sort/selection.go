//wagp to sort elements using selection sort

package main

import "fmt"

func main(){
var i,j,n,pos,temp int
var number[30] int

fmt.Println("Enter number of elements")
fmt.Scanln(&n)

for i=0;i<n;i++{
	fmt.Println("Enter element",i+1)
	fmt.Scanln(&number[i])
}

for i=0;i<n-1;i++{  
        pos = i;  
        for j=i+1;j<n;j++{  
        	if number[j]<number[pos]{  
            		pos=j;  
	    		}  
		}
        if pos!=j{
		temp=number[i]
		number[i]=number[pos]
		number[pos]=temp
		}   
	}  

fmt.Println(" ")
fmt.Println("Sorted elements after Selection sort are: ")
for i=0; i<n;i++{
	fmt.Println(number[i])
}
}