package misc

import "testing"

func TestOppositeSigns(t *testing.T) {
	neg, pos, zero := -1, 1, 0

	if !HaveOppositeSigns(neg, pos) {
		t.Error("-1 and 1 should have opposite signs.")
	}

	if !HaveOppositeSigns(neg, zero) {
		t.Error("-1 and 0 should have opposite signs.")
	}

	if HaveOppositeSigns(zero, pos) {
		t.Error("1 and 0 should not have opposite signs.")
	}
}

func TestMin(t *testing.T) {
	if Min(2, 4) != 2 {
		t.Error("Min between 2 and 4 should be 2")
	}

	if Min(-1, 4) != -1 {
		t.Error("Min between -1 and 4 should be -1")
	}

	if Min(8, 4) != 4 {
		t.Error("Min between 8 and 4 should be 4")
	}

	if Min(0, 0) != 0 {
		t.Error("Min between 0 and 0 should be 0")
	}

	if Min(-8, -42) != -42 {
		t.Error("Min between -8 and -42 should be  -42")
	}
}

func TestMax(t *testing.T) {
	if Max(2, 4) != 4 {
		t.Error("Max between 2 and 4 should be 4")
	}

	if Max(-1, 4) != 4 {
		t.Error("Max between -1 and 4 should be 4")
	}

	if Max(8, 4) != 8 {
		t.Error("Max between 8 and 4 should be 8")
	}

	if Max(0, 0) != 0 {
		t.Error("Max between 0 and 0 should be 0")
	}

	if Max(-8, -42) != -8 {
		t.Error("Max between -8 and -42 should be  -8")
	}
}
