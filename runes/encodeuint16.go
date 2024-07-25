package hexadecarunes

import (
	"github.com/reiver/go-hexadeca/symbols/rune"
)

func EncodeUint16(value uint16, symbolFunc func(byte)rune) (r3 rune, r2 rune, r1 rune, r0 rune) {
	r3,  r2  = EncodeByte(byte(0x00ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00ff &  value            ), symbolFunc)

	return
}

func EncodeUint16UsingLowerCaseSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, runesymbols.LowerCase)
}

func EncodeUint16UsingPersianSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, runesymbols.Persian)
}

func EncodeUint16UsingUpperCaseSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, runesymbols.UpperCase)
}
