package bytesymbols_test

import (
	"testing"

	"github.com/reiver/go-hexadeca/symbols/byte"
)

func TestFromLowerCase(t *testing.T) {

	tests := []struct{
		Symbol byte
		ExpectedByte byte
		ExpectedOK bool
	}{
		{
			Symbol:      '0',
			ExpectedByte: 0,
			ExpectedOK : true,
		},
		{
			Symbol:      '1',
			ExpectedByte: 1,
			ExpectedOK : true,
		},
		{
			Symbol:      '2',
			ExpectedByte: 2,
			ExpectedOK : true,
		},
		{
			Symbol:      '3',
			ExpectedByte: 3,
			ExpectedOK : true,
		},
		{
			Symbol:      '4',
			ExpectedByte: 4,
			ExpectedOK : true,
		},
		{
			Symbol:      '5',
			ExpectedByte: 5,
			ExpectedOK : true,
		},
		{
			Symbol:      '6',
			ExpectedByte: 6,
			ExpectedOK : true,
		},
		{
			Symbol:      '7',
			ExpectedByte: 7,
			ExpectedOK : true,
		},
		{
			Symbol:      '8',
			ExpectedByte: 8,
			ExpectedOK : true,
		},
		{
			Symbol:      '9',
			ExpectedByte: 9,
			ExpectedOK : true,
		},
		{
			Symbol:      'a',
			ExpectedByte: 10,
			ExpectedOK : true,
		},
		{
			Symbol:      'b',
			ExpectedByte: 11,
			ExpectedOK : true,
		},
		{
			Symbol:      'c',
			ExpectedByte: 12,
			ExpectedOK : true,
		},
		{
			Symbol:      'd',
			ExpectedByte: 13,
			ExpectedOK : true,
		},
		{
			Symbol:      'e',
			ExpectedByte: 14,
			ExpectedOK : true,
		},
		{
			Symbol:      'f',
			ExpectedByte: 15,
			ExpectedOK : true,
		},



		{
			Symbol:      'z',
			ExpectedByte: 0,
			ExpectedOK : false,
		},
	}

	for testNumber, test := range tests {

		actualByte, actualOK := bytesymbols.FromLowerCase(test.Symbol)

		{
			expected := test.ExpectedOK
			actual := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("SYMBOL: %q (%U)", test.Symbol, test.Symbol)
				continue
			}
		}

		{
			expected := test.ExpectedByte
			actual := actualByte

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' is not what was expected.", testNumber)
				t.Logf("EXPECTED: 0x%X", expected)
				t.Logf("ACTUAL:   0x%X", actual)
				t.Logf("SYMBOL: %q (%U)", test.Symbol, test.Symbol)
				continue
			}
		}
	}
}

func TestSymbolLowerCase(t *testing.T) {

	tests := []struct{
		Value byte
		Expected byte
	}{
		{
			Value:     0,
			Expected: '0',
		},
		{
			Value:     1,
			Expected: '1',
		},
		{
			Value:     2,
			Expected: '2',
		},
		{
			Value:     3,
			Expected: '3',
		},
		{
			Value:     4,
			Expected: '4',
		},
		{
			Value:     5,
			Expected: '5',
		},
		{
			Value:     6,
			Expected: '6',
		},
		{
			Value:     7,
			Expected: '7',
		},
		{
			Value:     8,
			Expected: '8',
		},
		{
			Value:     9,
			Expected: '9',
		},
		{
			Value:    10,
			Expected: 'a',
		},
		{
			Value:    11,
			Expected: 'b',
		},
		{
			Value:    12,
			Expected: 'c',
		},
		{
			Value:    13,
			Expected: 'd',
		},
		{
			Value:    14,
			Expected: 'e',
		},
		{
			Value:    15,
			Expected: 'f',
		},
	}

	for testNumber, test := range tests {

		actual := bytesymbols.LowerCase(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q (%d)", expected, expected)
			t.Logf("ACTUAL:   %q (%d)", actual, actual)
			t.Logf("VALUE:     %x (%d)", test.Value, test.Value)
			continue
		}

	}
}
