package hexadecarunes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// EncodeInt8 encodes the value of a int8 into hexadecimal.
func EncodeInt8(value int8, symbolFunc func(byte)rune) (mostSignificant rune, leastSignificant rune) {

	mostSignificant  = symbolFunc( byte(0x0f & (value >> 4)) )
	leastSignificant = symbolFunc( byte(0x0f & (value     )) )

	return
}

func EncodeInt8UsingLowerCaseSymbols(value int8) (mostSignificant rune, leastSignificant rune) {
	return EncodeInt8(value, runesymbols.LowerCase)
}

func EncodeInt8UsingPersianSymbols(value int8) (mostSignificant rune, leastSignificant rune) {
	return EncodeInt8(value, runesymbols.Persian)
}

func EncodeInt8UsingUpperCaseSymbols(value int8) (mostSignificant rune, leastSignificant rune) {
	return EncodeInt8(value, runesymbols.UpperCase)
}
