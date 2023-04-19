package main

import "fmt"

func main() {
	var costOfGoods int
	var costOfDelivery int
	var discountAmount int

	fmt.Println("Приветствую! Давайте рассчитаем общую стоимость товара.")
	fmt.Println("Для этого ниже потребуется ввести стоимость товара, доставки и размер скидки.")
	fmt.Println("Начнем с стоимости товара:")

	fmt.Scan(&costOfGoods)

	fmt.Println("Теперь введите стоимость доставки:")

	fmt.Scan(&costOfDelivery)

	fmt.Println("Теперь введите размер скидки:")

	fmt.Scan(&discountAmount)

	if costOfGoods >= 10000 {
		fmt.Println("Стоимость товара превышает 10000, поэтому вы получаете скидку 50% на доставку!")
		costOfDelivery /= 2

		if costOfGoods%2 == 0 {
			fmt.Println("Покупателю положен подарок!")
		}
	}

	fmt.Println("Стоимость товара:", costOfGoods)
	fmt.Println("Стоимость доставки:", costOfDelivery)
	fmt.Println("Размер скидки:", discountAmount)
	fmt.Println("---------")

	total := costOfGoods + costOfDelivery - discountAmount

	fmt.Println("Итого:", total)
}
