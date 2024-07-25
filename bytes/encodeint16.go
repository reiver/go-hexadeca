package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

func EncodeInt16(value int16, symbolFunc func(byte)byte) (r3 byte, r2 byte, r1 byte, r0 byte) {
	r3,  r2  = EncodeByte(byte(0x00ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00ff &  value            ), symbolFunc)

	return
}

func EncodeInt16UsingLowerCaseSymbols(value int16) (r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt16(value, bytesymbols.LowerCase)
}

func EncodeInt16UsingUpperCaseSymbols(value int16) (r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeInt16(value, bytesymbols.UpperCase)
}
