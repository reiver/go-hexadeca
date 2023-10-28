package hexadecarunes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

// EncodeByte encodes the value of a byte into hexadecimal.
func EncodeByte(value byte, symbolFunc func(byte)rune) (mostSignificant rune, leastSignificant rune) {

	mostSignificant  = symbolFunc( (0x0f & (value >> 4)) )
	leastSignificant = symbolFunc( (0x0f & (value     )) )

	return
}

func EncodeByteUsingLowerCaseSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, runesymbols.LowerCase)
}

func EncodeByteUsingPersianSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, runesymbols.Persian)
}

func EncodeByteUsingUpperCaseSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, runesymbols.UpperCase)
}
