package hexadeca_test

import (
	"testing"

	"sourcecode.social/reiver/go-hexadeca"
)

func TestSymbolLowerCase(t *testing.T) {

	tests := []struct{
		Value byte
		Expected rune
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

		actual := hexadeca.SymbolLowerCase(test.Value)
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

func TestSymbolPersian(t *testing.T) {

	tests := []struct{
		Value byte
		Expected rune
	}{
		{
			Value:     0,
			Expected: '۰',
		},
		{
			Value:     1,
			Expected: '۱',
		},
		{
			Value:     2,
			Expected: '۲',
		},
		{
			Value:     3,
			Expected: '۳',
		},
		{
			Value:     4,
			Expected: '۴',
		},
		{
			Value:     5,
			Expected: '۵',
		},
		{
			Value:     6,
			Expected: '۶',
		},
		{
			Value:     7,
			Expected: '۷',
		},
		{
			Value:     8,
			Expected: '۸',
		},
		{
			Value:     9,
			Expected: '۹',
		},
		{
			Value:    10,
			Expected: 'ا',
		},
		{
			Value:    11,
			Expected: 'ب',
		},
		{
			Value:    12,
			Expected: 'پ',
		},
		{
			Value:    13,
			Expected: 'ت',
		},
		{
			Value:    14,
			Expected: 'ث',
		},
		{
			Value:    15,
			Expected: 'ج',
		},
	}

	for testNumber, test := range tests {

		actual := hexadeca.SymbolPersian(test.Value)
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

func TestSymbolUpperCase(t *testing.T) {

	tests := []struct{
		Value byte
		Expected rune
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
			Expected: 'A',
		},
		{
			Value:    11,
			Expected: 'B',
		},
		{
			Value:    12,
			Expected: 'C',
		},
		{
			Value:    13,
			Expected: 'D',
		},
		{
			Value:    14,
			Expected: 'E',
		},
		{
			Value:    15,
			Expected: 'F',
		},
	}

	for testNumber, test := range tests {

		actual := hexadeca.SymbolUpperCase(test.Value)
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
