//wagp to compute all substrings from a given string 

package main
import "fmt"

func substrings(input string) {
for i:= 0;i<len(input);i++ {
       var substring string = string(input[i])
       fmt.Println(substring)
         for j:=i+1; j<len(input); j++ {
           substring = substring + string(input[j])
           fmt.Println(substring)
        }
    }
}

func main() {
	var input string
	fmt.Println("Enter any String")
	fmt.Scanln(&	input)
	substrings(input)	
}