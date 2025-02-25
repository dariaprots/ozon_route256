package json_file

import (
	"encoding/json"
	"os"
	"time"
)

type OutputData struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	UserID int64 	`json:"customerID"`
	Orders []Order	`json:"orders"`
}

type Order struct {	
	ID int64 				`json:"orderID"`
	GetTime time.Time		`json:"getTime"`
	GiveTime time.Time		`json:"giveTime"`
	GiveFlag bool			`json:"giveFlag"`
	ReturnTime time.Time	`json:"returnTime"`
	ReturnFlag time.Time	`json:"returnFlag"`
}

type Storage struct {
	Orders map[int64][]Order
	Path string
}


func NewStorage(path string) (*Storage, error) {
	f, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return &Storage{Path: path}, nil
}


func (s *Storage) readDataFromFile() (err error) {
	var file *os.File
	file, err = os.OpenFile(s.Path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s.Orders)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) writeDataFromFile() (err error) {
	var file *os.File
	file, err = os.OpenFile(s.Path, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetOrderFromCourier(orderID, customerID int64, storageInterval uint) (err error, answer string) {
	err = s.readDataFromFile()
	if err != nil {
		return err, ""
	}
	_, ok := s.Orders[customerID]
	if !ok {
		s.Orders[customerID] = make([]Order, 0)
	}

	// проверка, что заказа с таким айдишником нет, иначе ошибка
	// проверка, что срок хранения не в прошлом

	err = s.writeDataFromFile()
	if err != nil {
		return err, ""
	}
	answerMessage := "Заказ принят"
	return nil, answerMessage
}







// t1.Sub(t2).Hours() разница во времени