package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vyshkat/internal/app"
	"vyshkat/internal/di"
	"vyshkat/internal/domain"
)

func main() {
	fleet := di.NewDefault().Fleet

	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		fmt.Print("Выберите действие: ")
		line, _ := reader.ReadString('\n')
		switch strings.TrimSpace(line) {
		case "1":
			addVehicle(reader, fleet)
		case "2":
			addItem(reader, fleet)
		case "3":
			fmt.Printf("Единиц техники: %d\n", fleet.VehicleCount())
			fmt.Printf("Всего инвентаризируемых объектов: %d\n", fleet.InventoryCount())
		case "4":
			fmt.Printf("Суточное потребление парка: %.2f кВт·ч\n", fleet.TotalDailyEnergy())
		case "5":
			fmt.Println("Техника для новичков:")
			printVehicles(fleet.BeginnerFriendlyVehicles())
		case "6":
			fmt.Println("Вся техника и вещи:")
			printInventory(fleet.InventoryItems())
		case "7":
			fmt.Println("Демонстрация сценариев: добавление, отказ техосмотра, отчёты и список для новичков.")
			demonstrate(fleet)
		case "0":
			fmt.Println("До встречи!")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("=== ВышКат: учёт парка ===")
	fmt.Println("1. Добавить транспорт")
	fmt.Println("2. Добавить вещь на баланс (шлем / зарядная станция / прочее)")
	fmt.Println("3. Отчёт по количеству")
	fmt.Println("4. Суммарное суточное потребление энергии")
	fmt.Println("5. Техника для новичков")
	fmt.Println("6. Все техника и вещи (названия и инвентарные номера)")
	fmt.Println("7. Показать демонстрационные сценарии")
	fmt.Println("0. Выход")
}

func addVehicle(reader *bufio.Reader, fleet *app.Fleet) {
	name := ask(reader, "Название транспорта (например, Городской самокат): ")
	kind := ask(reader, "Тип (scooter/bicycle/e-bike): ")
	easy := readInt(reader, "Простота для новичка (1–10): ")
	energy := readFloat(reader, "Суточный расход энергии, кВт·ч (0, если не потребляет): ")
	vehicle, err := domain.NewVehicle(name, kind, easy, energy)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if err := fleet.AddVehicle(vehicle); err != nil {
		fmt.Println("Техника отклонена:", err)
		return
	}
	fmt.Printf("Техника принята. Инвентарный номер: %s\n", vehicle.InventoryNumber())
}

func addItem(reader *bufio.Reader, fleet *app.Fleet) {
	kind := readChoice(reader, "Вид вещи (1 — шлем, 2 — зарядная станция, 3 — прочее): ")
	name := ask(reader, "Название вещи (например, Шлем): ")
	energy := readFloat(reader, "Суточный расход энергии, кВт·ч (0, если не потребляет): ")

	var item domain.InventoryItem
	var err error
	switch kind {
	case 1:
		item, err = domain.NewHelmet(name, energy)
	case 2:
		item, err = domain.NewChargingStation(name, energy)
	default:
		item, err = domain.NewEquipment(name, energy)
	}
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if err := fleet.AddItem(item); err != nil {
		fmt.Println("Вещь не добавлена:", err)
		return
	}
	fmt.Printf("Вещь добавлена. Инвентарный номер: %s\n", item.InventoryNumber())
}

func demonstrate(fleet *app.Fleet) {
	scooter, _ := domain.NewVehicle("Пилотный самокат", "scooter", 8, 1.4)
	if err := fleet.AddVehicle(scooter); err != nil {
		fmt.Println("Самокат отклонён:", err)
	} else {
		fmt.Println("Самокат принят:", scooter.InventoryNumber())
	}
	hardVehicle, _ := domain.NewVehicle("Неисправный самокат", "scooter", 7, 1.8)
	hardVehicle.SetCondition(false)
	if err := fleet.AddVehicle(hardVehicle); err != nil {
		fmt.Println("Ожидаемый отказ техосмотра:", err)
	}
	helmet, _ := domain.NewHelmet("Шлем", 0)
	if err := fleet.AddItem(helmet); err != nil {
		fmt.Println("Шлем не добавлен:", err)
	} else {
		fmt.Println("Добавлена вещь:", helmet.InventoryNumber())
	}
	station, _ := domain.NewChargingStation("Зарядная станция", 3.5)
	if err := fleet.AddItem(station); err != nil {
		fmt.Println("Станция не добавлена:", err)
	} else {
		fmt.Println("Добавлена вещь:", station.InventoryNumber())
	}
	fmt.Printf("Количество техники: %d; объектов на балансе: %d\n", fleet.VehicleCount(), fleet.InventoryCount())
	fmt.Printf("Суточный расход: %.2f кВт·ч\n", fleet.TotalDailyEnergy())
	fmt.Println("Для новичков:")
	printVehicles(fleet.BeginnerFriendlyVehicles())
	fmt.Println("Инвентарный список:")
	printInventory(fleet.InventoryItems())
}

func printVehicles(vehicles []*domain.Vehicle) {
	if len(vehicles) == 0 {
		fmt.Println("(пусто)")
		return
	}
	for _, v := range vehicles {
		fmt.Printf("- %s [%s], простота %d/10, инв. № %s\n", v.Name(), v.Kind(), v.BeginnerEase(), v.InventoryNumber())
	}
}

func printInventory(items []domain.InventoryItem) {
	if len(items) == 0 {
		fmt.Println("(пусто)")
		return
	}
	for _, item := range items {
		fmt.Printf("- %s — инв. № %s\n", item.Name(), item.InventoryNumber())
	}
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}
func readInt(reader *bufio.Reader, prompt string) int {
	for {
		s := ask(reader, prompt)
		var value int
		if _, err := fmt.Sscan(s, &value); err == nil {
			return value
		}
		fmt.Println("Введите целое число.")
	}
}
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		s := ask(reader, prompt)
		var value float64
		if _, err := fmt.Sscan(s, &value); err == nil {
			return value
		}
		fmt.Println("Введите число.")
	}
}
func readChoice(reader *bufio.Reader, prompt string) int {
	for {
		value := readInt(reader, prompt)
		if value >= 1 && value <= 3 {
			return value
		}
		fmt.Println("Введите число от 1 до 3.")
	}
}
