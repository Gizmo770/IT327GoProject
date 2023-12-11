/*
    ulid: jdwald1
    Created on 3/26/22
*/

// package StackProgram;

/**
 * Class that creates a Stack using a singly linked list
 * 
 * @author JD Waldron
 */
public class Stack
{
    /**
     * Class that makes a given character into a Node
     * 
     * @author JD Waldron
     */
    public class Node
    {
        private char charData;
        private Node prev;

        /**
         * Class that makes a character into a Node object without a reference to
         * another Node
         * 
         * @param c The character to make into a Node
         */
        public Node(char c)
        {
            charData = c;
            prev = null;
        }

        /**
         * Class that makes a character into a Node object with a reference to another
         * Node
         * 
         * @param c       The character to make into a Node
         * @param nodeRef The Node that the new Node references
         */
        public Node(char c, Node nodeRef)
        {
            charData = c;
            prev = nodeRef;
        }
    }

    private Node head;

    /**
     * Constructor that sets up a blank Stack object
     */
    public Stack()
    {
        head = null;
    }

    /**
     * Method that checks to see if the stack is empty or not
     * 
     * @return True if stack is empty, false if not
     */
    public boolean isEmpty()
    {
        if (head == null)
        {
            return true;
        }
        return false;
    }

    /**
     * Method that adds a character to the top of the stack
     * 
     * @param item Character to add to the top of the stack
     */
    public void push(char item)
    {
        Node temp = head;
        if (head == null)
        {
            Node n = new Node(item);
            head = n;
        }
        else
        {
            Node n = new Node(item, temp);
            head = n;
        }
    }

    /**
     * Method that returns the character on the top of the stack and removes it
     * 
     * @return The character at the top of the stack
     */
    public char pop()
    {
        if (isEmpty())
        {
            throw new IllegalArgumentException("Stack is empty");
        }

        Node top = head;
        head = head.prev;
        return top.charData;
    }

    /**
     * Method that returns the character on the top of the stack without removing it
     * 
     * @return The character at the top of the stack
     */
    public char top()
    {
        if (!isEmpty())
        {
            return head.charData;
        }
        else
        {
            throw new IllegalArgumentException("Stack is empty");
        }
    }

    /**
     * Method that checks to see if two stacks are equal
     * 
     * @param stk The stack to check the current stack against
     * @return True if the stacks are the same, false if not
     */
    public boolean equals(Stack stk)
    {
        while (!isEmpty() && !stk.isEmpty())
        {
            if (pop() != stk.pop())
            {
                return false;
            }
        }
        if (!isEmpty() || !stk.isEmpty())
        {
            return false;
        }
        return true;
    }

    /**
     * Method that clears the stack
     */
    public void clear()
    {
        while (!isEmpty())
        {
            head = head.prev;
        }
    }

    /**
     * Method that displays the stack on the screen
     */
    public void displayStack()
    {
        if (head != null)
        {
            Node temp = head.prev;
            System.out.println(head.charData);
            while (temp.prev != null)
            {
                System.out.println(temp.charData);
                temp = temp.prev;
            }
            System.out.println(temp.charData);
        }
        else if (isEmpty())
        {
            System.out.println("Stack is empty");
        }
    }
}