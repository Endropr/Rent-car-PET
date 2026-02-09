package main

import (
	"errors"
	"fmt"
)

type Car struct {
	Name      string  `json:"name"`
	Model     string  `json:"model"`
	Speed     int     `json:"speed"`
	Year      int     `json:"year"`
	PriceHour float64 `json:"price_hour"`
}

type Order struct {
	Owner    string  `json:"owner"`
	RentCar  string  `json:"rent_car"`
	RentHour int     `json:"rent_hour"`
	Total    float64 `json:"total"`
	Date     string  `json:"date"`
}

type Card struct {
	Number   string
	CardDate string
	CardCVV  string
}

func (c Card) Pay(amount float64) error {
	// Проверка валида
	if len(c.Number) != 16 {
		return errors.New("номер карты должен содержать \u001b[37;1m16 цифр\033[0m")
	}
	if len(c.CardDate) != 4 {
		return errors.New("неверный формат даты (\u001b[37;1mнужно 4 цифры\033[0m)")
	}
	if len(c.CardCVV) != 3 {
		return errors.New("неверный CVV (\u001b[37;1mнужно 3 цифры\033[0m)")
	}

	fmt.Printf("💳 \u001b[34;1mОплата на сумму %.2f руб. прошла успешно!\033[0m)\n", amount)
	return nil
}
