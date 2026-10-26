package app

import (
	"errors"

	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

const beginnerEaseThreshold = 6

// транспорт и вещи на балансе
type Fleet struct {
	vehicles   []*domain.Vehicle
	items      []domain.InventoryItem
	inspection service.ServiceCenter
	numbers    *service.InventoryNumberGenerator
}

func NewFleet(numbers *service.InventoryNumberGenerator, inspection service.ServiceCenter) *Fleet {
	return &Fleet{inspection: inspection, numbers: numbers}
}

func (f *Fleet) AddVehicle(vehicle *domain.Vehicle) error {
	if vehicle == nil {
		return errors.New("транспорт не задан")
	}
	if err := f.inspection.Inspect(vehicle); err != nil {
		return err
	}
	vehicle.SetInventoryNumber(f.numbers.Next())
	f.vehicles = append(f.vehicles, vehicle)
	f.items = append(f.items, vehicle)
	return nil
}

func (f *Fleet) AddItem(item domain.InventoryItem) error {
	if item == nil {
		return errors.New("вещь не задана")
	}
	item.SetInventoryNumber(f.numbers.Next())
	f.items = append(f.items, item)
	return nil
}

func (f *Fleet) VehicleCount() int   { return len(f.vehicles) }
func (f *Fleet) InventoryCount() int { return len(f.items) }

func (f *Fleet) TotalDailyEnergy() float64 {
	var total float64
	for _, item := range f.items {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergyConsumption()
		}
	}
	return total
}

func (f *Fleet) BeginnerFriendlyVehicles() []*domain.Vehicle {
	result := make([]*domain.Vehicle, 0)
	for _, vehicle := range f.vehicles {
		if vehicle.BeginnerEase() >= beginnerEaseThreshold {
			result = append(result, vehicle)
		}
	}
	return result
}

func (f *Fleet) InventoryItems() []domain.InventoryItem {
	result := make([]domain.InventoryItem, len(f.items))
	copy(result, f.items)
	return result
}
