
# go-hexadeca

Package **hexadeca** implements hexadecimal encoding and decoding, for the Go programming language.
This is meant to be a better alternative to Go's built-in "hex" package.

Package **hexadeca** has functions for hexadecimal encoding and decoding with Go types:
* `byte` (i.e., `uint8`),
* `uint16`,
* `uint32`
* `uint64`
* `int64`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-hexadeca

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-hexadeca?status.svg)](https://godoc.org/sourcecode.social/reiver/go-hexadeca)

## Example

A couple examples of encoding a `byte` to hexidecimal.

```go
import "sourceccode.social/reiver/go-hexadeca"

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadeca.EncodeByteUsingLowerCaseSymbols(value)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadeca.EncodeByteUsingUpperCaseSymbols(value)
```

A couple examples of encoding a `uint16` to hexidecimal.

```go
import "sourceccode.social/reiver/go-hexadeca"

// ...

r3, r2, r1, r0 = hexadeca.EncodeUint16UsingLowerCaseSymbols(value)

// ...

r3, r2, r1, r0 = hexadeca.EncodeUint16UsingLUpperCaseSymbols(value)
```


A couple examples of encoding a `uint32` to hexidecimal.

```go
import "sourceccode.social/reiver/go-hexadeca"

// ...

r7, r6, r5, r4, r3, r2, r1, r0 = hexadeca.EncodeUint32UsingLowerCaseSymbols(value)

// ...

r7, r6, r5, r4, r3, r2, r1, r0 = hexadeca.EncodeUint32UsingLUpperCaseSymbols(value)
```


A couple examples of encoding a `uint64` to hexidecimal.

```go
import "sourceccode.social/reiver/go-hexadeca"

// ...

// value==18364758544493064720 -> r15=='f', r14=='e', r13=='d', r12=='c', r11=='b', r10=='a', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r11, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 = hexadeca.EncodeUint64UsingLowerCaseSymbols(value)

// ...

// value==18364758544493064720 -> r15=='F', r14=='E', r13=='D', r12=='C', r11=='B', r10=='A', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r11, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 = hexadeca.EncodeUint64UsingLUpperCaseSymbols(value)
```

## Import

To import package **hexadeca** use `import` code like the following:
```
import "sourcecode.social/reiver/go-hexadeca"
```

## Installation

To install package **hexadeca** do the following:
```
GOPROXY=direct https://sourcecode.social/reiver/go-hexadeca
```

## Author

Package **hexadeca** was written by [Charles Iliya Krempeaux](http://changelog.ca)
