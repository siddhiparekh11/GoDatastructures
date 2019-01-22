package main 


import (

	"fmt"
	"github.com/siddhiparekh11/GoDataStructures/strings"
	"github.com/siddhiparekh11/GoDataStructures/trees"
)


/*type Node struct {

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
}*/

func main() {

	fmt.Println(strings.Permute("abc"))
	//fmt.Println(string(carr[3]))

	arr := []int{18,36,9,12,6,10,8,1}

	var root *trees.Node

	for i:= 0;i<len(arr);i++{
		root = trees.Insert(arr[i],root)
	}

	trees.Inorder(root)



	/*arr := [...]int{121,2,3}

	parr := &arr

	fmt.Println((*parr)[0])

	Append(1)
	Append(2)
	Append(3)
	Append(4)
	Append(5)

	Print()
	fmt.Println(strings.Reverse("siddhi"))*/
}