package hexadeca

// SymbolLowerCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using lower-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, hexadeca.SymbolLowerCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'f' and and ‘leastSignificant’ would be 'e'.
func SymbolLowerCase(value byte) rune {
	const table string = "0123456789abcdef"

	var index int = int(value) % len(table)

	var result rune = rune(table[index])

	return result
}

// SymbolPersian is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using Persian symbols from Unicode UTF-8.
// I.e.,
//
// '۰' (zero), '۱' (one), '۲' (two), '۳' (three), '۴' (four), '۵' (five), '۶' (six), '۷' (seven), '۸' (eight) ,'۹' (nine), 'ا' (alef), 'ب ' (be), 'پ ' (pe), 'ت ' (te), 'ث ' (s̱e), 'ج' (jim)
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, hexadeca.SymbolPersian)
//
// If ‘value’ was 90 (== 0x5A) then ‘mostSignificiant’ would be '۵' and and ‘leastSignificant’ would be '\u0627'.
func SymbolPersian(value byte) rune {
	var table [16]rune = [16]rune{
		'\u06F0', // ۰ (zero)
		'\u06F1', // ۱ (one)
		'\u06F2', // ۲ (two)
		'\u06F3', // ۳ (three)
		'\u06F4', // ۴ (four)
		'\u06F5', // ۵ (five)
		'\u06F6', // ۶ (six)
		'\u06F7', // ۷ (seven)
		'\u06F8', // ۸ (eight)
		'\u06F9', // ۹ (nine)
		'\u0627', // ا (alef)
		'\u0628', // ب  (be)
		'\u067E', // پ  (pe)
		'\u062A', // ت  (te)
		'\u062B', // ث  (s̱e)
		'\u062C', // ج (jim)
	}

	var index int = int(value) % len(table)

	var result rune = table[index]

	return result
}

// SymbolUpperCase is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using upper-case symbols from ASCII / Unicode UTF-8.
// I.e.,
//
// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, hexadeca.SymbolUpperCase)
//
// If ‘value’ was 254 (== 0xFE) then ‘mostSignificiant’ would be 'F' and and ‘leastSignificant’ would be 'E'.
func SymbolUpperCase(value byte) rune {
	const table string = "0123456789ABCDEF"

	var index int = int(value) % len(table)

	var result rune = rune(table[index])

	return result
}