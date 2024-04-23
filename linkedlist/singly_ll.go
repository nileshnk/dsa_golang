package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

const (
	EMPTY_LIST       = "Linked List is Empty"
	POSITION_GREATER = "Position is greater that length of List"
)

func (ll *LinkedList) InsertBegin(value int) {
	if ll.head == nil {
		// insert at begining
		// node := make(Node{data: value, next: nil})
		var node Node = Node{data: value, next: nil}
		ll.head = &node
		return
	}

	// create a new node
	var node Node = Node{data: value, next: ll.head}
	ll.head = &node

}

func (ll *LinkedList) InsertEnd(value int) {
	if ll.head == nil {
		var node Node = Node{data: value, next: nil}
		ll.head = &node
		return
	}

	// create a new node
	var newNode Node = Node{data: value, next: nil}

	// iterate through the list and add the new node
	var curr *Node = ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &newNode
	return
}

func (ll *LinkedList) DeleteFirst() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next

	// By doing the above referencing, the first is no more reachable and can't be referenced by the program. and thus it is eligible for go's garbage collection.
	// however if you want to explicity free the memory, you can achieve so by setting it to nil
	// upon doing this, you should make sure you clean up any additional resources it may hold (like file handles, network connections) before setting to nil

	/*
		In this case, we can proceed as below
		deleteNode = ll.head;
		ll.head = ll.head.next;

		deleteNode = nil
	*/

}

func (ll *LinkedList) DeleteLast() {
	if ll.head == nil {
		return
	}

	var curr *Node = ll.head

	for curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil

}

func (ll *LinkedList) PrintList() {
	var curr *Node = ll.head

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func (ll *LinkedList) InsertAtPosition(value int, pos int) {
	if ll.head == nil {
		if pos > 1 {
			fmt.Println(POSITION_GREATER)
		} else {
			var newNode Node = Node{data: value, next: nil}
			ll.head = &newNode
		}
		return
	}

	var currPos int = 1
	var curr *Node = ll.head

	for curr != nil {
		if currPos >= pos-1 {
			break
		}

		currPos++
		curr = curr.next
	}

	if currPos < pos {
		fmt.Println(POSITION_GREATER)
		return
	}

	if pos == 1 {
		var newNode Node = Node{data: value, next: curr}
		ll.head = &newNode
		return
	}

	var newNode Node = Node{data: value, next: curr.next}

	curr.next = &newNode
}

func (ll *LinkedList) MiddleOfList() {
	if ll.head == nil {
		fmt.Println(EMPTY_LIST)
		return
	}

	var itr1 *Node = ll.head
	var itr2 *Node = ll.head

	for itr2 != nil && itr2.next != nil {
		itr1 = itr1.next
		itr2 = itr2.next.next
	}

	fmt.Println("Middle of list is", itr1.data)
}

func (ll *LinkedList) InsertSorted(value int) {
	newNode := CreateNewNode(value)
	if ll.head == nil {
		fmt.Println(EMPTY_LIST)
		// var newNode Node = new(Node)
		ll.head = newNode
		return
	}

	// iterate through the list
	var curr *Node = ll.head

	if value < curr.data {
		newNode.next = curr
		ll.head = newNode
		return
	}

	for curr.next != nil {
		if curr.data < value && curr.next != nil && curr.next.data > value {
			newNode.next = curr.next
			curr.next = newNode
			break
		}

		curr = curr.next
	}

	if curr.data < value && curr.next == nil {
		// insert at end
		var newNode Node = Node{data: value, next: nil}
		curr.next = &newNode
	}

	return
}

func (ll *LinkedList) NthNode(pos int) {
	if ll.head == nil {
		fmt.Println(EMPTY_LIST)
		return
	}

	// currPos := 0

	var curr *Node = ll.head
	var curr2 *Node = ll.head

	for curr2 != nil && pos > 0 {
		curr2 = curr2.next
		pos--
	}

	for curr2.next != nil {
		curr = curr.next
		curr2 = curr2.next
	}

	fmt.Println("This is the N'th Node", curr.data)
}

func (ll *LinkedList) ReverseLinkedList() {
	if ll.head == nil {
		fmt.Println(EMPTY_LIST)
		return
	}

	var prev *Node = nil
	var curr *Node = ll.head

	for curr != nil {
		var temp *Node = curr.next
		curr.next = prev
		prev = curr
		curr = temp
		fmt.Println("--------", curr)
	}

	ll.head = prev
}

// func (node Node) Node(value int) {
// 	node.data = value;
// 	node.next = nil;
// }

func CreateNewNode(value int) *Node {
	var node *Node = &Node{data: value, next: nil}
	return node
}

func TestFunction() {
	var LL LinkedList
	LL.InsertEnd(10)
	LL.InsertEnd(20)
	LL.InsertEnd(30)
	LL.InsertEnd(40)
	// LL.DeleteFirst()
	// LL.DeleteLast()

	// LL.InsertAtPosition(50, 1)
	// LL.InsertSorted(55)
	LL.NthNode(3)
	LL.MiddleOfList()
	LL.ReverseLinkedList()
	LL.PrintList()
}
