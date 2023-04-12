// Welcome to the C tutorial.
// What I'm writing after each "//" are comments, they're ignored by the compiler and don't interfere with the code.
// Use them to make your code more readable by specifying what each thing does, or even tag each code block for personal cataloguing.
// Code can get confusing after hundreds of lines and non-linear implementation.

#include <stdio.h>

// Including stdio.h is required, C does not have a built-in library, instead borrowing from the assembly
// libraries included in the OS.
// Also, you always have to input ; at the end of each statement, this tells the compiler you've ended that statement.

int main() {
	
	// Before beggining to code, you have to add the main function, it tells the OS when your program runs.
	// Int is an integer value, which means it requires returning an integer to finalize the code.
	
	printf ("Hello World."); 
    	// This is printf, a function that allows us to take a string and optional arguments and produce
    	// a formatted sequence of characters for output
	// in the example, we're inputting the standard Hello World.
	
	return 0;
    	// If the program returns 0, it means that it has been ran succesfully
	// Otherwise it will return a non-zero value, meaning an error.
    	// Int main() is mandatory for your code to run.
}

// I recommend rewriting the code yourself by memory to get accostumed to each function.
// Use GCC or Clang compiler to compile your C code into an executable file
