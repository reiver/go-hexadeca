package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"

	"github.com/reiver/go-erorr"
)

// DecodeByte decodes the value of a byte from hexadecimal.
func DecodeByte(symbolFunc func(byte)(byte,bool), mostSignificant byte, leastSignificant byte) (byte, error) {

	n1, ok1 := symbolFunc(mostSignificant)
	if !ok1 {
		var nada byte
		return nada, erorr.Errorf("hexadeca: cannot decode most-significant hexadecimal symbol (%q)", mostSignificant)
	}

	n0, ok0 := symbolFunc(leastSignificant)
	if !ok0 {
		var nada byte
		return nada, erorr.Errorf("hexadeca: cannot decode most-significant hexadecimal symbol (%q)", leastSignificant)
	}

	return (n1 << 4) | (n0), nil
}

func DecodeByteUsingLowerCaseSymbols(mostSignificant byte, leastSignificant byte) (byte, error) {
	return DecodeByte(bytesymbols.FromLowerCase, mostSignificant, leastSignificant)
}

func DecodeByteUsingUpperCaseSymbols(mostSignificant byte, leastSignificant byte) (byte, error) {
	return DecodeByte(bytesymbols.FromUpperCase, mostSignificant, leastSignificant)
}
