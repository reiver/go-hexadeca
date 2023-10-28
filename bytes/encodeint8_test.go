package hexadecabytes_test

import (
	"testing"

	"sourcecode.social/reiver/go-hexadeca/bytes"
)

func TestEncodeInt8UsingLowerCaseSymbols(t *testing.T) {

	tests := []struct{
		Value int8
		ExpectedMostSignificant byte
		ExpectedLeastSignificant byte
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
			Value:                  0x20,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x21,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x22,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x23,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x24,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x25,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x26,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x27,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x28,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x29,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x2A,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x2B,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x2C,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x2D,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x2E,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x2F,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x30,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x31,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x32,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x33,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x34,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x35,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x36,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x37,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x38,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x39,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x3A,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x3B,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x3C,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x3D,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x3E,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x3F,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x40,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x41,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x42,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x43,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x44,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x45,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x46,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x47,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x48,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x49,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x4A,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x4B,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x4C,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x4D,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x4E,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x4F,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x50,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x51,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x52,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x53,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x54,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x55,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x56,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x57,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x58,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x59,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x5A,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x5B,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x5C,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x5D,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x5E,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x5F,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x60,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x61,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x62,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x63,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x64,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x65,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x66,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x67,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x68,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x69,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x6A,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x6B,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x6C,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x6D,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x6E,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x6F,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x70,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x71,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x72,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x73,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x74,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x75,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x76,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x77,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x78,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x79,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x7A,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x7B,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x7C,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x7D,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x7E,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x7F,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x80 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x81 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x82 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x83 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x84 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x85 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x86 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x87 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x88 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x8A - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x8B - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x8C - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x8D - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x8E - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x8F - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x8F - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x90 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x91 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x92 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x93 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x94 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x95 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x96 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x97 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x98 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x99 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x9A - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x9B - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x9C - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x9D - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x9E - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x9F - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xA0 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xA1 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xA2 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xA3 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xA4 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xA5 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xA6 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xA7 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xA8 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xA9 - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xAA - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xAB - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xAC - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xAD - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xAE - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xAF - 0x100,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xB0 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xB1 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xB2 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xB3 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xB4 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xB5 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xB6 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xB7 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xB8 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xB9 - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xBA - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xBB - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xBC - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xBD - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xBE - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xBF - 0x100,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xC0 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xC1 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xC2 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xC3 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xC4 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xC5 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xC6 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xC7 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xC8 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xC9 - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xCA - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xCB - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xCC - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xCD - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xCE - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xCF - 0x100,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xD0 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xD1 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xD2 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xD3 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xD4 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xD5 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xD6 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xD7 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xD8 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xD9 - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xDA - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xDB - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xDC - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xDD - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xDE - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xDF - 0x100,
			ExpectedMostSignificant: 'd',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xE0 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xE1 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xE2 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xE3 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xE4 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xE5 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xE6 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xE7 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xE8 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xE9 - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xEA - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xEB - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xEC - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xED - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xEE - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xEF - 0x100,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xF0 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xF1 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xF2 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xF3 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xF4 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xF5 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xF6 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xF7 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xF8 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xF9 - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xFA - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xFB - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xFC - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xFD - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xFE - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xFF - 0x100,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: 'f',
		},
	}

	for testNumber, test := range tests {

		actualMostSignificant, actualLeastSignificant := hexadecabytes.EncodeInt8UsingLowerCaseSymbols(test.Value)

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

func TestEncodeInt8UsingUpperCaseSymbols(t *testing.T) {

	tests := []struct{
		Value int8
		ExpectedMostSignificant byte
		ExpectedLeastSignificant byte
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
			Value:                  0x20,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x21,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x22,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x23,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x24,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x25,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x26,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x27,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x28,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x29,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x2A,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x2B,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x2C,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x2D,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x2E,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x2F,
			ExpectedMostSignificant: '2',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x30,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x31,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x32,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x33,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x34,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x35,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x36,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x37,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x38,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x39,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x3A,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x3B,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x3C,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x3D,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x3E,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x3F,
			ExpectedMostSignificant: '3',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x40,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x41,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x42,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x43,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x44,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x45,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x46,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x47,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x48,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x49,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x4A,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x4B,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x4C,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x4D,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x4E,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x4F,
			ExpectedMostSignificant: '4',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x50,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x51,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x52,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x53,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x54,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x55,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x56,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x57,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x58,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x59,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x5A,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x5B,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x5C,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x5D,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x5E,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x5F,
			ExpectedMostSignificant: '5',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x60,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x61,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x62,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x63,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x64,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x65,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x66,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x67,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x68,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x69,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x6A,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x6B,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x6C,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x6D,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x6E,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x6F,
			ExpectedMostSignificant: '6',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x70,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x71,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x72,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x73,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x74,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x75,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x76,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x77,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x78,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x79,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x7A,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x7B,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x7C,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x7D,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x7E,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x7F,
			ExpectedMostSignificant: '7',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x80 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x81 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x82 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x83 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x84 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x85 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x86 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x87 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x88 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x89 - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x8A - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x8B - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x8C - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x8D - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x8E - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x8F - 0x100,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x90 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x91 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x92 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x93 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x94 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x95 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x96 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x97 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x98 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x99 - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x9A - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x9B - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x9C - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x9D - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x9E - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x9F - 0x100,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xA0 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xA1 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xA2 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xA3 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xA4 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xA5 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xA6 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xA7 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xA8 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xA9 - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xAA - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xAB - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xAC - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xAD - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xAE - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xAF - 0x100,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xB0 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xB1 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xB2 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xB3 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xB4 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xB5 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xB6 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xB7 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xB8 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xB9 - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xBA - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xBB - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xBC - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xBD - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xBE - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xBF - 0x100,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xC0 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xC1 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xC2 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xC3 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xC4 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xC5 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xC6 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xC7 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xC8 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xC9 - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xCA - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xCB - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xCC - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xCD - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xCE - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xCF - 0x100,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xD0 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xD1 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xD2 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xD3 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xD4 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xD5 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xD6 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xD7 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xD8 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xD9 - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xDA - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xDB - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xDC - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xDD - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xDE - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xDF - 0x100,
			ExpectedMostSignificant: 'D',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xF0 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xF1 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xF2 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xF3 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xF4 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xF5 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xF6 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xF7 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xF8 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xF9 - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xFA - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xFB - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xFC - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xFD - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xFE - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xFF - 0x100,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: 'F',
		},
	}

	for testNumber, test := range tests {

		actualMostSignificant, actualLeastSignificant := hexadecabytes.EncodeInt8UsingUpperCaseSymbols(test.Value)

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
