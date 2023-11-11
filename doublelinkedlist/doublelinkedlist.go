package doublelinkedlist

import "fmt"

// Node Class
type Node struct {
	previousNode *Node
	value        int
	nextNode     *Node
}

// DoubleLinkedList Class
type DoubleLinkedList struct {
	headNode *Node
}

// AddToHead method adds value to head of Double Linked List
func (doublelinkedlist *DoubleLinkedList) AddToHead(value int) {
	var newHeadNode = &Node{}
	newHeadNode.previousNode = nil
	newHeadNode.value = value
	newHeadNode.nextNode = nil
	if doublelinkedlist.headNode != nil {
		newHeadNode.nextNode = doublelinkedlist.headNode
		doublelinkedlist.headNode.previousNode = newHeadNode
	}
	doublelinkedlist.headNode = newHeadNode
}

// Iterate method prints value of Double Linked List
func (doublelinkedlist *DoubleLinkedList) Iterate() {
	var parsingNode *Node
	for parsingNode = doublelinkedlist.headNode; parsingNode != nil; parsingNode = parsingNode.nextNode {
		fmt.Println(parsingNode.value)
	}
}

// LastNode method returns last node of Double Linked List
func (doublelinkedlist *DoubleLinkedList) LastNode() *Node {
	var parsingNode, lastNode *Node
	for parsingNode = doublelinkedlist.headNode; parsingNode != nil; parsingNode = parsingNode.nextNode {
		if parsingNode.nextNode == nil {
			lastNode = parsingNode
			break
		}
	}
	return parsingNode
}

// AddToEnd method adds value to end of Double Linked List
func (doublelinkedlist *DoubleLinkedList) AddToEnd(newValue int) {
	var newLastNode = Node{}
	newLastNode.previousNode = nil
	newLastNode.value = newValue
	newLastNode.nextNode = nil
	var fetchLastNode = doublelinkedlist.LastNode()
	if fetchLastNode != nil {
		fetchLastNode.nextNode = newLastNode
		newLastNode.previousNode = fetchLastNode
	}
}

// NodeBetweenValues method adds value in between nodes of specified values in Double Linked List
func (doublelinkedlist *DoubleLinkedList) NodeBetweenValues(firstValue int, secondValue int) *Node {
	var nodeBetweenWith, parsingNode *Node
	for parsingNode = doublelinkedlist.headNode; parsingNode != nil; parsingNode = parsingNode.nextNode {
		if parsingNode.previousNode != nil && parsingNode.nextNode != nil {
			continue
		}
	}
}
