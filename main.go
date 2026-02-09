package main

import (
	"fmt"
	"time"
)

func main() {
	// Список машин
	cars := []Car{
		{Name: "\u001b[37;1mЛада\033[0m", Model: "\u001b[37;1mКорыто\033[0m", Speed: 10, Year: 2026, PriceHour: 100},
		{Name: "\u001b[37;1mBMW\033[0m", Model: "\u001b[37;1mM5\033[0m", Speed: 320, Year: 2025, PriceHour: 3400},
		{Name: "\u001b[37;1mTesla\033[0m", Model: "\u001b[37;1mModel 3\033[0m", Speed: 305, Year: 2025, PriceHour: 2800},
		{Name: "\u001b[37;1mЛада\033[0m", Model: "\u001b[37;1mВеста\033[0m", Speed: 180, Year: 2019, PriceHour: 1300},
		{Name: "\u001b[37;1mFord\033[0m", Model: "\u001b[37;1mMustang\033[0m", Speed: 275, Year: 2016, PriceHour: 2200},
	}

	// Вывод авто
	fmt.Println("\u001b[37;1m === ДОБРО ПОЖАЛОВАТЬ В RENT CAR! === \033[0m")
	fmt.Println("🚗 Выберите автомобиль для проката:")
	for i, car := range cars {
		fmt.Printf("[%d] %s %s (%d г.) — %.2f руб/час\n", i, car.Name, car.Model, car.Year, car.PriceHour)
	}

	// Выбор авто юзером
	var selection int
	fmt.Print("\nВведите номер машины для аренды: ")
	fmt.Scan(&selection)

	if selection < 0 || selection >= len(cars) {
		fmt.Println("❌ \u001b[31;1mОшибка: такой машины нет в списке.\033[0m")
		return
	}

	selectedCar := cars[selection]
	fmt.Printf("Вы выбрали: %s %s\n", selectedCar.Name, selectedCar.Model)

	// Оплата проката
	totalAmount := selectedCar.PriceHour // Стоимость за 1 час

	var cardNumber, cardDate, cardCVV string
	for {
		fmt.Printf("\nСумма к оплате: %.2f руб.\n", totalAmount)
		fmt.Print("\u001b[37;1m• Введите 16 цифр номера карты:\033[0m ")
		fmt.Scan(&cardNumber)
		fmt.Print("\u001b[37;1m• Введите срок действия (4 цифры):\033[0m ")
		fmt.Scan(&cardDate)
		fmt.Print("\u001b[37;1m• Введите CVV (3 цифры):\033[0m ")
		fmt.Scan(&cardCVV)

		payment := Card{
			Number:   cardNumber,
			CardDate: cardDate,
			CardCVV:  cardCVV,
		}

		// Вызов метода Pay из models.go
		if err := payment.Pay(totalAmount); err != nil {
			fmt.Println("❌ \u001b[31;1mОшибка оплаты:\033[0m", err)
			continue
		}

		break
	}

	// Cоздание и сохранение заказа
	newOrder := Order{
		Owner:    "User",
		RentCar:  selectedCar.Name + " " + selectedCar.Model,
		RentHour: 1,
		Total:    totalAmount,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}

	ordersToSave := []Order{newOrder}
	SaveOrders(ordersToSave)

	fmt.Println("\n🎉\u001b[33;1mПриятной поездки! Ваш заказ сохранен в\033[0m \u001b[33morders.json\033[0m")
}
