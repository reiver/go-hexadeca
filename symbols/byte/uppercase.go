package bytesymbols

// UpperCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using upper-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, bytesymbols.UpperCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'f' and and ‘leastSignificant’ would be 'e'.
func UpperCase(value byte) byte {
	const table string = "0123456789ABCDEF"

	var index int = int(value) % len(table)

	var result byte = table[index]

	return result
}
