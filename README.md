# B++
A B++ interpreter written in Go! 

## Installation
To install or update B++, do
```bash
go get -u github.com/Nv7-Github/Bpp
```

## Basic Usage
To run a file, do 
```bash
bpp -file filename
```
For example, to run the `kin` example, do 
```bash
bpp -file examples/kin.bpp
```

## Using this in your own code
You can use this in your own code. To do this, first, compile the program. You can do this by doing 
```go
prog, err := parser.Parse(src)
```
Handle `err` to find compile-time errors. Then, to execute the program and get the output, do 
```go
out, err := prog.Run()
```
Handle `err` to find runtime errors.

## Custom Executors and Debugging
In the example, we have been using `prog.Run` to execute the code. However, `prog.Run` is just a basic Executor. Lets look into how to make an Executor.

A program has 2 fields in it's struct. `Program` and `Memory`. `Memory` is a `map` in which the key is variable name, and the value is the value of the variable. This value is of type `Variable`.

A `Variable` contains 2 fields: `Type` and `Data`. To compare types, you use 
```go
type1.IsEqual(type2)
```

To make a custom executor, we use the `Program` value. The `Program` is a set of instructions, which are called `Executables`. These are simply functions that return a `Variable`.

The most basic executor would look like
```go
for _, instruction := range prog.Program {
  output, err := instruction(prog)
  if err != nil {
    panic(err)
  }
  fmt.Println(output.Data)
}
```

The builtin executor has some modifications to create output like the original B++. Let's look at some of these.

One of the simplest modifications we can do is not print out `NULL` values. We can do this simply by doing
```go
for _, instruction := range prog.Program {
  output, err := instruction(prog)
  if err != nil {
    panic(err)
  }
  if !output.Type.IsEqual(parser.NULL) {
    fmt.Println(output.Data)
  }
}
```

One of the most important modifications is the modification to printing `ARRAY`s. B++ prints arrays much differently from Go's `fmt.Println`. Let's look at the modification.
```go
for _, instruction := range prog.Program {
  output, err := instruction(prog)
  if err != nil {
    panic(err)
  }
  if output.Type.IsEqual(ARRAY) {
    fmt.Print("[ARRAY")
    for _, val := range ret.Data.([]Variable) {
      fmt.Print(" ")
      fmt.Print(val.Data)
    }
    fmt.Print("]\n")
    continue
  }
  fmt.Println(output.Data)
}
```
This modification checks if the data is an array. If it is, it type casts the data to an array of `Variable` and loops through them, adding the data to the output with a space.

You can look in `parser/run.go` for more modifications.