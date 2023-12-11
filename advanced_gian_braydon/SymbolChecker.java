/*
    ulid: jdwald1
    Created on 3/26/22
*/

// package StackProgram;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

/**
 * Class that checks the symbols of a given file using a stack
 * 
 * @author JD Waldron
 */
public class SymbolChecker
{
    private Stack stk;
    private Scanner in;
    private int lineNumber;

    /**
     * Constructor that sets up the stack and sets the line number to 1
     */
    public SymbolChecker()
    {

    }

    public static void main(String[] args)
    {
        SymbolChecker symChk = new SymbolChecker();
        symChk.setup("advanced_gian_braydon/symbols.txt");
    }

    /**
     * Method that sets up the Scanner and returns it for future use
     * 
     * @param filename Name of the file to check
     * @return The Scanner with the file
     */
    public Scanner setup(String filename)
    {

        try
        {
            in = new Scanner(new File(filename));
        }
        catch (FileNotFoundException e)
        {
            System.out.println(e);
            e.printStackTrace();
        }
        return in;
    }

    /**
     * Method that checks the file for matching brackets
     * 
     * @param in The Scanner with the file to check
     */
    public void runChecker(Scanner in)
    {
        boolean keepGoing = true;
        stk = new Stack();
        lineNumber = 1;

        while (in.hasNextLine())
        {
            String line = in.nextLine();
            char currChar;
            for (int i = 0; i < line.length(); i++)
            {
                while (keepGoing)
                {
                    currChar = line.charAt(i);

                    // If the current character is an open bracket, push it to the stack
                    if (currChar == '(' || currChar == '{' || currChar == '[')
                    {
                        stk.push(currChar);
                    }

                    // If the current character is a closing bracket, start checking
                    if (currChar == ')' || currChar == '}' || currChar == ']')
                    {
                        while (keepGoing)
                        {
                            // If the stack is empty, print that there is no matching symbol
                            if (stk.isEmpty())
                            {
                                System.out.println(
                                        "'" + currChar + "' on line " + lineNumber + " has no matching symbol.");
                                keepGoing = false;
                            }
                            else
                            {
                                break;
                            }
                        }
                        while (keepGoing)
                        {
                            if (currChar == ')')
                            {
                                // If the current character is ')', but the character at the top of the stack
                                // does not match it, end the program and print why it was ended
                                if (stk.top() != '(')
                                {
                                    System.out.println("'" + stk.top() + "' found on line " + lineNumber
                                            + " does not match '" + currChar + "'.");
                                    keepGoing = false;
                                }
                                // If the character does match, remove it from the top of the stack and keep
                                // going
                                else if (stk.top() == '(')
                                {
                                    stk.pop();
                                    break;
                                }
                            }

                            if (currChar == '}')
                            {
                                // If the current character is '}', but the character at the top of the stack
                                // does not match it, end the program and print why it was ended
                                if (stk.top() != '{')
                                {
                                    System.out.println("'" + stk.top() + "' found on line " + lineNumber
                                            + " does not match '" + currChar + "'.");
                                    keepGoing = false;
                                }
                                // If the character does match, remove it from the top of the stack and keep
                                // going
                                else if (stk.top() == '{')
                                {
                                    stk.pop();
                                    break;
                                }
                            }

                            if (currChar == ']')
                            {
                                // If the current character is ']', but the character at the top of the stack
                                // does not match it, end the program and print why it was ended
                                if (stk.top() != '[')
                                {
                                    System.out.println("'" + stk.top() + "' found on line " + lineNumber
                                            + " does not match '" + currChar + "'.");
                                    keepGoing = false;
                                }
                                // If the character does match, remove it from the top of the stack and keep
                                // going
                                else if (stk.top() == '[')
                                {
                                    stk.pop();
                                    break;
                                }
                            }
                        }
                    }
                    break;
                }
            }
            lineNumber++;
        }
        // Cases for if the end of the file is reached
        while (keepGoing)
        {
            // If at the end with elements still in the stack, print out the element on top
            if (!stk.isEmpty())
            {
                System.out.println("End of file reached with unmatched '" + stk.top() + "'.");
            }
            // Otherwise, print that everything matches
            else
            {
                System.out.println("All symbols correctly balanced.");
            }
            keepGoing = false;
        }
    }

    /**
     * Method that closes the Scanner
     * 
     * @param in The Scanner to close
     */
    public void cleanup(Scanner in)
    {
        in.close();
    }

    public void printFile(Scanner in)
    {

    }
}
