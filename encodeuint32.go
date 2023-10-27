package hexadeca

func EncodeUint32(value uint32, symbolFunc func(byte)rune) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
	r7,  r6  = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 3))), symbolFunc)
	r5,  r4  = EncodeByte(byte(0x00000000000000ff & (value >> (8 * 2))), symbolFunc)
	r3,  r2  = EncodeByte(byte(0x00000000000000ff & (value >> (8    ))), symbolFunc)
	r1,  r0  = EncodeByte(byte(0x00000000000000ff &  value            ), symbolFunc)

	return
}

func EncodeUint32UsingLowerCaseSymbols(value uint32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint32(value, SymbolLowerCase)
}

func EncodeUint32UsingPersianSymbols(value uint32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint32(value, SymbolPersian)
}

func EncodeUint32UsingUpperCaseSymbols(value uint32) (r7 rune, r6 rune, r5 rune, r4 rune, r3 rune, r2 rune, r1 rune, r0 rune) {
        return EncodeUint32(value, SymbolUpperCase)
}
