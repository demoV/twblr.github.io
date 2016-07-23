package data_structures

// Linked List Implementation
type LinkedList struct{
	first *Node
	last *Node
	length int
}
type Node struct{
	elem int
	next *Node
}

func (list *LinkedList) Add(elem int) {
	node := Node{elem,nil}
	if(list.length == 0){
		list.first = &node
		list.last = &node
		list.length ++
	}else{
		lastNode := list.first
		for(lastNode != nil){
			if(lastNode.next == nil){
				lastNode.next = &node
				list.last = &node
				list.length ++
				lastNode = &node
			}
			lastNode = lastNode.next
		}
	}

	return
}

func (list *LinkedList) Search(elem int) *Node {
	node := list.first
	for (node != nil) {
		if(node.elem == elem){
			return node
		}
		node = node.next
	}
	return nil
}
