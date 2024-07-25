package runesymbols_test

import (
	"testing"

	"github.com/reiver/go-hexadeca/symbols/rune"
)

func TestLowerCase(t *testing.T) {

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

		actual := runesymbols.LowerCase(test.Value)
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
