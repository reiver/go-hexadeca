package hexadeca

func EncodeUint16(value uint16, symbolFunc func(byte)rune) (r3 rune, r2 rune, r1 rune, r0 rune) {
	r3,  r2  = EncodeByte(byte(0x00ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00ff &  value            ), symbolFunc)

	return
}

func EncodeUint16UsingLowerCaseSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, SymbolLowerCase)
}

func EncodeUint16UsingPersianSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, SymbolPersian)
}

func EncodeUint16UsingUpperCaseSymbols(value uint16) (r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint16(value, SymbolUpperCase)
}
