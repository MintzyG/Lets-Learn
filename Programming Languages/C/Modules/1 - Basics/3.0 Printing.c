#include <stdio.h>
// Since we are going to print to the console it's imperative that we include 'stdio.h'
// as it's the header file that contains the printf() and scanf() functions

int main() {

    // Printing to the console is as easy as doing a printf() with text inside
    // * the text has to be inside double quotation marks, otherwise it's interpreted as a variable input
    
    printf("Hello World!");

    // Printf will only print to the same line, if we want a new line we have to specify it by using \n inside the text
    // You don't have to place a space between \n and your text, C knows to only read \n as the new line and not your actual text

    printf("\nHello! \nWorld!");
    // In this example, \n will act as if we pressed enter before typing the word in a text input program such as notepad or word
    // If you want to see it in action, try compiling this code.
    // Furthermore we can pass in variables to the printf() function to print them
    // And they don't need to be all characters, they can be, int, float or other types

    int x = 10;
    char letter = 'F';
    float f = 1.42;

    printf("%i \n", x);

    // The printf function only prints characters so if we want to print an int for example we have to format it
    // thats why you see the %i above, thats a way to format our int into a char to print it
    // formatting is as simple as putting '%' and following it with a letter to represent the type that will be formatted
    // Let's print all our types

    printf("Char: %c \n", letter);
    printf("Float: %f \n", f);
    printf("Int: %i \n", x);

    // A good thing formatting provides us with is multiple variable printing
    // Meaning we can pass more than one variable into printf and print them all at once using formatting

    printf("One line: %i, %c, %f", x, letter, f);

    // NOTE: the order here matters, if the variable you passed doesn't match the formatted symbol it will give you an error

}
