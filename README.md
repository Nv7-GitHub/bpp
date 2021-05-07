# B++
A B++ interpreter written in Go! Check [the docs](docs/docs.md) for more information on the B++ language and how it works, and [the programming guide](docs/lib.md) on how to import and use this library, with your code!

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
You can time how long it takes to run a B++ program, using --time or -t. For example:
```
bpp run -t <filename>
```

## Compiling Programs
Bpp also supports compiling B++ programs into a native, extremely high-performance executable! You can use --time or -t with this too.
> :warning: Arrays are not supported!

To compile a program, just do 
```
bpp build <filename>
```
You can also use -o or --output to specify the output file. For example, to compile the `kin` example, do:
```
bpp build -o kin examples/kin.bpp
```
You can also use --preserve or -p to keep the translated C++!
```
bpp build -p <filename>
```