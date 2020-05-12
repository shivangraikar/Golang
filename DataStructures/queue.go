package main
import "fmt"

func main(){
queue :=[]int{}
var op int

for{
fmt.Println("Enter your choice")
fmt.Println("\n 1:Insert\n 2:Delete\n 3:Peek\n 4:Display\n 5:Exit")
fmt.Scanln(&op)
if op == 1{
	var ele int
	fmt.Println("Enter element to be added")
	fmt.Scanln(&ele)
	queue = append(queue,ele)
	fmt.Println("Element", ele, "is inserted")
} else if op == 2{
	if len(queue) == 0{
		fmt.Println("Queue is empty")
	} else{
		var ele int
		n:= len(queue)
		ele = queue[0]
		queue = queue[1:n]
		fmt.Println("Element", ele,"is deleted")
	}
} else if op == 3{
	if len(queue) == 0{
		fmt.Println("Queue is empty")
	} else{
		var ele int
		n:= len(queue)
		ele = queue[n-1]
		fmt.Println("Last element is", ele)
	}
} else if op == 4{
	if len(queue) == 0{
		fmt.Println("Queue is empty")
	} else{
		n:= len(queue)
		queue = queue[:n]
		fmt.Println("Queue -->", queue)
	}
} else if op == 5{
	break
} else{
	fmt.Println("Invalid Option")
}
}
}