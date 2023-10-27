
# go-hexadeca

Package **hexadeca** implements hexadecimal encoding and decoding, for the Go programming language.
This is meant to be a better alternative to Go's built-in "hex" package.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-hexadeca

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-hexadeca?status.svg)](https://godoc.org/sourcecode.social/reiver/go-hexadeca)

## Example

```go
import "sourceccode.social/reiver/go-hexadeca"

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadeca.EncodeByteUsingLowerCaseSymbols(value)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadeca.EncodeByteUsingUpperCaseSymbols(value)
```

## Import

To import package **hexadeca** use `import` code like the follownig:
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
