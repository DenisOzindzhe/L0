package models

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
	{
	  "order_uid": "b563feb7b2b84b6test",
	  "track_number": "WBILMTESTTRACK",
	  "entry": "WBIL",
	  "delivery": {Delivery
	  "payment": Payment,
	  "items": [],
	  "locale": "en",
	  "internal_signature": "",
	  "customer_id": "test",
	  "delivery_service": "meest",
	  "shardkey": "9",
	  "sm_id": 99,
	  "date_created": "2021-11-26T06:22:19Z",
	  "oof_shard": "1"
	}
*/
type Order struct {
	OrderUID          string   `gorm:"primaryKey" json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `gorm:"foreignKey:phone" json:"delivery"`
	Payment           Payment  `gorm:"foreignKey:transaction" json:"payment"`
	Items             []Item   `gorm:"foreignKey:rid" json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"` //DateTime?
	OofShard          string   `json:"oof_shard"`
}

// Create метод структуры заказа Order который добовляет в базу новый заказ
func (order *Order) Create(dbConnection *gorm.DB) {
	err := dbConnection.Create(order).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable push data to db: %v\n", err)
	}
}

// Select метод структуры заказа Order для выбора данных из бд, по id - id сответсвует типу OrderUID из структуры Order
func (order *Order) Select(dbConnection *gorm.DB, id string) error {
	order.OrderUID = id
	err := dbConnection.Preload(clause.Associations).First(&order).Error

	return err
}

/*
func SelectAllOrders(dbConnection *gorm.DB) ([]Order, error) {
	orders := []Order{}
	err := dbConnection.Preload(clause.Associations).Find(&orders).Error

	return orders, err
}
*/
