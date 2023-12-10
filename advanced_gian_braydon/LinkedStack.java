package dataStructures;

/**
 *	LinkedStack class
 *
 *	implementation of a stack based on a linked structure
 *
 *	Written by Roger West, edited by Elham Buxton
 *	Adapted from code written by Mark Weiss in Data Structures and Problem
 *	Solving Using Java, 2nd edition, 2002
 */
public class LinkedStack<T> implements Stack<T>
{

	/*************
	 *	attributes
	 ************/
	
	/** reference to top node of stack */
	private StackNode<T> topOfStack;
	
	/** current number of items in stack */
	private int theSize;


	/***************
	 *	constructors
	 **************/

	/**
	 *	return a new, empty LinkedStack
	 */
	public LinkedStack()
	{
		// empty this LinkedStack
		clear();
	}
	
	
	/**************************************
	 *	methods inherited from class Object
	 *************************************/

	/**
	 *	return a String representation of the LinkedStack
	 *
	 *	items are listed in a single column, with the item at the top of the column
	 *	representing the top of the stack
	 */
	public String toString()
	{
		String s = "";
		StackNode<T> node = topOfStack;
		
		while(node != null)
		{
			s += node.item + "\n";
			node = node.next;
		}
		
		return s;
	}


	/*****************************************
	 *	methods inherited from interface Stack
	 ****************************************/
	
	/**
	 *	return the item at the top of the stack; the item is not removed
	 *
	 *	throws UnderflowException if stack is empty
	 */
	public T peek()
	{
		// throw exception if this LinkedStack is empty
		if(isEmpty())
			throw new UnderflowException("LinkedStack peek");

		// return item at top of this LinkedStack
		return topOfStack.item;
	}


	/**
	 *	remove the item at the top of the stack
	 *
	 *	throws UnderflowException if stack is empty
	 */
	public T pop()
	{
		// throw exception if this LinkedStack is empty
		if(isEmpty())
			throw new UnderflowException("LinkedStack pop");

		// store item currently at top of this LinkedStack
		T obj = topOfStack.item;
		
		// advance top of this LinkedStack to next item
		topOfStack = topOfStack.next;
		
		// subtract 1 from size of this LinkedStack
		theSize--;
		
		// return item that was removed
		return obj;
	}


	/**
	 *	add the specified item to the top of the stack
	 */
	public boolean push(T obj)
	{
		// create new node, and set top of this LinkedStack to reference new node
		topOfStack = new StackNode<T>(obj, topOfStack);
		
		// add 1 to size of this LinkedStack
		theSize++;
		
		// push successful
		return true;
	}


	/**********************************************
	 *	methods inherited from interface Collection
	 *********************************************/
	
	/**
	 *	add the specified item to the top of the stack
	 */
	public boolean add(T obj)
	{
		// invoke push alias method
		push(obj);
		
		// add successful
		return true;
	}


	/**
	 *	remove the item at the top of the stack
	 *
	 *	throws UnderflowException if stack is empty
	 */
	public boolean remove()
	{
		// invoke pop alias method
		pop();
		
		// remove successful
		return true;
	}


	/**
	 *	empty the LinkedStack
	 *
	 *	stack will be reset to the default capacity
	 *	size will be set to zero
	 */
	public void clear()
	{
		// reset reference to top of this LinkedStack
		topOfStack = null;
		
		// reset size to 0
		theSize = 0;
	}


	/**
	 *	return the number of items in the ArrayList
	 */
	public int size()
	{
		return theSize;
	}


	/**
	 *	return true if the ArrayStack is empty
	 */
	public boolean isEmpty()
	{
		return theSize == 0;
	}


	/**
	 *	return the item at the top of the stack; the item is not removed
	 *
	 *	throws UnderflowException if stack is empty
	 */
	public T get()
	{
		// invoke peek alias method
		return peek();
	}
	

	/****************
	 *	inner classes
	 ***************/

	/**
	 *	nested class StackNode
	 *
	 *	encapsulates the fundamental building block of a LinkedStack
	 *	contains a data item, and a reference to the next node in the stack
	 */
	private  class StackNode<E>
	{

		/*************
		 *	attributes
		 ************/
		
		/** the data item */
		E item;
		
		/** reference to the next node in the stack */
		StackNode<E> next;


		/***************
		 *	constructors
		 **************/

		/**
		 *	create a new StackNode containing the specified item, and reference
		 *	to the next node in the stack
		 */
		public StackNode(E obj, StackNode<E> node)
		{
			item = obj;
			next = node;
		}

	}

}