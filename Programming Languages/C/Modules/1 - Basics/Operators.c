// Operator
// Operators are special symbols or keywords used in programming languages to
// manipulate data values or variables.
// In C, there are several types of operators, including:
// - Arithmetic operators, which are used to perform mathematical operations
// like addition, subtraction, multiplication, and division
// - Relational operators, which are used to compare two values and return a
// boolean value (either true or false) depending on the result of the
// comparison
// - Logical operators, which are used to combine multiple conditions and
// return a boolean value based on the result of those conditions
// - Bitwise operators, which are used to perform operations at the bit level
// - Assignment operators, which are used to assign values to variables
// - Increment and decrement operators, which are used to increase or decrease
// the value of a variable by 1
//
// In this tutorial, we'll focus on the arithmetic, relational, and logical
// operators.

// Relational Operators
// C has six relational operators that can be used to compare two values and
// return a boolean value (either true or false) depending on the result of the
// comparison:
//
// - < less than
// - <= less than or equal to
// - > greater than
// - >= greater than or equal to
// - == equal to
// - != not equal to

// Arithmetic Operators
// C has several arithmetic operators that can be used to perform mathematical
// operations on numeric values:
// - + addition
// - - subtraction
// - * multiplication
// - / division
// - % modulus (returns the remainder of a division operation)

// Logical Operators
// C has three logical operators that can be used to combine multiple conditions
// and return a boolean value:
// - && logical AND (returns true if both conditions are true)
// - || logical OR (returns true if either condition is true)
// - ! logical NOT (returns the opposite of a condition)

#include <stdio.h>

int main() {
  int x = 10;
  int y = 5;

  // Arithmetic Operator Example: Addition
  int sum = x + y;
  printf("The sum of x and y is %d\n", sum);

  // Relational Operator Example: Greater Than
  if (x > y) {
    printf("x is greater than y\n");
  } else {
    printf("y is greater than x\n");
  }

  // Logical Operator Example: AND
  if (x > 0 && y > 0) {
    printf("Both x and y are positive numbers\n");
  } else {
    printf("At least one of x and y is not a positive number\n");
  }
  return 0;
}
