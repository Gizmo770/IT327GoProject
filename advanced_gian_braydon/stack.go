// Gian Garnica and Braydon Hughes
// Program that implements a stack and a symbol checker using said stack
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Node struct {
	symbol rune //a char in Go is basically a rune
	next   *Node
}

type Stack struct {
	head *Node
}

func (stack *Stack) Push(symbol rune) {
	newHead := &Node{symbol: symbol, next: stack.head}
	stack.head = newHead
}

func (stack *Stack) Pop() (rune, error) { //Do we need multiple return values??
	if stack.isEmpty() {
		return 0, errors.New("Stack is currently empty")
	}

	symbol := stack.head.symbol
	stack.head = stack.head.next
	return symbol, nil
}

func (stack *Stack) Peek() rune {
	if stack.isEmpty() {
		return 0 //0 char means the stack is empty
	}
	return stack.head.symbol
}

func (stack *Stack) isEmpty() bool {
	return stack.head == nil
}

type SymbolChecker struct {
	stack   Stack
	scanner *bufio.Scanner
	line    int
}

func (checker *SymbolChecker) Setup(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	checker.scanner = bufio.NewScanner(file)
	return nil
}

func (checker *SymbolChecker) RunChecker() {
	keepGoing := true
	checker.stack = Stack{}
	checker.line = 1

	for checker.scanner.Scan() {
		line := checker.scanner.Text()
		for _, currentChar := range line {
			for keepGoing {
				if currentChar == '(' || currentChar == '[' || currentChar == '{' {
					checker.stack.Push(currentChar)
				}
				if currentChar == ')' || currentChar == ']' || currentChar == '}' {
					for keepGoing && !checker.stack.isEmpty() {
						switch currentChar {
						case ')':
							if checker.stack.Peek() != '(' {
								fmt.Printf("%c on line %d has no matching symbols.\n", currentChar, checker.line)
								keepGoing = false
							} else {
									checker.stack.Pop()
								break
							}
						case ']':
							if checker.stack.Peek() != '[' {
								fmt.Printf("%c on line %d has no matching symbols.\n", currentChar, checker.line)
								keepGoing = false
							} else {
									checker.stack.Pop()
								break
							}
						case '}':
							if checker.stack.Peek() != '{' {
								fmt.Printf("%c on line %d has no matching symbols.\n", currentChar, checker.line)
								keepGoing = false
							} else {
									checker.stack.Pop()
								}
								break
							}
						}
					}
				}
				break
			}
		}
		checker.line++
	}

	for keepGoing && !checker.stack.isEmpty() {
		fmt.Println("End of file reached with unmatched %c \n", checker.stack.Peek())
		keepGoing = false
	}
	if keepGoing {
		fmt.Println("All symbols correctly balanced")
	}
}

func main() {
	checker := &SymbolChecker{}
	//err := checker.Setup()
	//if err != nil {}
	if err := checker.Setup("symbols.txt"); err != nil { //implicit line termination
		fmt.Println(err)
		return
	}
	checker.RunChecker()
}
