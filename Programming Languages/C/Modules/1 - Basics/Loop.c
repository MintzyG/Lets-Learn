// Welcome to the Control Structures tutorial.
// In this tutorial, we'll cover the use of loops in C programming.
// Loops are a type of control structure that can be used to execute a block of code repeatedly.
// Common examples of loops include for loops, while loops, and do-while loops.

#include <stdio.h>

int main(void) 
{
    // For loops are a type of loop that execute a block of code a specific number of times.
    // For example, let's say we want to print out "Hello, World!" ten times:
    printf("Hello, World! (10 times):\n");
    for(int i = 0; i < 10; i++) {
        printf("Hello, World!\n");
    }

    // While loops are a type of loop that execute a block of code while a certain condition is true.
    // For example, let's say we want to count from 0 to 10 using a while loop:
    int i = 0;

    printf("Counting from 0 to 10 (while loop):\n");
    while(i <= 10) {
        printf("%d ", i);
        i++; // i++ is the same as i += 1 or i = i + 1
    }
    printf("\n");

    // Do-while loops are a type of loop that execute a block of code at least once, and then continue executing
    // the code while a certain condition is true.
    // For example, let's say we want to ask the user for a positive integer using a do-while loop:
    int n = 0;

    do {
        printf("Please enter a positive integer: ");
        scanf("%d", &n);
    } while(n <= 0);

    // I recommend playing around with these examples and modifying the code to see how loops can be used in different ways.
    // Don't be afraid to experiment and try out new things!
    // Use a C compiler like GCC or Clang to compile your C code into an executable file.
    return 0;
}
