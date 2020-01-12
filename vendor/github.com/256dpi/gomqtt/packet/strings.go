package packet

import (
	"encoding/binary"
	"math"
)

// read length prefixed bytes
func readLPBytes(buf []byte, safe bool, t Type) ([]byte, int, error) {
	// check buffer
	if len(buf) < 2 {
		return nil, 0, makeError(t, "insufficient buffer size, expected 2, got %d", len(buf))
	}

	// read length
	length := int(binary.BigEndian.Uint16(buf))

	// check length
	if len(buf) < 2+length {
		return nil, 2, makeError(t, "insufficient buffer size, expected %d, got %d", 2+length, len(buf))
	}

	// get bytes
	bytes := buf[2 : 2+length]

	// return input buffer if not safe
	if !safe {
		return bytes, 2 + length, nil
	}

	// otherwise copy buffer
	cpy := make([]byte, length)
	copy(cpy, bytes)

	return cpy, 2 + length, nil
}

// read length prefixed string
func readLPString(buf []byte, t Type) (string, int, error) {
	bytes, n, err := readLPBytes(buf, false, t)
	return string(bytes), n, err
}

// write length prefixed bytes
func writeLPBytes(buf []byte, bytes []byte, t Type) (int, error) {
	// get length
	length := len(bytes)

	// check length
	if length > math.MaxUint16 {
		return 0, makeError(t, "length %d greater than allowed %d bytes", length, math.MaxUint16)
	}

	// check buffer
	if len(buf) < 2+length {
		return 0, makeError(t, "insufficient buffer size, expected %d, got %d", 2+length, len(buf))
	}

	// write length
	binary.BigEndian.PutUint16(buf, uint16(length))

	// write bytes
	copy(buf[2:], bytes)

	return 2 + length, nil
}

// write length prefixed string
func writeLPString(buf []byte, str string, t Type) (int, error) {
	return writeLPBytes(buf, []byte(str), t)
}
