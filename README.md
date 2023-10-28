
# go-hexadeca

Package **hexadeca** implements **hexadecimal** _encoding_ and _decoding_, for the Go programming language.
This is meant to be a better alternative to Go's built-in `"hex"` package.

Package **hexadeca** does _not_ impose any endianness for hexadecimal encoding.
The individual symbols are returned separately. For example:

```go
b1, b0 := hexadeca.EncodeByteUsingUpperCaseSymbols(u8)
```

```go
b3, b2, b1, b0 := hexadeca.EncodeUint16UsingLowerCaseSymbols(u16)
```

```go
r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint32UsingPersianSymbols(u32)
```

Package **hexadeca** has functions for hexadecimal encoding and decoding for the Go types:
* `byte` (i.e., `uint8`),
* `uint16`,
* `uint32`
* `uint64`
* `int8`
* `int16`
* `int32`
* `int64`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-hexadeca

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-hexadeca?status.svg)](https://godoc.org/sourcecode.social/reiver/go-hexadeca)

## Symbols

Package **hexadeca** lets you pick the 16 symbols to hexadecimal encode to and hexadecimal decode from.

Package **hexadeca** has built-in support for the following symbol sequences:

| Name       | zero  | one   | two   | three | four | five   | six   | seven | eight | nine  | ten   | eleven | twelve | thirteen | fourteen | fifteen |
|------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|--------|--------|----------|----------|---------|
| Lower-Case | `'0'` | `'1'` | `'2'` | `'3'` | `'4'` | `'5'` | `'6'` | `'7'` | `'8'` | `'9'` | `'a'` | `'b'`  | `'c'`  | `'d'`    | `'e'`    | `'f'`   |
| Upper-Case | `'0'` | `'1'` | `'2'` | `'3'` | `'4'` | `'5'` | `'6'` | `'7'` | `'8'` | `'9'` | `'A'` | `'B'`  | `'C'`  | `'D'`    | `'E'`    | `'F'`   |
| Persian    | `'۰'` | `'۱'` | `'۲'` | `'۳'` | `'۴'` | `'۵'` | `'۶'` | `'۷'` | `'۸'` | `'۹'` | `'ی'` | `'ک'`  | `'ل'`  | `'م'`    | `'ن'`    | `'س'`   |

Package **hexadeca** also lets you pick your own symbol sequences, using a function.

For example, one could create a symbol sequence function to make the hexadecimal encoding be:

| Name             | zero  | one   | two   | three | four | five   | six   | seven | eight | nine  | ten   | eleven | twelve | thirteen | fourteen | fifteen |
|------------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|--------|--------|----------|----------|---------|
| _Custom Example_ | `𝍠`   | `𝍡`   | `𝍢`   | `𝍣`   | `𝍤`   | `𝍥`   | `𝍦`   | `𝍧`   | `𝍨`   | `𝍩`   | `𝍪`   | `𝍫`    | `𝍬`    | `𝍭`      | `𝍮`      | `𝍯`     |
| _Custom Example_ | `❤️`   | `💜`  | `💙`  | `💛`  | `💚`  | `🤍`  | `🖤`  | `🧡`  | `🤎`  | `💕`  | `💔`  | `💖`   | `💗`   | `💓`     | `💞`     | `💝`    |

... or anything else.

### Samples

Here are some samples of hexadecimal numbers using each of the built-in symbol sequences:

| Name       | Decimal 49,374 | Decimal 19,229 | Decimal 255 |
|------------|----------------|----------------|-------------|
| Lower-Case | `c0de`         | `4b1d`         | `ff`        |
| Upper-Case | `C0DE`         | `4B1D`         | `FF`        |
| Persian    | `نم۰ل`         | `م۱ک۴`         | `'سس'`      |

### Custom Symbols

Package **hexadeca** also lets you use your own custom symbols by using your own symbol function.

For example:
```go
func CustomSymbol(value byte) rune {
        var table [16]rune = [16]rune{
                '𝍠', //  0
                '𝍡', //  1
                '𝍢', //  2
                '𝍣', //  3
                '𝍤', //  4
                '𝍥', //  5
                '𝍦', //  6
                '𝍧', //  7
                '𝍨', //  8
                '𝍩', //  9
                '𝍪', // 10
                '𝍫', // 11
                '𝍬', // 12
                '𝍭', // 13
                '𝍮', // 14
                '𝍯', // 15
        }

        var index int = int(value) % len(table)

        var result rune = table[index]

        return result
}

// ...

r1, r0 := hexadeca.EncodeRune(value, CustomSymbol)
```

Samples:  
* `𝍬𝍠𝍭𝍮`
* `𝍤𝍫𝍡𝍭` 

## Import

To import package **hexadeca** use `import` code like the following:
```
import "sourcecode.social/reiver/go-hexadeca"
```

## Installation

To install package **hexadeca** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-hexadeca
```

## Author

Package **hexadeca** was written by [Charles Iliya Krempeaux](http://changelog.ca)
