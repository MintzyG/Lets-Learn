
// Like in C, every program needs to create a main function.
// Scala has two types of variables: mutable and immutable

def main(args: Array[String]) = {
    // Let's start with val. It is an immutable variable.
    val number: Int = 5 // Once declared as 5, you can't change its value.

    println(number)
    // OUTPUT: 5

    //Var are mutables
    var name = "John" // It starts as John

    name = "Paul" // Name now is getting "Paul"
    println(name)
    // OUTPUT: "Paul"

    // Note: In Scala, you can specify the type before declaring a value to the variable. It's optional.
    var myValue: Int = 15
    print(myValue)
    // OUTPUT: 15    
}