package models

/*
 "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
*/
type Delivery struct {
	Name    string `json:"name"`
	Phone   string `gorm:"primaryKey" json:"phone"` //Единственное поле которое может быть уникальным - поэтому является primary key
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"adress"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
