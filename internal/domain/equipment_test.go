package domain

import "testing"

func TestNewEquipmentValidation(t *testing.T) {
	if _, err := NewEquipment("   ", 0); err == nil {
		t.Error("blank name must be rejected")
	}
	if _, err := NewEquipment("Шлем", -1); err == nil {
		t.Error("negative energy must be rejected")
	}
	if _, err := NewHelmet("", 0); err == nil {
		t.Error("helmet with blank name must be rejected")
	}
	if _, err := NewChargingStation("Станция", -2); err == nil {
		t.Error("station with negative energy must be rejected")
	}
}

func TestEquipmentSubtypesWorkAsInventoryItems(t *testing.T) {
	helmet, _ := NewHelmet("Шлем", 0)
	station, _ := NewChargingStation("Зарядная станция", 3.5)
	equipment, _ := NewEquipment("Док-станция", 0.5)

	items := []InventoryItem{helmet, station, equipment}
	for _, item := range items {
		item.SetInventoryNumber("VK-0001")
		if item.InventoryNumber() != "VK-0001" {
			t.Errorf("%T did not accept inventory number", item)
		}
	}
	total := helmet.DailyEnergyConsumption() + station.DailyEnergyConsumption() + equipment.DailyEnergyConsumption()
	if total != 4.0 {
		t.Fatalf("total energy = %v, want 4.0", total)
	}
}
