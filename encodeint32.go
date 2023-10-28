package hexadeca

import (
	"sourcecode.social/reiver/go-hexadeca/bytes"
	"sourcecode.social/reiver/go-hexadeca/runes"
)

func EncodeInt32UsingLowerCaseSymbols(value int32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeInt32UsingLowerCaseSymbols(value)
}

func EncodeInt32UsingPersianSymbols(value int32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	return hexadecarunes.EncodeInt32UsingPersianSymbols(value)
}

func EncodeInt32UsingUpperCaseSymbols(value int32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeInt32UsingUpperCaseSymbols(value)
}
