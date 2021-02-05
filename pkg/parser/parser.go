package parser

import (
	"bytes"
	"encoding/hex"
	"math"
	"strconv"
)

// BYTELEN length of bytes
const BYTELEN = 320

// IdxOfHead .
func IdxOfHead(hex string) (int, int) {
	if len(hex) < BYTELEN {
		return 0, 0
	}

	for i := 0; i < len(hex)-1; i++ {
		if hex[i:i+2] == "26" {
			return i - 4, i - 4 + BYTELEN
		}
	}
	return 0, 0
}

// IsFine .
func IsFine(hex string) bool {
	return hex == "00"
}

// IsReady .
func IsReady(hex byte) bool {
	return hex == 5
}

// IsProcess .
func IsProcess(hex string) bool {
	return hex == "0a"
}

// AckFinish .
func AckFinish(hex string) bool {
	return hex == "1e"
}

// HexToFloat32 .
func HexToFloat32(hex string) float32 {
	n, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0
	}
	f := math.Float32frombits(uint32(n))
	return f
}

// HexToInt16 .
func HexToInt16(hex string) int16 {
	n, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0
	}

	return int16(n)
}

// HexToBool .
func HexToBool(hex string) bool {
	return hex == "00"
}

// HexToString .
func HexToString(h string) string {
	r, _ := hex.DecodeString(h)

	buf := bytes.Buffer{}
	for _, v := range r {
		buf.WriteByte(v)
	}

	return buf.String()
}
