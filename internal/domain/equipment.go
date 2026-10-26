package domain

import (
	"errors"
	"strings"
)

// Equipment — любая инвентаризируемая вещь: шлем, зарядная станция и т. п.
type Equipment struct {
	name        string
	dailyEnergy float64
	inventory   string
}

func NewEquipment(name string, dailyEnergy float64) (*Equipment, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("название не должно быть пустым")
	}
	if dailyEnergy < 0 {
		return nil, errors.New("расход энергии не может быть отрицательным")
	}
	return &Equipment{name: name, dailyEnergy: dailyEnergy}, nil
}

func (e *Equipment) Name() string                    { return e.name }
func (e *Equipment) DailyEnergyConsumption() float64 { return e.dailyEnergy }
func (e *Equipment) InventoryNumber() string         { return e.inventory }
func (e *Equipment) SetInventoryNumber(number string) {
	e.inventory = number
}

// шлем
type Helmet struct {
	*Equipment
}

func NewHelmet(name string, dailyEnergy float64) (*Helmet, error) {
	base, err := NewEquipment(name, dailyEnergy)
	if err != nil {
		return nil, err
	}
	return &Helmet{Equipment: base}, nil
}

// зарядная станция
type ChargingStation struct {
	*Equipment
}

func NewChargingStation(name string, dailyEnergy float64) (*ChargingStation, error) {
	base, err := NewEquipment(name, dailyEnergy)
	if err != nil {
		return nil, err
	}
	return &ChargingStation{Equipment: base}, nil
}
