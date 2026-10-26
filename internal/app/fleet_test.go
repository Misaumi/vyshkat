package app

import (
	"errors"
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

type fakeInspection struct {
	accept bool
	calls  int
}

func (f *fakeInspection) Inspect(v *domain.Vehicle) error {
	f.calls++
	if !f.accept {
		return errors.New("fake: rejected")
	}
	return nil
}

func newTestFleet(center service.ServiceCenter) *Fleet {
	return NewFleet(service.NewInventoryNumberGenerator("TEST"), center)
}

func TestVehicleAcceptedAndGetsInventoryNumber(t *testing.T) {
	center := &fakeInspection{accept: true}
	fleet := newTestFleet(center)
	v, _ := domain.NewVehicle("Самокат", "scooter", 8, 1.5)
	if err := fleet.AddVehicle(v); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.InventoryNumber() != "TEST-0001" {
		t.Fatalf("inventory number = %q", v.InventoryNumber())
	}
	if fleet.VehicleCount() != 1 || fleet.InventoryCount() != 1 {
		t.Fatal("counts were not updated")
	}
	if center.calls != 1 {
		t.Fatalf("inspection calls = %d", center.calls)
	}
}

func TestRejectedVehicleIsNotAdded(t *testing.T) {
	center := &fakeInspection{accept: false}
	fleet := newTestFleet(center)
	v, _ := domain.NewVehicle("Плохой самокат", "scooter", 8, 1)
	if err := fleet.AddVehicle(v); err == nil {
		t.Fatal("expected rejection")
	}
	if fleet.VehicleCount() != 0 || fleet.InventoryCount() != 0 {
		t.Fatal("rejected vehicle was added")
	}
	if v.InventoryNumber() != "" {
		t.Fatal("rejected vehicle received inventory number")
	}
}

func TestBeginnerRule(t *testing.T) {
	fleet := newTestFleet(&fakeInspection{accept: true})
	easy, _ := domain.NewVehicle("Простой", "scooter", 6, 1)
	hard, _ := domain.NewVehicle("Сложный", "e-bike", 5, 2)
	if err := fleet.AddVehicle(easy); err != nil {
		t.Fatal(err)
	}
	if err := fleet.AddVehicle(hard); err != nil {
		t.Fatal(err)
	}
	got := fleet.BeginnerFriendlyVehicles()
	if len(got) != 1 || got[0].Name() != "Простой" {
		t.Fatalf("unexpected beginner list: %#v", got)
	}
}

func TestTotalEnergyIncludesVehiclesAndEquipment(t *testing.T) {
	fleet := newTestFleet(&fakeInspection{accept: true})
	v, _ := domain.NewVehicle("Самокат", "scooter", 7, 1.25)
	e, _ := domain.NewEquipment("Зарядный шкаф", 3.5)
	if err := fleet.AddVehicle(v); err != nil {
		t.Fatal(err)
	}
	if err := fleet.AddItem(e); err != nil {
		t.Fatal(err)
	}
	if got := fleet.TotalDailyEnergy(); got != 4.75 {
		t.Fatalf("energy = %v, want 4.75", got)
	}
}

func TestItemGetsInventoryNumber(t *testing.T) {
	fleet := newTestFleet(&fakeInspection{accept: true})
	helmet, _ := domain.NewHelmet("Шлем", 0)
	station, _ := domain.NewChargingStation("Зарядная станция", 3.5)
	if err := fleet.AddItem(helmet); err != nil {
		t.Fatal(err)
	}
	if err := fleet.AddItem(station); err != nil {
		t.Fatal(err)
	}
	if helmet.InventoryNumber() != "TEST-0001" || station.InventoryNumber() != "TEST-0002" {
		t.Fatalf("numbers = %q / %q", helmet.InventoryNumber(), station.InventoryNumber())
	}
}
