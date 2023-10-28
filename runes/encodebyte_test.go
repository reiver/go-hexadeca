package hexadecarunes_test

import (
	"testing"

	"sourcecode.social/reiver/go-hexadeca/runes"
)

func TestEncodeByteUsingLowerCaseSymbols(t *testing.T) {

	tests := []struct{
		Value byte
		ExpectedMostSignificant rune
		ExpectedLeastSignificant rune
	}{
		{
			Value:                  0x00,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x01,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x02,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x03,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x04,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x05,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x06,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x07,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x08,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x09,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x0A,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x0B,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x0C,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x0D,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x0E,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x0F,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x10,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x11,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x12,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x13,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x14,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x15,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x16,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x17,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x18,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x19,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x1A,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x1B,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x1C,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x1D,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x1E,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x1F,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'f',
		},



		{
			Value:                  0x47,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '7',
		},



		{
			Value:                  0x5C,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'c',
		},



		{
			Value:                  0xD0,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xD1,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xD2,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xD3,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xD4,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xD5,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xD6,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xD7,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xD8,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xD9,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xDA,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xDB,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xDC,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xDD,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xDE,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xDF,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'f',
		},



		{
			Value:                  0xF9,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xFA,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xFB,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xFC,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xFD,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xFE,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xFF,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'f',
		},
	}

	for testNumber, test := range tests {

		actualMostSignificant, actualLeastSignificant := hexadecarunes.EncodeByteUsingLowerCaseSymbols(test.Value)

		{
			actual := actualMostSignificant
			expected := test.ExpectedMostSignificant

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the most-significant symbol is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:     %X (%d)", test.Value, test.Value)
				continue
			}
		}

		{
			actual := actualLeastSignificant
			expected := test.ExpectedLeastSignificant

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the least-significant symbol is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    %X (%d)", test.Value, test.Value)
				continue
			}
		}
	}
}

func TestEncodeByteUsingUpperCaseSymbols(t *testing.T) {

	tests := []struct{
		Value byte
		ExpectedMostSignificant rune
		ExpectedLeastSignificant rune
	}{
		{
			Value:                  0x00,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x01,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x02,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x03,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x04,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x05,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x06,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x07,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x08,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x09,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x0A,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x0B,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x0C,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x0D,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x0E,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x0F,
			ExpectedMostSignificant: '0',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x10,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x11,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x12,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x13,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x14,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x15,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x16,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x17,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x18,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x19,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x1A,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x1B,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x1C,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x1D,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x1E,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x1F,
			ExpectedMostSignificant: '1',
			ExpectedLeastSignificant: 'F',
		},



		{
			Value:                  0x47,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '7',
		},



		{
			Value:                  0x5C,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'C',
		},



		{
			Value:                  0xD0,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xD1,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xD2,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xD3,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xD4,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xD5,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xD6,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xD7,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xD8,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xD9,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xDA,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xDB,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xDC,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xDD,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xDE,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xDF,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'F',
		},



		{
			Value:                  0xF9,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xFA,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xFB,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xFC,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xFD,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xFE,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xFF,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'F',
		},
	}

	for testNumber, test := range tests {

		actualMostSignificant, actualLeastSignificant := hexadecarunes.EncodeByteUsingUpperCaseSymbols(test.Value)

		{
			actual := actualMostSignificant
			expected := test.ExpectedMostSignificant

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the most-significant symbol is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:     %X (%d)", test.Value, test.Value)
				continue
			}
		}

		{
			actual := actualLeastSignificant
			expected := test.ExpectedLeastSignificant

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the least-significant symbol is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    %X (%d)", test.Value, test.Value)
				continue
			}
		}
	}
}
