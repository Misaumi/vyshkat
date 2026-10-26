package service

import "fmt"

type InventoryNumberGenerator struct {
	prefix string
	next   int
}

func NewInventoryNumberGenerator(prefix string) *InventoryNumberGenerator {
	return &InventoryNumberGenerator{prefix: prefix, next: 1}
}

func (g *InventoryNumberGenerator) Next() string {
	number := fmt.Sprintf("%s-%04d", g.prefix, g.next)
	g.next++
	return number
}
