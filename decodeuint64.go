package hexadeca

import (
	"github.com/reiver/go-hexadeca/bytes"
//	"github.com/reiver/go-hexadeca/runes"
)

func DecodeUint64UsingLowerCaseSymbols(b15 byte, b14 byte, b13 byte, b12 byte, b11 byte, b10 byte, b9 byte, b8 byte, b7 byte, b6 byte, b5 byte, b4 byte, b3 byte, b2 byte, b1 byte, b0 byte) (uint64, error) {
	return hexadecabytes.DecodeUint64UsingLowerCaseSymbols(b15, b14, b13, b12, b11, b10, b9, b8, b7, b6, b5, b4, b3, b2, b1, b0)
}

//func DecodeUint64UsingPersianSymbols(r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) (uint64, error) {
//	return hexadecarunes.DecodeUint64UsingPersianSymbols(r15, r14, r13, r12, r11, r10, r9, r8, r7, r6, r5, r4, r3, r2, r1, r0)
//}

func DecodeUint64UsingUpperCaseSymbols(b15 byte, b14 byte, b13 byte, b12 byte, b11 byte, b10 byte, b9 byte, b8 byte, b7 byte, b6 byte, b5 byte, b4 byte, b3 byte, b2 byte, b1 byte, b0 byte) (uint64, error) {
	return hexadecabytes.DecodeUint64UsingUpperCaseSymbols(b15, b14, b13, b12, b11, b10, b9, b8, b7, b6, b5, b4, b3, b2, b1, b0)
}
