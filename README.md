# Glisp
### Simple lisp implementation on Golang just for fun and self education.

## Usage
`go build` to build binary

`./glisp run <filepath>` run glisp code from the file

`./glisp eval "<glisp_expression>"` evaluate glisp expression from command line

## Supported syntax
### Operators
| operation | arguments | description | example |
|---|---|---|---|
| sum | 2...N (num) | add N numbers | `(sum 1 2 3) // 6`|
| sub | 2...N (num) | chain subtraction  for N numbers | `(sub 1 2 3) // -4`|
| mult | 2...N (num) | multiply N numbers | `(mult 1 2 3) // 6` |
| sum | 2...N (num) | add N numbers | `(div 25 5 2) // 2.5` |
| print | 1 any | print arg to console log | `(print "hello world") // glisp>> hello world` |