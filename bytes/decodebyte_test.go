package hexadecabytes_test

import (
	"testing"

	"github.com/reiver/go-hexadeca/bytes"
)

func TestDecodeByteUsingLowerCaseSymbols(t *testing.T) {

	tests := []struct{
		MostSignificant byte
		LeastSignificant byte
		Expected byte
	}{
		{
			MostSignificant: '0',
			LeastSignificant: '0',
			Expected:       0x00,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '1',
			Expected:       0x01,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '2',
			Expected:       0x02,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '3',
			Expected:       0x03,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '4',
			Expected:       0x04,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '5',
			Expected:       0x05,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '6',
			Expected:       0x06,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '7',
			Expected:       0x07,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '8',
			Expected:       0x08,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '9',
			Expected:       0x09,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'a',
			Expected:       0x0a,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'b',
			Expected:       0x0b,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'c',
			Expected:       0x0c,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'd',
			Expected:       0x0d,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'e',
			Expected:       0x0e,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'f',
			Expected:       0x0f,
		},
		{
			MostSignificant: '1',
			LeastSignificant: '0',
			Expected:       0x10,
		},



		{
			MostSignificant: '2',
			LeastSignificant: '1',
			Expected:       0x21,
		},



		{
			MostSignificant: '5',
			LeastSignificant: '0',
			Expected:       0x50,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '1',
			Expected:       0x51,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '2',
			Expected:       0x52,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '3',
			Expected:       0x53,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '4',
			Expected:       0x54,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '5',
			Expected:       0x55,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '6',
			Expected:       0x56,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '7',
			Expected:       0x57,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '8',
			Expected:       0x58,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '9',
			Expected:       0x59,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'a',
			Expected:       0x5a,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'b',
			Expected:       0x5b,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'c',
			Expected:       0x5c,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'd',
			Expected:       0x5d,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'e',
			Expected:       0x5e,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'f',
			Expected:       0x5f,
		},



		{
			MostSignificant: 'f',
			LeastSignificant: '9',
			Expected:       0xf9,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'a',
			Expected:       0xfa,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'b',
			Expected:       0xfb,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'c',
			Expected:       0xfc,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'd',
			Expected:       0xfd,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'e',
			Expected:       0xfe,
		},
		{
			MostSignificant: 'f',
			LeastSignificant: 'f',
			Expected:       0xff,
		},
	}

	for testNumber, test := range tests {

		actual, err := hexadecabytes.DecodeByteUsingLowerCaseSymbols(test.MostSignificant, test.LeastSignificant)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("MOST-SIGNIFICANT: %q (%U)", test.MostSignificant, test.MostSignificant)
			t.Logf("LEAST-SIGNIFICANT: %q (%U)", test.LeastSignificant, test.LeastSignificant)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual byte is not what was expected.", testNumber)
			t.Logf("EXPECTED:        0x%02x (%d)", expected, expected)
			t.Logf("ACTUAL:          0x%02x (%d)", actual, actual)
			t.Logf("MOST-SIGNIFICANT: %q (%U)", test.MostSignificant, test.MostSignificant)
			t.Logf("LEAST-SIGNIFICANT: %q (%U)", test.LeastSignificant, test.LeastSignificant)
			continue
		}
	}
}

func TestDecodeByteUsingUpperCaseSymbols(t *testing.T) {

	tests := []struct{
		MostSignificant byte
		LeastSignificant byte
		Expected byte
	}{
		{
			MostSignificant: '0',
			LeastSignificant: '0',
			Expected:       0x00,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '1',
			Expected:       0x01,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '2',
			Expected:       0x02,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '3',
			Expected:       0x03,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '4',
			Expected:       0x04,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '5',
			Expected:       0x05,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '6',
			Expected:       0x06,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '7',
			Expected:       0x07,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '8',
			Expected:       0x08,
		},
		{
			MostSignificant: '0',
			LeastSignificant: '9',
			Expected:       0x09,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'A',
			Expected:       0x0A,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'B',
			Expected:       0x0B,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'C',
			Expected:       0x0C,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'D',
			Expected:       0x0D,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'E',
			Expected:       0x0E,
		},
		{
			MostSignificant: '0',
			LeastSignificant: 'F',
			Expected:       0x0F,
		},
		{
			MostSignificant: '1',
			LeastSignificant: '0',
			Expected:       0x10,
		},



		{
			MostSignificant: '2',
			LeastSignificant: '1',
			Expected:       0x21,
		},



		{
			MostSignificant: '5',
			LeastSignificant: '0',
			Expected:       0x50,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '1',
			Expected:       0x51,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '2',
			Expected:       0x52,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '3',
			Expected:       0x53,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '4',
			Expected:       0x54,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '5',
			Expected:       0x55,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '6',
			Expected:       0x56,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '7',
			Expected:       0x57,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '8',
			Expected:       0x58,
		},
		{
			MostSignificant: '5',
			LeastSignificant: '9',
			Expected:       0x59,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'A',
			Expected:       0x5A,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'B',
			Expected:       0x5B,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'C',
			Expected:       0x5C,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'D',
			Expected:       0x5D,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'E',
			Expected:       0x5E,
		},
		{
			MostSignificant: '5',
			LeastSignificant: 'F',
			Expected:       0x5F,
		},



		{
			MostSignificant: 'F',
			LeastSignificant: '9',
			Expected:       0xF9,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'A',
			Expected:       0xFA,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'B',
			Expected:       0xFB,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'C',
			Expected:       0xFC,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'D',
			Expected:       0xFD,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'E',
			Expected:       0xFE,
		},
		{
			MostSignificant: 'F',
			LeastSignificant: 'F',
			Expected:       0xFF,
		},
	}

	for testNumber, test := range tests {

		actual, err := hexadecabytes.DecodeByteUsingUpperCaseSymbols(test.MostSignificant, test.LeastSignificant)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("MOST-SIGNIFICANT: %q (%U)", test.MostSignificant, test.MostSignificant)
			t.Logf("LEAST-SIGNIFICANT: %q (%U)", test.LeastSignificant, test.LeastSignificant)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual byte is not what was expected.", testNumber)
			t.Logf("EXPECTED:        0x%02x (%d)", expected, expected)
			t.Logf("ACTUAL:          0x%02x (%d)", actual, actual)
			t.Logf("MOST-SIGNIFICANT: %q (%U)", test.MostSignificant, test.MostSignificant)
			t.Logf("LEAST-SIGNIFICANT: %q (%U)", test.LeastSignificant, test.LeastSignificant)
			continue
		}
	}
}
