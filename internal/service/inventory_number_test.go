package service

import "testing"

func TestNextProducesSequentialNumbers(t *testing.T) {
	g := NewInventoryNumberGenerator("VK")
	want := []string{"VK-0001", "VK-0002", "VK-0003"}
	for _, expected := range want {
		if got := g.Next(); got != expected {
			t.Fatalf("Next() = %q, want %q", got, expected)
		}
	}
}

func TestNextUsesCustomPrefixAndFourDigits(t *testing.T) {
	g := NewInventoryNumberGenerator("TEST")
	if got := g.Next(); got != "TEST-0001" {
		t.Fatalf("Next() = %q, want TEST-0001", got)
	}
	for i := 0; i < 9; i++ {
		g.Next()
	}
	if got := g.Next(); got != "TEST-0011" {
		t.Fatalf("Next() = %q, want TEST-0011", got)
	}
}
