package hexadecarunes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/rune"
)

func EncodeUint64(value uint64, symbolFunc func(byte)rune) (r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	r15, r14 = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 7))), symbolFunc)
	r13, r12 = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 6))), symbolFunc)
	r11, r10 = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 5))), symbolFunc)
	r9,  r8  = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 4))), symbolFunc)
	r7,  r6  = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 3))), symbolFunc)
	r5,  r4  = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 2))), symbolFunc)
	r3,  r2  = EncodeByte(byte(0x00000000000000ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00000000000000ff &  value            ), symbolFunc)

	return
}

func EncodeUint64UsingLowerCaseSymbols(value uint64) (r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint64(value, runesymbols.LowerCase)
}

func EncodeUint64UsingPersianSymbols(value uint64) (r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint64(value, runesymbols.Persian)
}

func EncodeUint64UsingUpperCaseSymbols(value uint64) (r15 rune, r14 rune, r13 rune, r12 rune, r11 rune, r10 rune, r9 rune, r8 rune, r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint64(value, runesymbols.UpperCase)
}
