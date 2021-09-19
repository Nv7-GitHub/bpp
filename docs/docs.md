# B++ Documentation
B++ is a programming language initially developed by the developers for The Brain of TWOW Central. Check out the source code at [their GitHub repository](https://github.com/AeroAstroid/TheBrainOfTWOWCentral)! 

## Table of contents
- [Introduction](#introduction)
- [Hello, World!](#hello-world!)
- [Variables](#variables)
- [Basic Functions](#basic-functions)
- [Comparison](#comparison)
- [Blocks](#blocks)
- [Type Conversions](#type-conversions)
- [Import Statements](#import-statements)
- [Builtin Functions](#builtin-functions)

## Introduction
In B++, everything is a tag. A tag is made of square brackets, with a function call in them! You can also provide tags as input to another tag. Arguments to a tag are seperated by spaces. For example:
```bpp
[MATH 5 * 7]
[CONCAT "hello w" "orld"]
[IF [COMPARE 6 != 4] "6 is not 4" "6 is 4"]
```
You can also do comments using a "#". For example:
```bpp
# This is a comment.
```

## Hello, World!
In B++, the return value of a tag is automatically printed. That means that, to make a hello, world!, you just need to do:
```bpp
"Hello, World!"
```

## Variables
Variables are made using the DEFINE and VAR statements. To define a variable, use:
```bpp
[DEFINE helloworld "Hello, World!"]
```
You can also use DEFINE to change a variable. 

To get the value of a variable, use 
```bpp
[VAR helloworld]
```
We can make a hello world program using variables by doing:
```bpp
[DEFINE helloworld "Hello, World!"]
[VAR helloworld]
```

## Data Types
B++ is a type-safe language. There are 4 types in B++:
- Strings (words/letters)
- Integers (whole numbers)
- Floats (decimals)
- Arrays (lists)

### Strings
Strings can be defined like any other variable. You can get a letter of a string using the INDEX function. For example:
```bpp
[INDEX "Hi!" 0]
```
This gets the first letter of the string "Hi!", or "H". Note that the first letter has an index of 0.

You can also get the length of a string using
```bpp
[LENGTH "Hello, World!"]
```

### Floats and Integers
You can do math on floats an integers, using [the MATH function](#math). 
Integers are defined by doing:
```bpp
[DEFINE a 7]
```
Floats are defined by doing:
```bpp
[DEFINE b 0.21]
```

### Arrays
Arrays can have any type as values. You can even store an array in an array! Define arrays using the ARRAY function. For example, the following program makes an array with the values 1, 2, 3 and 4:
```bpp
[ARRAY 1 2 3 4]
```
You can get a value at an index using.
```bpp
[INDEX [ARRAY 1 2 3 4] 0]
```
This gets the first element in the array. Note that the first element has an index of 0.

You can get the length of an array using 
```bpp
[LENGTH [ARRAY 1 2 3 4]]
```
This would return 4, which is the number of items in the array.

## Basic Functions
There are a few basic functions which you will most likely ue a lot. They are explained below.

### Math
The first one is math. To do math, simply use the MATH tag with a value, operator, and another value. For example:
```bpp
[MATH 100 + 100]
```
Supported operators are:

| Operator | Math Function |
| --- | --- |
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `^` | Power |

### String Formatting
To format a string, you would use the CONCAT function. This accepts any number of strings and concatenates them. For example:
```bpp
[CONCAT "Hello" ", " "World" "!"]
```
Prints "Hello, World!".

## Comparison
To compare values, you use the COMPARE function. For example:
```bpp
[COMPARE 6 = 4]
``` 
In B++, there aren't booleans. COMPARE just returns 1 if true, and 0 if false. 

B++ Supports many comparison operators:
| Operator | What it Does |
| --- | --- |
| `=` | Equals |
| `!=` | Not Equal |
| `>` | Greater Than |
| `<` | Less Than |
| `>=` | Greater Than or Equal To |
| `<=` | Less Than or Equal To |

If statements are ternary. Simply just do:
```bpp
[IF [COMPARE 6 != 4] "6 is not 4" "6 is 4"]
```
To make an if statement. To have more than one instruction in an IF statement, check out [GOTOs](#goto-statements).

## Blocks
Blocks allow block if statements, loops, and functions!

### Block If Statements
What if you need to have multiple lines of code in an if statement? Use a block if statement!
```bpp
[IFB [COMPARE 1 == 1]]
  "Awesome!"
  "Everything works!"
[ELSE]
  "Is 1 not equal to 1?"
  "Thats not good..."
[ENDIF]
```
> Note: You don't need to indent the contents of block statements, but it makes it cleaner and easier to read.

### Loops
B++ supports loops, in the form of `WHILE` loops! For example, to print the numbers 1-100:
```bpp
# Define i, which we will be using to control the number of iterations
[DEFINE i 1]

# While i is less than 100, do something
[WHILE [COMPARE [VAR i] <= 100]]
  # Print i
  [VAR i]

  # Increase i by 1
  [DEFINE i [MATH [VAR i] + 1]]
[ENDWHILE]
```

### Functions
Functions allow code to be put in blocks and run in a safe environment, for example, to add 2 numbers:
```bpp
# Define add, which accepts a, which is an integer, and b, which is also an integer
[FUNCTION ADD [PARAM a INT] [PARAM b INT]]
# Add the two numbers
[DEFINE result [MATH [VAR a] + [VAR b]]]
# Return the value
[RETURN [INT [VAR result]]]
```
Now, to add 1 and 2, using this function, you would just do
```bpp
[ADD 1 2]
```

Side note: You can use 
```bpp
[RETURN [NULL]]
```
To return a blank value.

> :warning: You can't access global variables, only variables defined in the function and the parameters.

B++ supports recursion too! For example, to make a factorial function:
```bpp
# Define factorial function
[FUNCTION FACTORIAL [PARAM inp INT]]

# If the number is over 1, then take the factorial of 1 less than the number and multiply that with the number
[IFB [COMPARE [VAR inp] >= 1]]
  # Take factorial of 1 less than number
  [DEFINE mul [FACTORIAL [MATH [VAR inp] - 1]]]
  # Return the number multiplied by the input
  [DEFINE result [MATH [VAR mul] * [VAR inp]]]
[ELSE] 
  # Otherwise, return 1
  [DEFINE result 1]
[ENDIF]

# Return the result
[RETURN [INT [VAR result]]]
```
Now, just use
```bpp
[FACTORIAL 10]
```
To get 10 factorial, or `3628800`!

## Type Conversions
Sometimes, you need to convert types. To do so, just do:
```bpp
[FLOAT 100]
```
or
```bpp
[STRING 0.1]
```
You can use `INT`, `FLOAT`, `STRING`, and `ARRAY` in any combination!

## Import Statements
B++ also supports multiple files! To do this, use
```bpp
[IMPORT "file.bpp"]
```
This executes the file when the statement is reached. That means conditional imports are also possible, for example
```bpp
[IF [COMPARE [ARGS 0] = "yes"] [IMPORT "yes.bpp"] [IMPORT "no.bpp"]]
```
Make sure to include the other files when running the code!

## Builtin Functions
B++ has many builtin functions, which are listed below.

| Function Signature | Description |
| --- | --- |
| `[CHOOSE val]` | Gets a random index of `val`, which can be an array or a string. |
| `[CHOOSECHAR val]` | Gets a random character of `val`, which must be a string. |
| `[RANDINT lower upper]` | Gets a random integer within the range `lower`, `upper`. |
| `[RANDOM lower upper]` | Gets a random float in the range `lower`, `upper`. |
| `[FLOOR val]` | Gets the floor, or rounds down float `val`. |
| `[CEIL val]` | Gets the ceiling, or rounds up float `val`. |
| `[ROUND val]` | Rounds float `val` to the nearest integer, or whole number. |
| `[ARGS index]` | Gets argument with index `index` of argument array. |
