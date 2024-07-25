package hexadecabytes_test

import (
	"testing"

	"github.com/reiver/go-hexadeca/bytes"
)

func TestEncodeByteUsingLowerCaseSymbols(t *testing.T) {

	tests := []struct{
		Value byte
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
			Value:                  0x80,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x81,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x82,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x83,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x84,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x85,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x86,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x87,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x88,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x8A,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x8B,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x8C,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x8D,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x8E,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x8F,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x8F,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0x90,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x91,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x92,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x93,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x94,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x95,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x96,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x97,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x98,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x99,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x9A,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0x9B,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0x9C,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0x9D,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0x9E,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0x9F,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xA0,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xA1,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xA2,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xA3,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xA4,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xA5,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xA6,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xA7,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xA8,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xA9,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xAA,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xAB,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xAC,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xAD,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xAE,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xAF,
			ExpectedMostSignificant: 'a',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xB0,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xB1,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xB2,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xB3,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xB4,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xB5,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xB6,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xB7,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xB8,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xB9,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xBA,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xBB,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xBC,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xBD,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xBE,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xBF,
			ExpectedMostSignificant: 'b',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xC0,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xC1,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xC2,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xC3,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xC4,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xC5,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xC6,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xC7,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xC8,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xC9,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xCA,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xCB,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xCC,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xCD,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xCE,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xCF,
			ExpectedMostSignificant: 'c',
			ExpectedLeastSignificant: 'f',
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
			Value:                  0xE0,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xE1,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xE2,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xE3,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xE4,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xE5,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xE6,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xE7,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xE8,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xE9,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xEA,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'a',
		},
		{
			Value:                  0xEB,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'b',
		},
		{
			Value:                  0xEC,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'c',
		},
		{
			Value:                  0xED,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'd',
		},
		{
			Value:                  0xEE,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'e',
		},
		{
			Value:                  0xEF,
			ExpectedMostSignificant: 'e',
			ExpectedLeastSignificant: 'f',
		},
		{
			Value:                  0xF0,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xF1,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xF2,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xF3,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xF4,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xF5,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xF6,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xF7,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xF8,
			ExpectedMostSignificant: 'f',
			ExpectedLeastSignificant: '8',
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

		actualMostSignificant, actualLeastSignificant := hexadecabytes.EncodeByteUsingLowerCaseSymbols(test.Value)

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
			Value:                  0x80,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x81,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x82,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x83,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x84,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x85,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x86,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x87,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x88,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x89,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x8A,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x8B,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x8C,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x8D,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x8E,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x8F,
			ExpectedMostSignificant: '8',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0x90,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0x91,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0x92,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0x93,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0x94,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0x95,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0x96,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0x97,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0x98,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0x99,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0x9A,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0x9B,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0x9C,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0x9D,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0x9E,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0x9F,
			ExpectedMostSignificant: '9',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xA0,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xA1,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xA2,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xA3,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xA4,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xA5,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xA6,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xA7,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xA8,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xA9,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xAA,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xAB,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xAC,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xAD,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xAE,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xAF,
			ExpectedMostSignificant: 'A',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xB0,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xB1,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xB2,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xB3,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xB4,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xB5,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xB6,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xB7,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xB8,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xB9,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xBA,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xBB,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xBC,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xBD,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xBE,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xBF,
			ExpectedMostSignificant: 'B',
			ExpectedLeastSignificant: 'F',
		},
		{
			Value:                  0xC0,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xC1,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xC2,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xC3,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xC4,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xC5,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xC6,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xC7,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xC8,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '8',
		},
		{
			Value:                  0xC9,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: '9',
		},
		{
			Value:                  0xCA,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'A',
		},
		{
			Value:                  0xCB,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'B',
		},
		{
			Value:                  0xCC,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'C',
		},
		{
			Value:                  0xCD,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'D',
		},
		{
			Value:                  0xCE,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'E',
		},
		{
			Value:                  0xCF,
			ExpectedMostSignificant: 'C',
			ExpectedLeastSignificant: 'F',
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
			Value:                  0xF0,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '0',
		},
		{
			Value:                  0xF1,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '1',
		},
		{
			Value:                  0xF2,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '2',
		},
		{
			Value:                  0xF3,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '3',
		},
		{
			Value:                  0xF4,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '4',
		},
		{
			Value:                  0xF5,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '5',
		},
		{
			Value:                  0xF6,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '6',
		},
		{
			Value:                  0xF7,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '7',
		},
		{
			Value:                  0xF8,
			ExpectedMostSignificant: 'F',
			ExpectedLeastSignificant: '8',
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

		actualMostSignificant, actualLeastSignificant := hexadecabytes.EncodeByteUsingUpperCaseSymbols(test.Value)

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
