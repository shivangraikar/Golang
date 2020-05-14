//wagp to accept numbers and sort as positive negative and zeros


package main
import "fmt"

func main() {
	var data[100] int
	var n int
	pc, nc, zc := 0, 0, 0
	fmt.Println("Enter number of elements in array")
	fmt.Scanln(&n)
	
	if n>len(data){
		fmt.Println("Array size exceeds max limit")
	} else{
		for 	i:=0;i<n;i++{
			var ele int
			fmt.Println("Enter any integer to store in array")
			fmt.Scanln(&ele)
			data[i] = ele
		}
	}	

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	for i := 0; i < n; i++ {
		if data[i] < 0 {
			nc++
		} else if data[i] > 0 {
			pc++
		} else {
			zc++
		}
	}
	fmt.Println("\nPositives: ", pc, "\nNegatives: ", nc, "\nZeros: ", zc)
}
