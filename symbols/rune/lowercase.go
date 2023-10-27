package runesymbols

import (
	"sourcecode.social/reiver/go-hexadeca/symbols/byte"
)

// LowerCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using lower-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, runesymbols.LowerCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'f' and and ‘leastSignificant’ would be 'e'.
func LowerCase(value byte) rune {
	return rune(bytesymbols.LowerCase(value))
}
