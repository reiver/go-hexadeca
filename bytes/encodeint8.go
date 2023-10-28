package hexadecabytes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

// EncodeInt8 encodes the value of a int8 into hexadecimal.
func EncodeInt8(value int8, symbolFunc func(byte)byte) (mostSignificant byte, leastSignificant byte) {

	mostSignificant  = symbolFunc( byte(0x0f & (value >> 4)) )
	leastSignificant = symbolFunc( byte(0x0f & (value     )) )

	return
}

func EncodeInt8UsingLowerCaseSymbols(value int8) (mostSignificant byte, leastSignificant byte) {
	return EncodeInt8(value, bytesymbols.LowerCase)
}

func EncodeInt8UsingUpperCaseSymbols(value int8) (mostSignificant byte, leastSignificant byte) {
	return EncodeInt8(value, bytesymbols.UpperCase)
}
