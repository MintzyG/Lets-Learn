#include <stdio.h>

int main() {

    // Now that we learned how to print text, lets learn how to read text
    // For that we use scanf() a function from stdio.h
    // Primeiramente para poder ler um valor temos que ter como armazena-lo

    int x;

    // Now that we have where to store our input lets try and get a value
    // Scanf expects us to tell it how to format our value and where to store it
    // For that do a simple format and use &variable to reference the variable you want to use to store your value

    scanf("%i", &x);

    // Be aware that any input will be formatted to what you choose
    // So nothing prohibits a user to inpu the letter 'B' for example consequently inputing some unkown int as the value
    
    // Lets print our input

    printf("%i", x);

}
