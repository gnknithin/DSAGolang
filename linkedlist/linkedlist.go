package linkedlist
import "fmt"

// Node Class
type Node Struct{
	value int
	nextNode *Node
}

// Linked List Class
type LinkedList struct{
	headNode *Node
}

// AddToHead method adds to head of Linked List
func (linkedlist *LinkedList) AddToHead(value int){
	var addingHead = Node{}
	addingHead.value = value
	if addingHead.nextNode != nil{
		addingHead.nextNode = linkedlist.headNode
	}
	linkedlist.headNode = &addingHead
}

// IterateList method iterates over Linked List
func (linkedlist *LinkedList) IterateList(){
	var parsingNode *Node
	for parsingNode= linkedlist.headNode; parsingNode!= nil;parsingNode= parsingNode.nextNode{
		fmt.Println(parsingNode.value)
	}
}

// LastNode method returns last node of Linked List
func (linkedlist *LinkedList) LastNode () *Node{
	var lastNode *Node
	var parsingNode *Node
	for parsingNode=linkedlist.headNode;parsingNode!=nil;parsingNode=parsingNode.nextNode{
		if parsingNode.nextNode == nil{
			lastNode = parsingNode
			break
		}
	}
	return lastNode
}

// AddToEnd method adds to end of Linked List
func (linkedlist *LinkedList) AddToEnd (value int){
	var newLastNode = &Node{}
	newLastNode.value = value
	newLastNode.nextNode = nil
	var fetchLastNode = linkedlist.LastNode()
	if fetchLastNode != nil{
		fetchLastNode.nextNode = newLastNode
	}
}

// NodeWithValue method return Node with given value from LinkedList
func (linkedlist *LinkedList) NodeWithValue (findValue int) *Node{
	var nodeWith *Node
	var parsingNode *Node
	for parsingNode = linkedlist.headNode;parsingNode!=nil;parsingNode=parsingNode.nextNode{
		if parsingNode.value == findValue{
			nodeWith=parsingNode
			break
		}
	}
	return nodeWith
}

// AddAfter method adds a Node after the given value in LinkedList
func (linkedlist *LinkedList) AddAfter (givenValue int, value int){
	var newNode = &Node{}
	newNode.value = value
	newNode.nextNode = nil
	var findNodeWith = linkedlist.NodeWithValue()
	if findNodeWith !=nil{
		newNode.nextNode = findNodeWith.nextNode
		findNodeWith.nextNode = newNode
	}
}