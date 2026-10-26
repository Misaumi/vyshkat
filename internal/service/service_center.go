package service

import (
	"errors"
	"vyshkat/internal/domain"
)

type ServiceCenter interface {
	Inspect(vehicle *domain.Vehicle) error
}

type BasicServiceCenter struct{}

func NewBasicServiceCenter() *BasicServiceCenter { return &BasicServiceCenter{} }

func (s *BasicServiceCenter) Inspect(vehicle *domain.Vehicle) error {
	if !vehicle.ConditionOK() {
		return errors.New("устройство не прошло техосмотр")
	}
	return nil
}
