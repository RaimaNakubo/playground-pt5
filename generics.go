package main

import (
	"fmt"
	"strings"
)

// this is an example of a generic type;
// List represents a node of a singly-linked list that holds value of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// LinkedList represents a singly-linked list itself, that can contain nodes of any type.
type LinkedList[T any] struct {
	head *List[T]
}

func main() {

	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) //compairing a slice of ints to int with function Index()

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) //compairing a slice of strings to string with the same function Index()
	fmt.Println()

	baseSentence := "I sleep holding my pillow."

	stringSentence := strings.Split(baseSentence, " ") //sentence is a slice of strings
	words := initListAsString()                        //words is an empty single-linked list of strings
	//filling words list with strings from sentence
	for wordsCounter := range stringSentence {
		//fmt.Printf("node %v:, content: %v\n", wordsCounter, word)
		words.addStringFront(stringSentence[wordsCounter])
	}
	//now words contains all words from sentence
	fmt.Println("Singly-linked list of strings contents:")
	words.displayListOfStrings()
	fmt.Println()

	//using the same linked list for a different type of data
	byteSentence := []byte(baseSentence)
	bytes := initListAsByte()
	for wordsCounter := range byteSentence {
		bytes.addByteFront(byteSentence[wordsCounter])
	}
	fmt.Println("Singly-linked list of bytes contents:")
	bytes.displayListOfBytes()
}

// this is an example of a generic function;
// function Index returns index of x value inside s slice, or -1 if x value have not been found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s { //v and x both have type T, type T have costraint: "comparable",
		if v == x { //so operator == can be used inside this function
			return i
		}
	}
	return -1
}

// function initListAsString() initializes new empty singly-linked list that contain nodes of string type
func initListAsString() *LinkedList[string] {
	return &LinkedList[string]{}
}

// function initListAsByte() initializes new empty singly-linked list that contain nodes of byte type
func initListAsByte() *LinkedList[byte] {
	return &LinkedList[byte]{}
}

// method addStringFront() adds a new node of argument's value to the end of recieved singly-linked list
func (l *LinkedList[string]) addStringFront(element string) {
	newNode := &List[string]{ // creating a new node for the list
		val:  element, // new node contains value of the element added to the list
		next: nil,     // and the next node of that node is the end of the list
	}

	if l.head == nil { // if the list itself is empty (list do not contain any node yet),
		l.head = newNode // then the zero-node(the head of the list) will become our newly-created node
		return           // and nothing happens after; list now contains a single node: {end of the list, added element's value}
	}

	// if the list itself is not empty (list already contains one or more nodes),
	currentNode := l.head         // then from the head of the list
	for currentNode.next != nil { // we're seaching for a node that is next to the end of the list
		currentNode = currentNode.next
	}
	currentNode.next = newNode // to replace the end of the list with our newly-created node;
	// list now contains: {{{{nil, added element's value}, value N}, value 1}, value 0}
}

// method addByteFront() adds a new node of argument's value to the end of recieved singly-linked list
func (l *LinkedList[byte]) addByteFront(element byte) {
	newNode := &List[byte]{
		val:  element,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

// method displayListOfStrings() prints all content of a recieved singly-linked list of bytes
func (l *LinkedList[string]) displayListOfStrings() {
	currentNode := l.head
	for i := 0; currentNode != nil; i++ {
		fmt.Printf("node: %v, value: %v\n", i, currentNode.val)
		currentNode = currentNode.next
	}
}

// method displayListOfBytes() prints all content of a recieved singly-linked list of strings
func (l *LinkedList[byte]) displayListOfBytes() {
	currentNode := l.head
	buf := []byte{}
	for i := 0; currentNode != nil; i++ {
		fmt.Printf("node: %v, value: %v, character: %q\n", i, currentNode.val, currentNode.val)
		buf = append(buf, currentNode.val)
		currentNode = currentNode.next
	}

	fmt.Println("Result: ", buf)
}
