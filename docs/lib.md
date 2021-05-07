# B++ Go Library
B++ has 3 libraries: `parser`, `membuild`, and `compiler`. Here is what they do:
| Library Name | What it does |
| --- | --- |
| [parser](#parser) | Parses B++ source code and generates an Abstract Syntax Tree (AST), checks types |
| [membuild](#membuild) | Compiles the AST to an array of functions along with some pre-processing, which you can now run |
| [compiler](#compiler) | Converts B++ AST into C++ source code  |

## Parser
The parser returns an AST tree when provided with B++ source code. To get the ast tree, you do:
```go
prog, err := parser.Parse(src)
if err != nil {
  panic(err)
}
```
And thats it! Lets look at the AST tree now.

Everything is a statement, even the function calls. Let's look at the interface for a statement:
```go
type Statement interface {
	Line() int
	Type() DataType
}
```
You can also find the types in `parser/types.go`.

Now, lets look at some samples. 

All the statements are stored in a `Program` object. Within the Program, there is an array of statements. Each statement corresponds to a line. 

Functions are added by adding to the `parsers` map. For example, here is the code for the `DEFINE` statement:
```go
parsers["DEFINE"] = StatementParser{
  Parse: func(args []Statement, line int) (Statement, error) {
    return &DefineStmt{
      Label:          args[0],
      Value:          args[1],
      BasicStatement: &BasicStatement{line: line},
    }, nil
  },
  Signature: []DataType{IDENTIFIER, ANY},
}
```
A `StatementParser` has 2 things: it's `Parse` function, which converts the arguments into a statement, and it's signature, which says what types it accepts. 

Data is stored in the `Data` struct. This is what the Data struct looks like:
```go
type Data struct {
	*BasicStatement
	kind DataType
	Data interface{}
}
```
The `Data` value can be type-casted based on the type. To get the type, use `<val>.Type()`. 

Types are stored as bitmasks. This allows functions to accept multiple types. For example, a function can add `STRING | ARRAY` to its signature to accept a string or an array. To check if a type is equal to something, don't use `==`. This will most likely not work. Instead use `a.IsEqual(b)`.

Lets look at the types.
| Type Name | Description |
| --- | --- |
| `STRING` | Stores a string |
| `INT` | Stores an integer |
| `FLOAT` | Stores a `float64` |
| `ARRAY` | Stores an integer |
| `IDENTIFIER` | Stores an identifier, like in a `SECTION`, `DEFINE`, or `VAR` statement |
| `NULL` | It's blank, `Data` is probably `nil` |
| `VARIADIC` | Used for variadic arguments, explained below |
| `NUMBER` | Integer or a float |
| `ANY` | A collection of all types with values |

Now, let's look at the `VARIADIC` type. This allows accepting any number of arguments to a function. This is used in the `ARRAY` function. Let's take a look at that:
```go
parsers["ARRAY"] = StatementParser{
  Parse: func(args []Statement, line int) (Statement, error) {
    return &ArrayStmt{
      Values:         args,
      BasicStatement: &BasicStatement{line: line},
    }, nil
  },
  Signature: []DataType{ANY, VARIADIC},
}
```
Note that the signature for `ARRAY` has `VARIADIC`. When a signature is 2 values long, and has `VARIADIC` as the second value, a function will accept any number of values with the type of the first value. `ARRAY`s can store any type, so the first value in this case is `ANY`.

You can check out the Go Reference for more information on each type of statement.

## Membuild
Info on how to use `membuild` coming soon!

## Compiler
Info on how to use `compiler` coming soon!