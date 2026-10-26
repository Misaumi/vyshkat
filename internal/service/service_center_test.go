package service

import (
	"testing"

	"vyshkat/internal/domain"
)

func TestInspectAcceptsServiceableVehicle(t *testing.T) {
	center := NewBasicServiceCenter()
	v, _ := domain.NewVehicle("Самокат", "scooter", 8, 1)
	if err := center.Inspect(v); err != nil {
		t.Fatalf("serviceable vehicle must be accepted: %v", err)
	}
}

func TestInspectRejectsBrokenVehicle(t *testing.T) {
	center := NewBasicServiceCenter()
	v, _ := domain.NewVehicle("Самокат", "scooter", 8, 1)
	v.SetCondition(false)
	if err := center.Inspect(v); err == nil {
		t.Fatal("broken vehicle must be rejected")
	}
}
