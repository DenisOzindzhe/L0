package db

/*
https://github.com/go-gorm/postgres
Драйвер для подключения postgres
*/
import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
  PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
}), &gorm.Config{})
*/
//NewConnection устанавливает соединение с бд postgres и возвращает базу данных
func NewConnection(dsn string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(dsn)) // Подключение к бд
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Postgres: %v\n", err) //Выход в случае отсуствия подключения к бд
		os.Exit(1)
	}

	return conn //вернуть бд
}
