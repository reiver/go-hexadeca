package runesymbols

// Persian is used with hexadeca.EncodeByte() to encode a byte into hexadecimal symbols using Persian symbols from Unicode UTF-8.
// I.e.,
//
// '۰' [zero], '۱' [one], '۲' [two], '۳' [three], '۴' [four], '۵' [five], '۶' [six], '۷' [seven], '۸' [eight], '۹' [nine], 'ی' (ye) [ten], 'ک' (kāf) [eleven], 'ل' (lâm) [twelve], 'م' (mim) [thirteen], 'ن' (nun) [fourteen], 'س'  (sin) [fifteen]
//
// Example usage:
//
//	mostSignificiant, leastSignificant := hexadeca.EncodeByte(value, runesymbols.Persian)
//
// If ‘value’ was 90 (== 0x5A) then ‘mostSignificiant’ would be '۵' and and ‘leastSignificant’ would be '\u0627'.
func Persian(value byte) rune {
	var table [16]rune = [16]rune{
		'\u06F0', // ۰ [zero]
		'\u06F1', // ۱ [one]
		'\u06F2', // ۲ [two]
		'\u06F3', // ۳ [three]
		'\u06F4', // ۴ [four]
		'\u06F5', // ۵ [five]
		'\u06F6', // ۶ [six]
		'\u06F7', // ۷ [seven]
		'\u06F8', // ۸ [eight]
		'\u06F9', // ۹ [nine]
		'\u06CC', // ی (ye) [ten]
		'\u06A9', // ک (kāf) [eleven]
		'\u0644', // ل (lâm) [twelve]
		'\u0645', // م (mim) [thirteen]
		'\u0646', // ن (nun) [fourteen]
		'\u0633', // س  (sin) [fifteen]
	}

	var index int = int(value) % len(table)

	var result rune = table[index]

	return result
}
