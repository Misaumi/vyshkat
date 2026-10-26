package main

import (
	"bufio"
	"strings"
	"testing"

	"vyshkat/internal/app"
	"vyshkat/internal/service"
)

func newTestFleet() *app.Fleet {
	return app.NewFleet(service.NewInventoryNumberGenerator("VK"), service.NewBasicServiceCenter())
}

func readerOf(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}

func TestReadIntRetriesOnInvalidInput(t *testing.T) {
	if got := readInt(readerOf("не число\n8\n"), "Простота: "); got != 8 {
		t.Fatalf("readInt = %d, want 8", got)
	}
}

func TestReadFloatRetriesOnInvalidInput(t *testing.T) {
	if got := readFloat(readerOf("точка\n1.5\n"), "Энергия: "); got != 1.5 {
		t.Fatalf("readFloat = %v, want 1.5", got)
	}
}

func TestReadChoiceRejectsOutOfRange(t *testing.T) {
	if got := readChoice(readerOf("0\n4\n2\n"), "Вид вещи: "); got != 2 {
		t.Fatalf("readChoice = %d, want 2", got)
	}
}

func TestAddVehicle(t *testing.T) {
	fleet := newTestFleet()
	addVehicle(readerOf("Городской самокат\nscooter\n8\n1.5\n"), fleet)
	if fleet.VehicleCount() != 1 {
		t.Fatalf("vehicle count = %d", fleet.VehicleCount())
	}
	if got := fleet.InventoryItems()[0].InventoryNumber(); got != "VK-0001" {
		t.Fatalf("inventory number = %q", got)
	}
}

func TestAddVehicleWithInvalidEase(t *testing.T) {
	fleet := newTestFleet()
	addVehicle(readerOf("Самокат\nscooter\n0\n1.5\n"), fleet)
	if fleet.VehicleCount() != 0 {
		t.Fatal("invalid vehicle was added")
	}
}

func TestAddItem(t *testing.T) {
	fleet := newTestFleet()
	addItem(readerOf("1\nШлем\n0\n"), fleet)
	addItem(readerOf("2\nЗарядная станция\n3.5\n"), fleet)
	addItem(readerOf("3\nДок-станция\n0.5\n"), fleet)
	if fleet.InventoryCount() != 3 {
		t.Fatalf("inventory count = %d", fleet.InventoryCount())
	}
	if got := fleet.InventoryItems()[1].Name(); got != "Зарядная станция" {
		t.Fatalf("item = %q", got)
	}
}

func TestDemonstrate(t *testing.T) {
	fleet := newTestFleet()
	demonstrate(fleet)
	if fleet.VehicleCount() != 1 {
		t.Fatalf("vehicle count = %d, want 1", fleet.VehicleCount())
	}
	if fleet.InventoryCount() != 3 {
		t.Fatalf("inventory count = %d, want 3", fleet.InventoryCount())
	}
	if got := fleet.TotalDailyEnergy(); got != 4.9 {
		t.Fatalf("energy = %v, want 4.9", got)
	}
}
