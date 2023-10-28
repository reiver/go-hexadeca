package hexadeca

import (
	"sourcecode.social/reiver/go-hexadeca/bytes"
)

func EncodeByte(value byte, symbolFunc func(byte)byte) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeByte(value, symbolFunc)
}

func EncodeByteUsingLowerCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeByteUsingLowerCaseSymbols(value)
}

func EncodeByteUsingUpperCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeByteUsingUpperCaseSymbols(value)
}
