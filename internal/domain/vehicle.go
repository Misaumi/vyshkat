package domain

import (
	"errors"
	"strings"
)

type Vehicle struct {
	name        string
	kind        string
	ease        int
	dailyEnergy float64
	inventory   string
	condition   bool
}

func NewVehicle(name, kind string, ease int, dailyEnergy float64) (*Vehicle, error) {
	name = strings.TrimSpace(name)
	kind = strings.TrimSpace(kind)
	if name == "" {
		return nil, errors.New("название не должно быть пустым")
	}
	if kind == "" {
		return nil, errors.New("тип транспорта не должен быть пустым")
	}
	if ease < 1 || ease > 10 {
		return nil, errors.New("простота должна быть от 1 до 10")
	}
	if dailyEnergy < 0 {
		return nil, errors.New("расход энергии не может быть отрицательным")
	}
	return &Vehicle{name: name, kind: kind, ease: ease, dailyEnergy: dailyEnergy, condition: true}, nil
}

func (v *Vehicle) Name() string                     { return v.name }
func (v *Vehicle) Kind() string                     { return v.kind }
func (v *Vehicle) BeginnerEase() int                { return v.ease }
func (v *Vehicle) DailyEnergyConsumption() float64  { return v.dailyEnergy }
func (v *Vehicle) InventoryNumber() string          { return v.inventory }
func (v *Vehicle) SetInventoryNumber(number string) { v.inventory = number }
func (v *Vehicle) SetCondition(ok bool)             { v.condition = ok }
func (v *Vehicle) ConditionOK() bool                { return v.condition }
