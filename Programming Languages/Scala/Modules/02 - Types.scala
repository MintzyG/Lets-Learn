// We are about to discuss basics types in Scala.

def main(args: Array[String]) = {

    // ------------Strings ------------
    // Like so many other programming languages, we have Strings in Scala.
    // Strings are a sequence of caracteres.

    var name: String = "Jonathan" // It's a word
    var number: String = "18" // Note: The "" makes it be a string, unlike numerical values, strings cannot be used in mathematical operations
    //--------------Char------------
    var letter: Char = 'b'; // "A Char data type in Scala can store a single character. Note that 'b' and 'B' are different characters, so be careful when using them in your code."

    // -------------Int-------------
    var value: Int = 15 // Int can stores values up to 32 bits.
    // -------------Long-------------
    var bigValue: Long = 58175431210L // For number larger than 32 bits, you use Long. Note: You need to add a "l" at the end to specify the long value.

    // -------------Float-------------
    var floatValue: Float = 3.14f // Float are number with decimal points, and can store values up to 32 bits. Note that you need to add an 'f' at the end of the number to specify a Float value."


    // -------------Double--------------
    var floatDouble: Double = 3.14 // For number larger than 32 bits. 

    // --------------Boolean--------------
    val isTrue: Boolean = true // Booleans can hold just two values. True and False. They are logical operands. 
    val isFalse: Boolean = false 

}