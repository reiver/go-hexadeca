package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

func EncodeInt64(value int64, symbolFunc func(byte)byte) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
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

func EncodeInt64UsingLowerCaseSymbols(value int64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt64(value, bytesymbols.LowerCase)
}

func EncodeInt64UsingUpperCaseSymbols(value int64) (r15 byte, r14 byte, r13 byte, r12 byte, r11 byte, r10 byte, r9 byte, r8 byte, r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt64(value, bytesymbols.UpperCase)
}
