package di

import (
	"vyshkat/internal/app"
	"vyshkat/internal/service"
)

// собранные зависимости
type Container struct {
	Fleet *app.Fleet
}

// New собирает приложение из переданных зависимостей
func New(inspection service.ServiceCenter, numbers *service.InventoryNumberGenerator) *Container {
	return &Container{Fleet: app.NewFleet(numbers, inspection)}
}

// сборка для запуска программы
func NewDefault() *Container {
	return New(service.NewBasicServiceCenter(), service.NewInventoryNumberGenerator("VK"))
}
