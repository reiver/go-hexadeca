package hexadecabytes_test

import (
	"testing"

	"github.com/reiver/go-hexadeca/bytes"
)

func TestDecodeUint64UsingLowerCaseSymbols(t *testing.T) {

	tests := []struct{
		Symbol0 byte
		Symbol1 byte
		Symbol2 byte
		Symbol3 byte
		Symbol4 byte
		Symbol5 byte
		Symbol6 byte
		Symbol7 byte
		Symbol8 byte
		Symbol9 byte
		Symbol10 byte
		Symbol11 byte
		Symbol12 byte
		Symbol13 byte
		Symbol14 byte
		Symbol15 byte
		Expected uint64
	}{
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000000000,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '1',
			Expected: 0x0000000000000001,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '2',
			Expected: 0x0000000000000002,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '3',
			Expected: 0x0000000000000003,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 'f',
			Symbol0:                  '0',
			Expected: 0x00000000000000f0,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                'f',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000000f00,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               'f',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000000000f000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              'f',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000000000f0000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             'f',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000f00000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            'f',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000000f000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           'f',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000000f0000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          'f',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000f00000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         'f',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000f000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       'f',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000f0000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      'f',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000f00000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     'f',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000f000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    'f',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00f0000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   'f',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0f00000000000000,
		},



		{
			Symbol15:  'f',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0xf000000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '1',
			Symbol13:    '2',
			Symbol12:     '3',
			Symbol11:      '4',
			Symbol10:       '5',
			Symbol9:         '6',
			Symbol8:          '7',
			Symbol7:           '8',
			Symbol6:            '9',
			Symbol5:             'a',
			Symbol4:              'b',
			Symbol3:               'c',
			Symbol2:                'd',
			Symbol1:                 'e',
			Symbol0:                  'f',
			Expected: 0x0123456789abcdef,
		},



		{
			Symbol15:  'f',
			Symbol14:   'e',
			Symbol13:    'd',
			Symbol12:     'c',
			Symbol11:      'b',
			Symbol10:       'a',
			Symbol9:         '9',
			Symbol8:          '8',
			Symbol7:           '7',
			Symbol6:            '6',
			Symbol5:             '5',
			Symbol4:              '4',
			Symbol3:               '3',
			Symbol2:                '2',
			Symbol1:                 '1',
			Symbol0:                  '0',
			Expected: 0xfedcba9876543210,
		},
	}

	for testNumber, test := range tests {

		actual, err := hexadecabytes.DecodeUint64UsingLowerCaseSymbols(
			test.Symbol15,
			test.Symbol14,
			test.Symbol13,
			test.Symbol12,
			test.Symbol11,
			test.Symbol10,
			test.Symbol9,
			test.Symbol8,
			test.Symbol7,
			test.Symbol6,
			test.Symbol5,
			test.Symbol4,
			test.Symbol3,
			test.Symbol2,
			test.Symbol1,
			test.Symbol0,
		)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SYMBOL-15: %q (%U)", test.Symbol15, test.Symbol15)
			t.Logf("SYMBOL-14: %q (%U)", test.Symbol14, test.Symbol14)
			t.Logf("SYMBOL-13: %q (%U)", test.Symbol13, test.Symbol13)
			t.Logf("SYMBOL-12: %q (%U)", test.Symbol12, test.Symbol12)
			t.Logf("SYMBOL-11: %q (%U)", test.Symbol11, test.Symbol11)
			t.Logf("SYMBOL-10: %q (%U)", test.Symbol10, test.Symbol10)
			t.Logf("SYMBOL-9:  %q (%U)", test.Symbol9,  test.Symbol9)
			t.Logf("SYMBOL-8:  %q (%U)", test.Symbol8,  test.Symbol8)
			t.Logf("SYMBOL-7:  %q (%U)", test.Symbol7,  test.Symbol7)
			t.Logf("SYMBOL-6:  %q (%U)", test.Symbol6,  test.Symbol6)
			t.Logf("SYMBOL-5:  %q (%U)", test.Symbol5,  test.Symbol5)
			t.Logf("SYMBOL-4:  %q (%U)", test.Symbol4,  test.Symbol4)
			t.Logf("SYMBOL-3:  %q (%U)", test.Symbol3,  test.Symbol3)
			t.Logf("SYMBOL-2:  %q (%U)", test.Symbol2,  test.Symbol2)
			t.Logf("SYMBOL-1:  %q (%U)", test.Symbol1,  test.Symbol1)
			t.Logf("SYMBOL-0:  %q (%U)", test.Symbol0,  test.Symbol0)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual decoded value is not what was expected.", testNumber)
			t.Logf("EXPECTED: 0x%016x (%d)", expected, expected)
			t.Logf("ACTUAL:   0x%016x (%d)", actual, actual)
			t.Logf("SYMBOL-15: %q (%U)", test.Symbol15, test.Symbol15)
			t.Logf("SYMBOL-14:  %q (%U)", test.Symbol14, test.Symbol14)
			t.Logf("SYMBOL-13:   %q (%U)", test.Symbol13, test.Symbol13)
			t.Logf("SYMBOL-12:    %q (%U)", test.Symbol12, test.Symbol12)
			t.Logf("SYMBOL-11:     %q (%U)", test.Symbol11, test.Symbol11)
			t.Logf("SYMBOL-10:      %q (%U)", test.Symbol10, test.Symbol10)
			t.Logf("SYMBOL-9:        %q (%U)", test.Symbol9,  test.Symbol9)
			t.Logf("SYMBOL-8:         %q (%U)", test.Symbol8,  test.Symbol8)
			t.Logf("SYMBOL-7:          %q (%U)", test.Symbol7,  test.Symbol7)
			t.Logf("SYMBOL-6:           %q (%U)", test.Symbol6,  test.Symbol6)
			t.Logf("SYMBOL-5:            %q (%U)", test.Symbol5,  test.Symbol5)
			t.Logf("SYMBOL-4:             %q (%U)", test.Symbol4,  test.Symbol4)
			t.Logf("SYMBOL-3:              %q (%U)", test.Symbol3,  test.Symbol3)
			t.Logf("SYMBOL-2:               %q (%U)", test.Symbol2,  test.Symbol2)
			t.Logf("SYMBOL-1:                %q (%U)", test.Symbol1,  test.Symbol1)
			t.Logf("SYMBOL-0:                 %q (%U)", test.Symbol0,  test.Symbol0)
			continue
		}
	}
}

func TestDecodeUint64UsingUpperCaseSymbols(t *testing.T) {

	tests := []struct{
		Symbol0 byte
		Symbol1 byte
		Symbol2 byte
		Symbol3 byte
		Symbol4 byte
		Symbol5 byte
		Symbol6 byte
		Symbol7 byte
		Symbol8 byte
		Symbol9 byte
		Symbol10 byte
		Symbol11 byte
		Symbol12 byte
		Symbol13 byte
		Symbol14 byte
		Symbol15 byte
		Expected uint64
	}{
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000000000,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '1',
			Expected: 0x0000000000000001,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '2',
			Expected: 0x0000000000000002,
		},
		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '3',
			Expected: 0x0000000000000003,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 'F',
			Symbol0:                  '0',
			Expected: 0x00000000000000F0,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                'F',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000000F00,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               'F',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000000000F000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              'F',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000000000F0000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             'F',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000000F00000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            'F',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000000F000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           'F',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000000F0000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          'F',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000000F00000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         'F',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000000F000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       'F',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00000F0000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      'F',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0000F00000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     'F',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x000F000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '0',
			Symbol13:    'F',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x00F0000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   'F',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0x0F00000000000000,
		},



		{
			Symbol15:  'F',
			Symbol14:   '0',
			Symbol13:    '0',
			Symbol12:     '0',
			Symbol11:      '0',
			Symbol10:       '0',
			Symbol9:         '0',
			Symbol8:          '0',
			Symbol7:           '0',
			Symbol6:            '0',
			Symbol5:             '0',
			Symbol4:              '0',
			Symbol3:               '0',
			Symbol2:                '0',
			Symbol1:                 '0',
			Symbol0:                  '0',
			Expected: 0xF000000000000000,
		},



		{
			Symbol15:  '0',
			Symbol14:   '1',
			Symbol13:    '2',
			Symbol12:     '3',
			Symbol11:      '4',
			Symbol10:       '5',
			Symbol9:         '6',
			Symbol8:          '7',
			Symbol7:           '8',
			Symbol6:            '9',
			Symbol5:             'A',
			Symbol4:              'B',
			Symbol3:               'C',
			Symbol2:                'D',
			Symbol1:                 'E',
			Symbol0:                  'F',
			Expected: 0x0123456789ABCDEF,
		},



		{
			Symbol15:  'F',
			Symbol14:   'E',
			Symbol13:    'D',
			Symbol12:     'C',
			Symbol11:      'B',
			Symbol10:       'A',
			Symbol9:         '9',
			Symbol8:          '8',
			Symbol7:           '7',
			Symbol6:            '6',
			Symbol5:             '5',
			Symbol4:              '4',
			Symbol3:               '3',
			Symbol2:                '2',
			Symbol1:                 '1',
			Symbol0:                  '0',
			Expected: 0xFEDCBA9876543210,
		},
	}

	for testNumber, test := range tests {

		actual, err := hexadecabytes.DecodeUint64UsingUpperCaseSymbols(
			test.Symbol15,
			test.Symbol14,
			test.Symbol13,
			test.Symbol12,
			test.Symbol11,
			test.Symbol10,
			test.Symbol9,
			test.Symbol8,
			test.Symbol7,
			test.Symbol6,
			test.Symbol5,
			test.Symbol4,
			test.Symbol3,
			test.Symbol2,
			test.Symbol1,
			test.Symbol0,
		)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SYMBOL-15: %q (%U)", test.Symbol15, test.Symbol15)
			t.Logf("SYMBOL-14: %q (%U)", test.Symbol14, test.Symbol14)
			t.Logf("SYMBOL-13: %q (%U)", test.Symbol13, test.Symbol13)
			t.Logf("SYMBOL-12: %q (%U)", test.Symbol12, test.Symbol12)
			t.Logf("SYMBOL-11: %q (%U)", test.Symbol11, test.Symbol11)
			t.Logf("SYMBOL-10: %q (%U)", test.Symbol10, test.Symbol10)
			t.Logf("SYMBOL-9:  %q (%U)", test.Symbol9,  test.Symbol9)
			t.Logf("SYMBOL-8:  %q (%U)", test.Symbol8,  test.Symbol8)
			t.Logf("SYMBOL-7:  %q (%U)", test.Symbol7,  test.Symbol7)
			t.Logf("SYMBOL-6:  %q (%U)", test.Symbol6,  test.Symbol6)
			t.Logf("SYMBOL-5:  %q (%U)", test.Symbol5,  test.Symbol5)
			t.Logf("SYMBOL-4:  %q (%U)", test.Symbol4,  test.Symbol4)
			t.Logf("SYMBOL-3:  %q (%U)", test.Symbol3,  test.Symbol3)
			t.Logf("SYMBOL-2:  %q (%U)", test.Symbol2,  test.Symbol2)
			t.Logf("SYMBOL-1:  %q (%U)", test.Symbol1,  test.Symbol1)
			t.Logf("SYMBOL-0:  %q (%U)", test.Symbol0,  test.Symbol0)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual decoded value is not what was expected.", testNumber)
			t.Logf("EXPECTED: 0x%016x (%d)", expected, expected)
			t.Logf("ACTUAL:   0x%016x (%d)", actual, actual)
			t.Logf("SYMBOL-15: %q (%U)", test.Symbol15, test.Symbol15)
			t.Logf("SYMBOL-14:  %q (%U)", test.Symbol14, test.Symbol14)
			t.Logf("SYMBOL-13:   %q (%U)", test.Symbol13, test.Symbol13)
			t.Logf("SYMBOL-12:    %q (%U)", test.Symbol12, test.Symbol12)
			t.Logf("SYMBOL-11:     %q (%U)", test.Symbol11, test.Symbol11)
			t.Logf("SYMBOL-10:      %q (%U)", test.Symbol10, test.Symbol10)
			t.Logf("SYMBOL-9:        %q (%U)", test.Symbol9,  test.Symbol9)
			t.Logf("SYMBOL-8:         %q (%U)", test.Symbol8,  test.Symbol8)
			t.Logf("SYMBOL-7:          %q (%U)", test.Symbol7,  test.Symbol7)
			t.Logf("SYMBOL-6:           %q (%U)", test.Symbol6,  test.Symbol6)
			t.Logf("SYMBOL-5:            %q (%U)", test.Symbol5,  test.Symbol5)
			t.Logf("SYMBOL-4:             %q (%U)", test.Symbol4,  test.Symbol4)
			t.Logf("SYMBOL-3:              %q (%U)", test.Symbol3,  test.Symbol3)
			t.Logf("SYMBOL-2:               %q (%U)", test.Symbol2,  test.Symbol2)
			t.Logf("SYMBOL-1:                %q (%U)", test.Symbol1,  test.Symbol1)
			t.Logf("SYMBOL-0:                 %q (%U)", test.Symbol0,  test.Symbol0)
			continue
		}
	}
}

