# Glisp
#### Simple lisp implementation on Golang just for fun and self education.
### Code examples
#### Factorial

```
(fn factorial [num]
    (if (eql num 1)
        num
        (mult
            (factorial (sub num 1))
            num)))
(print (factorial 10)) // 3628800
```
#### increment list items
```
(def arr (list 1 2 3 4 5 6 7 8 9))
(print arr) // 1 2 3 4 5 6 7 8 9

(fn inc [i] (sum i 1))
(def arr2 (map inc arr))
(print arr2) // 2 3 4 5 6 7 8 9 10

```
#### sum list items
```
(def nums (list 1 2 3 4 5))

(fn sumNums [nums]
    (if (eql (length nums) 1)
        (get 0 nums)
        (sum
            (sumNums (sublist nums 0 (sub (length nums) 2)))
            (get (sub (length nums) 1) nums))
))

(print (sumNums nums))
```

## Usage
`go build` to build binary

`./glisp run <filepath>` run glisp code from the file

`./glisp eval "<glisp_expression>"` evaluate glisp expression from command line

## Supported syntax
### Operators

#### Calculus
| operation | arguments | description | example |
|---|---|---|---|
| sum | 2...N (num) | add N numbers | `(sum 1 2 3) // 6`|
| sub | 2...N (num) | chain subtraction  for N numbers | `(sub 1 2 3) // -4`|
| mult | 2...N (num) | multiply N numbers | `(mult 1 2 3) // 6` |
| div | 2...N (num) | add N numbers | `(div 25 5 2) // 2.5` |
| if | 3 (bool, successValue, failValue) | if condition call successValue on true and failValue on false | `(if true 5 2) // 5` |
| def | 2 (name, value) | define a variable | `(def a 10) // 10` |
| print | 1 any | print arg to console log | `(print "hello world") // hello world` |

### Logical
| operation | arguments | description | example |
|---|---|---|---|
| if | 3 (bool, successValue, failValue) | if condition call successValue on true and failValue on false | `(if true 5 2) // 5` |
| eql | 2 (v1, v2) | check if variables are equal, type including | `(eql 10 10) // true` `(eql "10" 10) // false` |
| more | 2 (v1, v2) | check if v1 > v2 | `(more 10 5) // true` |
| moreEq | 2 (v1, v2) | check if v1 >= v2 | `(moreEq 10 10) // true` |

#### Defining
| operation | arguments | description | example |
|---|---|---|---|
| def | 2 (name, value) | define a variable | `(def a 10) // 10` |
| fn | 3 (name, [args...], exp) | define a function | `(fn inc [arg] (sum arg 1))` `(inc 5) // 6` |

### List
| operation | arguments | description | example |
|---|---|---|---|
| list | 0...N | create a list value | `(list 1 2 3) // (1 2 3)` |
| map | 2 (fn, list) | you know map | `(fn inc [arg] (sum arg 1)) (map inc (list 1 2 3)) // (2 3 4)` |

#### CMD
| operation | arguments | description | example |
|---|---|---|---|
| print | 1 any | print arg to console log | `(print "hello world") // hello world` |

### Nested expressions
```
(print (sum (div 100 2 (sum 2 3)) (sub (20 7 3)))) // 20
```
### Variables 
```
(def a 5)
(def b (mult a a))
(print (sum a b)) // 30
```
## Contribution
#### Tests
`go test -v ./...`