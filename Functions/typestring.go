//wagp to compute all substrings from a given string 

package main
import "fmt"
import "time"

func substrings(input string) {
for i:= 0;i<1;i++ {
       var substring string = string(input[i])
       fmt.Println(substring)
         for j:=i+1; j<len(input); j++ {
           substring = substring + string(input[j])
	   time.Sleep(1* time.Second)
           fmt.Println(substring)
        }
    }
}

func main() {
	var input string
	fmt.Println("Enter any String")
	fmt.Scanln(&input)
	substrings(input)	
}