package hexadecabytes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

// EncodeByte encodes the value of a byte into hexadecimal.
func EncodeByte(value byte, symbolFunc func(byte)byte) (mostSignificant byte, leastSignificant byte) {

	mostSignificant  = symbolFunc( (0x0f & (value >> 4)) )
	leastSignificant = symbolFunc( (0x0f & (value     )) )

	return
}

func EncodeByteUsingLowerCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return EncodeByte(value, bytesymbols.LowerCase)
}

func EncodeByteUsingUpperCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return EncodeByte(value, bytesymbols.UpperCase)
}
