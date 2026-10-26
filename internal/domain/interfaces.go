package domain

// EnergyConsumer — объект с суточным расходом энергии.
type EnergyConsumer interface {
	DailyEnergyConsumption() float64
}

// InventoryItem — объект на балансе: наименование и инвентарный номер.
type InventoryItem interface {
	Name() string
	InventoryNumber() string
	SetInventoryNumber(number string)
}
