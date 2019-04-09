package bit

import "testing"

func TestSetBit(t *testing.T) {
	i := 192

	res := SetBit(i, 0)

	if res != 193 {
		t.Error("192: 1100 0000 if we set the bit at pos 0 we should have 1100 0001 which is the binary representation of 193")
	}
}

func TestGetBit(t *testing.T) {
	i := 80

	if GetBit(i, 0) {
		t.Error("80: 0101 0000 bit at pos 0 should not be set")
	}

	if !GetBit(i, 4) {
		t.Error("80: 0101 0000 bit at pos 4 should be set")
	}
}
