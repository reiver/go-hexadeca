package hexadecabytes

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

func EncodeUint64(value uint64, symbolFunc func(byte)byte) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
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

func EncodeUint64UsingLowerCaseSymbols(value uint64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeUint64(value, bytesymbols.LowerCase)
}

func EncodeUint64UsingUpperCaseSymbols(value uint64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeUint64(value, bytesymbols.UpperCase)
}
