package main

// QueueNode will store the value and the next node as well
type QueueNode struct {
	val  any
	Next *QueueNode
}

// Queue structure is tell us what our head is and what tail should be with length of the list
type Queue struct {
	head   *QueueNode
	tail   *QueueNode
	length int
}

func New() *Queue {
	//var queue *Queue = new(Queue)
	//return list

	queue := new(Queue)
	return queue // Will return a parameter type of *Queue
}

// enqueue it will be added new value into queue
func (queue *Queue) Enqueue(n any) {
	//var newNode *QueueNode // When pointer tto the QueueNode is used, we use newNode, instead of &newNode  ***
	var newNode QueueNode // create new QueueNode
	newNode.val = n       // set the data

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
func (queue *Queue) Dequeue() any {
	if queue.IsEmpty() {
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
func (queue *Queue) IsEmpty() bool {
	return queue.length == 0
}

// len is return the length of queue
func (queue *Queue) Len() int {
	return queue.length
}

// frontQueue it will return the front data
func (queue *Queue) FrontQueue() any {
	return queue.head.val
}

// backQueue it will return the back data
func (queue *Queue) BackQueue() any {
	return queue.tail.val
}
