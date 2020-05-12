package main
import "fmt"

func main(){
stack :=[]int{}
var op int

for{
fmt.Println("Enter your choice")
fmt.Println("\n 1:Push\n 2:Pop\n 3:Peek\n 4:Display\n 5:Exit")
fmt.Scanln(&op)
if op == 1{
	var ele int
	fmt.Println("Enter element to be pushed on stack")
	fmt.Scanln(&ele)
	stack = append(stack,ele)
	fmt.Println("Element", ele, "is inserted")
} else if op == 2{
	if len(stack) == 0{
		fmt.Println("Stack is empty")
	} else{
		var ele int
		n:= len(stack)
		ele = stack[n-1]
		stack = stack[0:n-1]
		fmt.Println("Element", ele,"is popped")
	}
} else if op == 3{
	if len(stack) == 0{
		fmt.Println("Stack is empty")
	} else{
		var ele int
		n:= len(stack)
		ele = stack[n-1]
		fmt.Println("Last element is", ele)
	}
} else if op == 4{
	if len(stack) == 0{
		fmt.Println("Stack is empty")
	} else{
		n:= len(stack)
		stack = stack[:n]
		fmt.Println("Queue -->", stack)
	}
} else if op == 5{
	break
} else{
	fmt.Println("Invalid Option")
}
}
}