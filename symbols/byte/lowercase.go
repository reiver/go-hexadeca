package bytesymbols

// LowerCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using lower-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, bytesymbols.LowerCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'f' and and ‘leastSignificant’ would be 'e'.
func LowerCase(value byte) byte {
	const table string = "0123456789abcdef"

	var index int = int(value) % len(table)

	var result byte = table[index]

	return result
}
