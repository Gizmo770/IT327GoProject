/* 
    ulid: jdwald1

    Created on: 4/2/22
*/

// package StackProgram;

import java.util.Scanner;

/**
 * Class that tests the SymbolChecker class. I know it's not JUnit, but I
 * couldn't get the runChecker() method to work for the life of me.
 */
public class SymbolCheckerTest
{
    public static void main(String[] args)
    {
        SymbolChecker symChk = new SymbolChecker();

        Scanner case1 = symChk.setup("advanced_gian_braydon/symbols.txt");

        // System.out.println("Expected output of Case 1: '(' found on line 12 does not match '}'.");
        // System.out.println();
        // System.out.println("Actual output of Case 1: ");
        symChk.runChecker(case1);
        symChk.cleanup(case1);
        System.out.println();
    }
}
