// Welcome to the C tutorial.
// Naturally, you can already see how to create comments.

#include <stdio.h>

// Including stdio.h is required, C does not have a built-in library, instead borrowing from the assembly
// libraries included in the OS.
// Also, you always have to input ; at the end of each statement, this tells the compiler you've ended that statement.

int main() {
	// Before beggining to code, you have to add the main function, it tells the OS when your program runs.
	// Int is an integer value, which means it requires returning an integer to finalize the code.
	printf ("Hello World"); // Our usual Hello World.
	return 0; //If the program returns 0, it means that it has been ran succesfully, if not, we encountered an error.
}


// Use a GCC compiler to compile C into machine code, if needed.
