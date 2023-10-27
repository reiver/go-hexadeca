
# go-hexadeca

Package **hexadeca** implements **hexadecimal** _encoding_ and _decoding_, for the Go programming language.
This is meant to be a better alternative to Go's built-in `"hex"` package.

Package **hexadeca** has functions for hexadecimal encoding and decoding with Go types:
* `byte` (i.e., `uint8`),
* `uint16`,
* `uint32`
* `uint64`
* `int64`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-hexadeca

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-hexadeca?status.svg)](https://godoc.org/sourcecode.social/reiver/go-hexadeca)

## Symbols

Package **hexadeca** lets you pick the 16 symbols to hexadecimal encode to and hexadecimal decode from.
For example —

### Lower-Case Symbols

`'0'`, `'1'`, `'2'`, `'3'`, `'4'`, `'5'`, `'6'`, `'7'`, `'8'`, `'9'`, `'a'`, `'b'`, `'c'`, `'d'`, `'e'`, `'f'`.

Sample:  
`c0de`

### Upper-Case Symbols

`'0'`, `'1'`, `'2'`, `'3'`, `'4'`, `'5'`, `'6'`, `'7'`, `'8'`, `'9'`, `'A'`, `'B'`, `'C'`, `'D'`, `'E'`, `'F'`.

Sample:  
`C0DE`

### Persian Symbols

`'۰'` [zero], `'۱'` [one], `'۲'` [two], `'۳'` [three], `'۴'` [four], `'۵'` [five], `'۶'` [six], `'۷'` [seven], `'۸'` [eight], `'۹'` [nine], `'ی'` (ye) [ten], `'ک'` (kāf) [eleven], `'ل'` (lâm) [twelve], `'م'` (mim) [thirteen], `'ن'` (nun) [fourteen], `'س'`  (sin) [fifteen]

Sample:  
`ل۰من`

### Custom Symbols

As well as defining you own using function:
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

Sample:  
`𝍬𝍠𝍭𝍮`

## hexadecimal encode `byte`

To hexadecimal encode a `byte` use any of these:

```go
import "sourcecode.social/reiver/go-hexadeca/enc/bytes"

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadecabytes.EncodeByteUsingLowerCaseSymbols(value)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadecabytes.EncodeByteUsingUpperCaseSymbols(value)

```

Or:

```go
import "sourcecode.social/reiver/go-hexadeca/enc/runes"

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadecarunes.EncodeByteUsingLowerCaseSymbols(value)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadecarunes.EncodeByteUsingUpperCaseSymbols(value)

```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca/enc/bytes"
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadecabytes.EncodeByte(value, bytesymbols.LowerCase)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadecabytes.EncodeByte(value, bytesymbols.UpperCase)
```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca/enc/runes"
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// ...

// value==254 -> mostSignificant=='f', leastSignificant=='e'
mostSignificant, leastSignificant := hexadecarunes.EncodeByte(value, runesymbols.LowerCase)

// ...

// value==254 -> mostSignificant=='F', leastSignificant=='E'
mostSignificant, leastSignificant := hexadecarunes.EncodeByte(value, runesymbols.UpperCase)
```

## hexadecimal encode `uint16`

To hexadecimal encode a `uint16` use any of these:

```go
import "sourcecode.social/reiver/go-hexadeca/enc/bytes"

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16UsingLowerCaseSymbols(value)

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16UsingUpperCaseSymbols(value)

```

Or:

```go
import "sourcecode.social/reiver/go-hexadeca/enc/runes"

// ...

r3, r2, r1, r0 := hexadecarunes.EncodeUint16UsingLowerCaseSymbols(value)

// ...

r3, r2, r1, r0 := hexadecarunes.EncodeUint16UsingUpperCaseSymbols(value)

```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca/enc/bytes"
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16(value, bytesymbols.LowerCase)

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16(value, bytesymbols.UpperCase)
```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca/enc/bytes"
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16(value, runesymbols.LowerCase)

// ...

r3, r2, r1, r0 := hexadecabytes.EncodeUint16(value, runesymbols.UpperCase)
```

## hexadecimal encode `uint32`

To hexadecimal encode a `uint32` use any of these:

```go
import "sourcecode.social/reiver/go-hexadeca"

// ...

r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint32UsingLowerCaseSymbols(value)

// ...

r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint32UsingUpperCaseSymbols(value)

```

Or:

```go
import  (
	"sourcecode.social/reiver/go-hexadeca"
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// ...

r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint32(value, runesymbols.LowerCase)

// ...

r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint32(value, runesymbols.UpperCase)
```

## hexadecimal encode `uint64`

To hexadecimal encode a `uint64` use any of these:

```go
import "sourcecode.social/reiver/go-hexadeca"

// ...

// value==18364758544493064720 -> r15=='f', r14=='e', r13=='d', r12=='c', r11=='b', r10=='a', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint64UsingLowerCaseSymbols(value)

// ...

// value==18364758544493064720 -> r15=='F', r14=='E', r13=='D', r12=='C', r11=='B', r10=='A', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint64UsingUpperCaseSymbols(value)

```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca"
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// ...

// value==18364758544493064720 -> r15=='f', r14=='e', r13=='d', r12=='c', r11=='b', r10=='a', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint64(value, runesymbols.LowerCase)

// ...

// value==18364758544493064720 -> r15=='F', r14=='E', r13=='D', r12=='C', r11=='B', r10=='A', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeUint64(value, runesymbols.UpperCase)
```

## hexadecimal encode `int64`

To hexadecimal encode a `int64` use any of these:

```go
import "sourcecode.social/reiver/go-hexadeca"

// ...

// value==18364758544493064720 -> r15=='f', r14=='e', r13=='d', r12=='c', r11=='b', r10=='a', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeInt64UsingLowerCaseSymbols(value)

// ...

// value==18364758544493064720 -> r15=='F', r14=='E', r13=='D', r12=='C', r11=='B', r10=='A', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeInt64UsingUpperCaseSymbols(value)

```

Or:

```go
import (
	"sourcecode.social/reiver/go-hexadeca"
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// ...

// value==18364758544493064720 -> r15=='f', r14=='e', r13=='d', r12=='c', r11=='b', r10=='a', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeInt64(value, runesymbols.LowerCase)

// ...

// value==18364758544493064720 -> r15=='F', r14=='E', r13=='D', r12=='C', r11=='B', r10=='A', r9=='9', r8=='8', r7=='7', r6=='6', r5=='5', r4=='4', r3=='3', r2=='2', r1=='1', r0=='0'
r15, r14, r13, r12, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0 := hexadeca.EncodeInt64(value, runesymbols.UpperCase)
```

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
