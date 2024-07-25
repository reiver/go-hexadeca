package hexadecarunes

import (
	"github.com/reiver/go-hexadeca/symbols/rune"
)

func EncodeInt32(value int32, symbolFunc func(byte)rune) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	r7,  r6  = EncodeByte(byte(0x000000ff & (value >> (8 * 3))), symbolFunc)
	r5,  r4  = EncodeByte(byte(0x000000ff & (value >> (8 * 2))), symbolFunc)
	r3,  r2  = EncodeByte(byte(0x000000ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x000000ff &  value            ), symbolFunc)

	return
}

func EncodeInt32UsingLowerCaseSymbols(value int32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt32(value, runesymbols.LowerCase)
}

func EncodeInt32UsingPersianSymbols(value int32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt32(value, runesymbols.Persian)
}

func EncodeInt32UsingUpperCaseSymbols(value int32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt32(value, runesymbols.UpperCase)
}
