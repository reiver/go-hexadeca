package hexadecarunes

import (
	"github.com/reiver/go-hexadeca/symbols/rune"
)

func EncodeInt16(value int16, symbolFunc func(byte)rune) (r3 rune, r2 rune, r1 rune, r0 rune) {
	r3,  r2  = EncodeByte(byte(0x00ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00ff &  value            ), symbolFunc)

	return
}

func EncodeInt16UsingLowerCaseSymbols(value int16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt16(value, runesymbols.LowerCase)
}

func EncodeInt16UsingPersianSymbols(value int16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt16(value, runesymbols.Persian)
}

func EncodeInt16UsingUpperCaseSymbols(value int16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeInt16(value, runesymbols.UpperCase)
}
