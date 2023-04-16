#include <stdio.h>

int main() {

// C, just like any other language has data types, they are used to declare the type and size of a given variable
//
// int myInt;
// Here for example myInt is of the type int (used to store intergers) and has a size of 4 bytes
//
// C has a lot of data types, but here we will explain only the most commonly used ones
//
//---------------------int
// int is defined by an Integer value which is a whole number, either positive, negative or zero.

    int x = -1;
    int y = 0;
    int z = 1;

    printf("%i, %i, %i \n", x, y, z);


//---------------------char
// Char is defined by characters, it allows for ASCII characters such as letters.
// The characters are defined by numbers, char uses integer values to define characters.
// This allows us to represent the letter 'A' as 65 and 'a' as 97 for example

    char A = 65;
    char B = 'B'; // Would've been 66
    char a = 97;
    char b = 'b'; // Would've been 98

    printf("%c, %c, %c, %c \n", A, B, a, b);

    // If you want to write the character instead of the number use '' simple quotes


//---------------------float
// Like Integers, this uses numbers, however, we can now use non-whole numbers, such as 6.7.

    float f1 = -0.1;
    float f2 = 0.0;
    float f3 = 0.1;

    printf("%f, %f, %f \n", f1, f2, f3);

    // Trying to write a float with a very big decimal part will cause troubles
    // because the computer can't achieve such accuracy
    // In the example below the number is rounded to 0.00000000
    // So be mindful of your floats

    float f4 = 0.0000001;

    printf("%f \n", f4);


//---------------------double
// This works just like float, however it takes double the usual memory, allowing you to write high precision numbers.
// The double then avoids the problem the float has, only losing accuracy after the 15th decimal place
// It may not be fully printed in the console but it does store more data

    double d1 = 0.00234;

    printf("%lf \n", d1);

// A in depth explanation of other C data types can be found in the 2.1 submodule of Variables
// In here, to keep it brief we only mentioned the most commonly used ones
}
