#include <stdio.h>

int main() {

    // Now that we learned how to print text, lets learn how to read text
    // For that we use scanf() a function from stdio.h
    // First of all, to read a value we need to have some way to store said value.

    int x;

    // When scanf starts it will just wait for the user to give it an input without notifiying them that its waiting
    // So lets warn the user to input a number so they dont just stay there looking at the screen waiting for something to happen

    printf("Enter a base 10 number: ");

    // Now that we warned the user and have where to store our input lets try and get a value from the user
    // Scanf expects us to tell it how to format and where to store the value thats being received
    // For that do a simple conditional format and use &variable to reference the variable you want to use to store your value

    scanf("%i", &x);

    // Be aware that any input will be formatted to what you choose
    // So nothing prohibits a user to input the letter 'B' for example consequently inputing some unkown int as the value

    // Lets print our number
    printf("%i", x);

}
