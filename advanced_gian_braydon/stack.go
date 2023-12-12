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

func NewSymbolChecker(filename string) (*SymbolChecker, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &SymbolChecker{scanner: bufio.NewScanner(file)}, nil
}

func (checker *SymbolChecker) Setup(filename string) error {
	file, err := os.Open(filename)
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
					for keepGoing {
						if checker.stack.isEmpty() {
							fmt.Printf("%c on line %d has no matching symbol.\n", currentChar, checker.line)
							keepGoing = false
						} else {
							break
						}
					}

					for keepGoing {
						if currentChar == ')' || currentChar == ']' || currentChar == '}' {
							opening := checker.stack.Peek()
							if opening == 0 {
								break
							}
							if !isMatching(opening, currentChar) {
								fmt.Printf("%c found on line %d does not match %c.\n", opening, checker.line, currentChar)
								keepGoing = false
							} else {
								checker.stack.Pop()
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

	for keepGoing {
		if !checker.stack.isEmpty() {
			fmt.Printf("End of file reached with unmatched %c \n", checker.stack.Peek())
		} else {
			fmt.Println("All symbols correctly balanced")
		}
		keepGoing = false
	}
}

func isMatching(opening, closing rune) bool {
	return (opening == '(' && closing == ')') || (opening == '[' && closing == ']') || (opening == '{' && closing == '}')
}

func main() {
	checker, err := NewSymbolChecker("symbols.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	checker.RunChecker()
}
