package main 


import (

	"fmt"
)


type Node struct {

	data int
	next *Node
}


var head *Node

func Append(val int) {

	node := Node {val,nil}
	if(head==nil){
		head = &node
	}else{
		temp := head
		for {
				if temp.next==nil {
					break
				}
				temp=temp.next
		}
		temp.next=&node
	}

}

func Print(){

	if head==nil {
		fmt.Println("The list is empty")
		return
	}

	temp := head

	for {
		fmt.Println(temp.data)
		if temp.next == nil {
			break
		}
		temp= temp.next
	}
}

func main() {

	carr := []rune("siddhi")
	fmt.Println(string(carr[3]))

	Append(1)
	Append(2)
	Append(3)
	Append(4)
	Append(5)

	Print()
}