package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

func EncodeInt32(value int32, symbolFunc func(byte)byte) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
	r7,  r6  = EncodeByte(byte(0x000000ff & (value >> (8 * 3))), symbolFunc)
	r5,  r4  = EncodeByte(byte(0x000000ff & (value >> (8 * 2))), symbolFunc)
	r3,  r2  = EncodeByte(byte(0x000000ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x000000ff &  value            ), symbolFunc)

	return
}

func EncodeInt32UsingLowerCaseSymbols(value int32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt32(value, bytesymbols.LowerCase)
}

func EncodeInt32UsingUpperCaseSymbols(value int32) (r7 byte, r6 byte, r5 byte, r4 byte, r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt32(value, bytesymbols.UpperCase)
}
