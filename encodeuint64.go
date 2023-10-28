package hexadeca

import (
	"sourcecode.social/reiver/go-hexadeca/bytes"
	"sourcecode.social/reiver/go-hexadeca/runes"
)

func EncodeUint64UsingLowerCaseSymbols(value uint64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint64UsingLowerCaseSymbols(value)
}

func EncodeUint64UsingPersianSymbols(value uint64) (r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	return hexadecarunes.EncodeUint64UsingPersianSymbols(value)
}

func EncodeUint64UsingUpperCaseSymbols(value uint64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	return hexadecabytes.EncodeUint64UsingUpperCaseSymbols(value)
}
