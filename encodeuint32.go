package hexadeca

import (
	"github.com/reiver/go-hexadeca/bytes"
	"github.com/reiver/go-hexadeca/runes"
)

func EncodeUint32UsingLowerCaseSymbols(value uint32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint32UsingLowerCaseSymbols(value)
}

func EncodeUint32UsingPersianSymbols(value uint32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	return hexadecarunes.EncodeUint32UsingPersianSymbols(value)
}

func EncodeUint32UsingUpperCaseSymbols(value uint32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint32UsingUpperCaseSymbols(value)
}
