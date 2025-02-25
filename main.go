package main

import (
	"bufio"
	"fmt"
	"os"

	json_file "github.com/dariaprots/ozon_route256/storage/json"
)

type Storage interface {
	GetOrderFromCourier() (err error)
	ReturnOrderToCourier() (err error)
	GiveOrderToCustomer() (err error)
	GetOrderList() (err error)
	AcceptRefundFromCustomer() (err error)
	GetRefundList() (err error)
}

var (
	storage Storage
	err error
) 

func main() {
	fmt.Println("Чтобы узнать список доступных команд, введите 'help'.")
	fmt.Println("Введите команду:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandText := scanner.Text()

	switch commandText {
	case "help":
		fmt.Println("Список доступных команд:")
		fmt.Println("Команда 'getorder' позволяет принять заказ от курьера.")
		fmt.Println("Команда 'returnorder' позволяет вернуть заказ курьеру.")
		fmt.Println("Команда 'giveorder' позволяет выдать заказ клиенту.")
		fmt.Println("Команда 'getorderlist' позволяет получить список заказов.")
		fmt.Println("Команда 'acceptrefund' позволяет принять возврат от клиента.")
		fmt.Println("Команда 'getrefundlist' позволяет получить список возвратов.")
	case "getorder":
		
		storage, err = json_file.NewStorage("storage/json/data.json")
		

		
	

	case "returnorder":
	
	case "giveorder": 

	case "getorderlist":

	case "acceptrefund":

	case "getrefundlist":
	
	default:
		fmt.Println("Такой команды не существует! \nПопробуйте ввести команду из доступных.")
	}
}