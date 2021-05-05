# B++
A B++ interpreter written in Go! Check docs/docs.md for more information on the B++ language, and docs/programming.md on how to use this as a library, with your code!

## Installation
To install or update B++, do
```bash
go get -u github.com/Nv7-Github/Bpp
```

## Usage
To run a file, do 
```bash
bpp run <filename>
```
For example, to run the `kin` example, do 
```bash
bpp run examples/kin.bpp
```
B++ programs support arguments. To pass arguments, use --args with comma-seperated values. For example:
```bash
bpp run --args arg1,arg2,arg3 <filename>
```

## Compiling Programs
Bpp also supports compiling B++ programs!
> :warning: **Arrays are not supported!**