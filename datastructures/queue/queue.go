package queue

// Queue Linked-List

// Node will be store the value and the next node as well
type Node struct {
	val  any
	Next *Node
}

// Queue structure is tell us what our head is and what tail should be with length of the list
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

// enqueue it will be added new value into queue
func (queue *Queue) enqueue(n any) {
	var newNode Node // create new Node
	newNode.val = n  // set the data

	if queue.tail != nil {
		queue.tail.Next = &newNode
	}

	queue.tail = &newNode

	if queue.head == nil {
		queue.head = &newNode
	}
	queue.length++
}

// dequeue it will be removed the first value into queue (First In First Out)
func (queue *Queue) dequeue() any {
	if queue.isEmpty() {
		return -1 // if is empty return -1
	}
	data := queue.head.val

	queue.head = queue.head.Next

	if queue.head == nil {
		queue.tail = nil
	}

	queue.length--
	return data
}

// isEmpty it will check our list is empty or not
func (queue *Queue) isEmpty() bool {
	return queue.length == 0
}

// len is return the length of queue
func (queue *Queue) len() int {
	return queue.length
}

// frontQueue it will return the front data
func (queue *Queue) frontQueue() any {
	return queue.head.val
}

// backQueue it will return the back data
func (queue *Queue) backQueue() any {
	return queue.tail.val
}
