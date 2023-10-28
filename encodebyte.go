package hexadeca

import (
	"sourcecode.social/reiver/go-hexadeca/bytes"
	"sourcecode.social/reiver/go-hexadeca/runes"
)

func EncodeByteUsingLowerCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeByteUsingLowerCaseSymbols(value)
}

func EncodeByteUsingPersianSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return hexadecarunes.EncodeByteUsingPersianSymbols(value)
}

func EncodeByteUsingUpperCaseSymbols(value byte) (mostSignificant byte, leastSignificant byte) {
	return hexadecabytes.EncodeByteUsingUpperCaseSymbols(value)
}
