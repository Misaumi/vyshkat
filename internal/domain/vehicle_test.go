package domain

import "testing"

func TestNewVehicle(t *testing.T) {
	v, err := NewVehicle("Городской самокат", "scooter", 8, 1.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.Name() != "Городской самокат" || v.Kind() != "scooter" || v.BeginnerEase() != 8 {
		t.Fatalf("vehicle = %q/%q/%d", v.Name(), v.Kind(), v.BeginnerEase())
	}
	if v.DailyEnergyConsumption() != 1.5 {
		t.Fatalf("energy = %v", v.DailyEnergyConsumption())
	}
}

func TestNewVehicleValidation(t *testing.T) {
	cases := []struct {
		name   string
		kind   string
		ease   int
		energy float64
	}{
		{"", "scooter", 5, 1},
		{"Самокат", "", 5, 1},
		{"Самокат", "scooter", 0, 1},
		{"Самокат", "scooter", 11, 1},
		{"Самокат", "scooter", 5, -0.5},
	}
	for _, c := range cases {
		if _, err := NewVehicle(c.name, c.kind, c.ease, c.energy); err == nil {
			t.Errorf("NewVehicle(%q, %q, %d, %v): expected error", c.name, c.kind, c.ease, c.energy)
		}
	}
}

func TestVehicleCondition(t *testing.T) {
	v, _ := NewVehicle("Самокат", "scooter", 7, 1)
	if !v.ConditionOK() {
		t.Fatal("new vehicle must be serviceable")
	}
	v.SetCondition(false)
	if v.ConditionOK() {
		t.Fatal("condition must be switchable")
	}
}
