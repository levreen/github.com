package prose

import "testing"

func TestTwoElement(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't mathc expected value")
	}
}

func TestThreeElement(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}
}
