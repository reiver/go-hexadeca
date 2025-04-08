package hexadecabytes

import (
	"github.com/reiver/go-hexadeca/symbols/byte"
)

func DecodeUint64(symbolFunc func(byte)(byte,bool), s15 byte, s14 byte, s13 byte, s12 byte, s11 byte, s10 byte, s9 byte, s8 byte, s7 byte, s6 byte, s5 byte, s4 byte, s3 byte, s2 byte, s1 byte, s0 byte) (uint64, error) {

	b7, err := DecodeByte(symbolFunc, s15, s14)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b6, err := DecodeByte(symbolFunc, s13, s12)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b5, err := DecodeByte(symbolFunc, s11, s10)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b4, err := DecodeByte(symbolFunc,  s9,  s8)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b3, err := DecodeByte(symbolFunc,  s7,  s6)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b2, err := DecodeByte(symbolFunc,  s5,  s4)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b1, err := DecodeByte(symbolFunc,  s3,  s2)
	if nil != err {
		var nada uint64
		return nada, err
	}

	b0, err := DecodeByte(symbolFunc,  s1,  s0)
	if nil != err {
		var nada uint64
		return nada, err
	}

	return	(uint64(b7)  << (8* 7)) |
		(uint64(b6)  << (8* 6)) |
		(uint64(b5)  << (8* 5)) |
		(uint64(b4)  << (8* 4)) |
		(uint64(b3)  << (8* 3)) |
		(uint64(b2)  << (8* 2)) |
		(uint64(b1)  << (8   )) |
		(uint64(b0)           ), nil
}


func DecodeUint64UsingLowerCaseSymbols(s15 byte, s14 byte, s13 byte, s12 byte, s11 byte, s10 byte, s9 byte, s8 byte, s7 byte, s6 byte, s5 byte, s4 byte, s3 byte, s2 byte, s1 byte, s0 byte) (uint64, error) {
        return DecodeUint64(bytesymbols.FromLowerCase, s15, s14, s13, s12, s11, s10, s9, s8, s7, s6, s5, s4, s3, s2, s1, s0)
}

func DecodeUint64UsingUpperCaseSymbols(s15 byte, s14 byte, s13 byte, s12 byte, s11 byte, s10 byte, s9 byte, s8 byte, s7 byte, s6 byte, s5 byte, s4 byte, s3 byte, s2 byte, s1 byte, s0 byte) (uint64, error) {
        return DecodeUint64(bytesymbols.FromUpperCase, s15, s14, s13, s12, s11, s10, s9, s8, s7, s6, s5, s4, s3, s2, s1, s0)
}
