# B++ Go Library
B++ has 3 libraries: `parser`, `membuild`, and `compiler`. Here is what they do:
| Library Name | What it does |
| --- | --- |
| `parser` | Parses B++ source code and generates an Abstract Syntax Tree (AST), checks types |
| `membuild` | Compiles the AST to an array of functions along with some pre-processing, which you can now run |
| `compiler` | Converts B++ AST into C++ source code  |