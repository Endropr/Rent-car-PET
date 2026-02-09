package main

import (
	"encoding/json"
	"os"
)

func SaveCars(cars []Car) {
	data, _ := json.MarshalIndent(cars, "", "  ")
	os.WriteFile("cars.json", data, 0644)
}
func SaveOrders(orders []Order) {
	data, _ := json.MarshalIndent(orders, "", "  ")
	os.WriteFile("orders.json", data, 0644)
}
