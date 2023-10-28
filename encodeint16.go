package hexadeca

import (
	"sourcecode.social/reiver/go-hexadeca/bytes"
	"sourcecode.social/reiver/go-hexadeca/runes"
)

func EncodeInt16UsingLowerCaseSymbols(value int16) (r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeInt16UsingLowerCaseSymbols(value)
}

func EncodeInt16UsingPersianSymbols(value int16) (r3 rune, r2 rune, r1 rune, r0 rune) {
	return hexadecarunes.EncodeInt16UsingPersianSymbols(value)
}

func EncodeInt16UsingUpperCaseSymbols(value int16) (r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeInt16UsingUpperCaseSymbols(value)
}
