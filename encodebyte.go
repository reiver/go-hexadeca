package hexadeca

// EncodeByte encodes the value of a byte into hexadecimal.
func EncodeByte(value byte, symbolFunc func(byte)rune) (mostSignificant rune, leastSignificant rune) {

	mostSignificant  = symbolFunc( (0xf0 & value) >> 4 )
	leastSignificant = symbolFunc( (0x0f & value)      )

	return
}

func EncodeByteUsingLowerCaseSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, SymbolLowerCase)
}

func EncodeByteUsingPersianSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, SymbolLowerCase)
}

func EncodeByteUsingUpperCaseSymbols(value byte) (mostSignificant rune, leastSignificant rune) {
	return EncodeByte(value, SymbolUpperCase)
}
