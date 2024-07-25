package hexadeca

import (
	"github.com/reiver/go-hexadeca/bytes"
	"github.com/reiver/go-hexadeca/runes"
)

func EncodeInt8UsingLowerCaseSymbols(value int8) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeInt8UsingLowerCaseSymbols(value)
}

func EncodeInt8UsingPersianSymbols(value int8) (mostSignificant rune, leastSignificant rune) {
	return hexadecarunes.EncodeInt8UsingPersianSymbols(value)
}

func EncodeInt8UsingUpperCaseSymbols(value int8) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeInt8UsingUpperCaseSymbols(value)
}
