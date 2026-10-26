package di

import (
	"errors"
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

func TestDefaultContainerBuildsWorkingFleet(t *testing.T) {
	c := NewDefault()
	if c.Fleet == nil {
		t.Fatal("container must provide a Fleet")
	}
	v, err := domain.NewVehicle("Самокат", "scooter", 8, 1.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := c.Fleet.AddVehicle(v); err != nil {
		t.Fatalf("fleet from container must accept vehicle: %v", err)
	}
	if got := v.InventoryNumber(); got != "VK-0001" {
		t.Fatalf("inventory number = %q, want VK-0001", got)
	}
}

func TestContainerAllowsInspectionSubstitution(t *testing.T) {
	center := &rejectingCenter{}
	c := New(center, service.NewInventoryNumberGenerator("VK"))

	v, _ := domain.NewVehicle("Плохой самокат", "scooter", 8, 1)
	if err := c.Fleet.AddVehicle(v); err == nil {
		t.Fatal("vehicle must be rejected by substituted inspection")
	}
	if c.Fleet.VehicleCount() != 0 {
		t.Fatal("rejected vehicle was added")
	}
	if v.InventoryNumber() != "" {
		t.Fatal("rejected vehicle received inventory number")
	}
}

type rejectingCenter struct{}

func (rejectingCenter) Inspect(*domain.Vehicle) error {
	return errors.New("fake: rejected")
}
