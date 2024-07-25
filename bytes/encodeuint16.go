package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

func EncodeUint16(value uint16, symbolFunc func(byte)byte) (r3 byte, r2 byte, r1 byte, r0 byte) {
	r3,  r2  = EncodeByte(byte(0x00ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00ff &  value            ), symbolFunc)

	return
}

func EncodeUint16UsingLowerCaseSymbols(value uint16) (r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeUint16(value, bytesymbols.LowerCase)
}

func EncodeUint16UsingUpperCaseSymbols(value uint16) (r3 byte, r2 byte, r1 byte, r0 byte) {
        return EncodeUint16(value, bytesymbols.UpperCase)
}
