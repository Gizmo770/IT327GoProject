// Gian Garnica and Braydon Hughes
// Program that implements a stack and a symbol checker using said stack
package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	symbol rune  //char in stack
	next   *Node //pointer to next char in stack
}

type Stack struct {
	head *Node //char at the top of the stack
}

func (stack *Stack) Push(symbol rune) {
	newHead := &Node{symbol: symbol, next: stack.head} //Creates new instance of Node struct using data from stack struct and parameter
	stack.head = newHead                               //Adds the Node to top of stack
}

func (stack *Stack) Pop() (rune, error) {
	if stack.isEmpty() { //Checks to make sure stack has values in it
		return 0, errors.New("Stack is currently empty")
	}

	symbol := stack.head.symbol //If stack has values, points to value after current value and returns current value
	stack.head = stack.head.next
	return symbol, nil
}

func (stack *Stack) Peek() rune {
	if stack.isEmpty() {
		return 0 //0 equates to an empty stack since we are just using symbols
	}
	return stack.head.symbol //Shows item at top of stack without removing it
}

func (stack *Stack) isEmpty() bool {
	return stack.head == nil //Checks for pointer to first item in stack
}

type SymbolChecker struct { //Struct that holds the Stack, the reader of the data, and a counter for lines
	stack   Stack
	scanner *bufio.Scanner
	line    int
}

func NewSymbolChecker(input string) (*SymbolChecker, error) {
	return &SymbolChecker{scanner: bufio.NewScanner(strings.NewReader(input))}, nil //Creates new SymbolChecker instance with scanner initialized to read the inputted string
}

func (checker *SymbolChecker) RunChecker() { //Function that parses string and actually does the symbol checking
	keepGoing := true
	checker.stack = Stack{} //Initializes empty stack
	checker.line = 1

	for checker.scanner.Scan() { //Scanner moves to next token
		line := checker.scanner.Text()     //Gets current token
		for _, currentChar := range line { //Iterates over each character in the line
			for keepGoing {
				if currentChar == '(' || currentChar == '[' || currentChar == '{' { //If open symbol is found, it's pushed to the stack
					checker.stack.Push(currentChar)
				}
				if currentChar == ')' || currentChar == ']' || currentChar == '}' { //If closing symbol is found, checks begin
					for keepGoing {
						if checker.stack.isEmpty() { //Empty stack means there is no matching symbol
							fmt.Printf("%c on line %d has no matching symbol.\n", currentChar, checker.line)
							keepGoing = false
						} else {
							break
						}
					}

					for keepGoing {
						if currentChar == ')' || currentChar == ']' || currentChar == '}' {
							opening := checker.stack.Peek() //Check character at top of stack
							if opening == 0 {
								break
							}
							if !isMatching(opening, currentChar) { //If closing symbol doesn't match character at the top of stack, mismatch has occurred
								fmt.Printf("%c found on line %d does not match %c.\n", opening, checker.line, currentChar)
								keepGoing = false //Loop should exit
							} else {
								checker.stack.Pop() //Closing symbol does match opening symbol, remove from the top of the stack and continue
								break
							}
						}
					}
				}
				break
			}
		}
		checker.line++ //Move to next line to parse
	}

	for keepGoing {
		if !checker.stack.isEmpty() { //If finished parsing and symbols remain on stack, show what remains unmatched by popping rest of stack
			fmt.Print("End of file reached with unmatched symbols: ")
			for !checker.stack.isEmpty() {
				if symbol, err := checker.stack.Pop(); err == nil {
					fmt.Printf("%c ", symbol)
				} else {
					fmt.Println(err)
				}
			}
			fmt.Println()
		} else {
			fmt.Println("All symbols correctly balanced") //Otherwise, parsing completed and symbols all matched
		}
		keepGoing = false
	}
}

func isMatching(opening, closing rune) bool { //Simple helper that compares opening and closing symbols and returns true if matching, false if not
	return (opening == '(' && closing == ')') || (opening == '[' && closing == ']') || (opening == '{' && closing == '}')
}

func main() {
	//Test cases: for each test case, a string is created that is added to a new SymbolChecker instance, checked to ensure it was
	//created successfully, and then run against expected output.

	//Case 1: symbols all properly balanced
	balanced := "(Is this){balanced}[enough?]"
	balancedChecker, err := NewSymbolChecker(balanced)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balanced Symbol Test: ")
	fmt.Println("Expected: All symbols correctly balanced")
	balancedChecker.RunChecker()

	//Case 2: Open symbols left, each without a closing symbol
	unbalancedOpen := "{Is this}[balanced{[enough?]"
	unbalancedOpenChecker, err := NewSymbolChecker(unbalancedOpen)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Unbalanced with open Symbol Test: ")
	fmt.Println("End of file reached with unmatched symbols: { [")
	unbalancedOpenChecker.RunChecker()

	//Opening and closing symbols do not match each other
	mismatched := "{Is this}(balanced}[enough?]"
	mismatchedChecker, err := NewSymbolChecker(mismatched)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Unmatched Symbol Test: ")
	fmt.Println("Expected: ( found on line 1 does not match }.")
	mismatchedChecker.RunChecker()

	//Closing symbol does not match to any opening symbol, also shows multi-line use
	unbalancedClosed := "(){[]()}\n{}{}\n]{}"
	unbalancedClosedChecker, err := NewSymbolChecker(unbalancedClosed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Multi-line symbol test: ")
	fmt.Println("Expected: ] on line 3 has no matching symbol")
	unbalancedClosedChecker.RunChecker()

}
