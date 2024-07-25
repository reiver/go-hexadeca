package runesymbols

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

// UpperCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using upper-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, runesymbols.UpperCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'F' and and ‘leastSignificant’ would be 'E'.
func UpperCase(value byte) rune {
	return rune(bytesymbols.UpperCase(value))
}
