package hexadeca

import (
	"github.com/reiver/go-hexadeca/bytes"
	"github.com/reiver/go-hexadeca/runes"
)

func EncodeUint16UsingLowerCaseSymbols(value uint16) (r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint16UsingLowerCaseSymbols(value)
}

func EncodeUint16UsingPersianSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
	return hexadecarunes.EncodeUint16UsingPersianSymbols(value)
}

func EncodeUint16UsingUpperCaseSymbols(value uint16) (r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint16UsingUpperCaseSymbols(value)
}
