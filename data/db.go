package data

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

type (
	/** Структура схемы БД. */
	Schema struct {
		Data []string
	}
	/** Инициализированные данные по умолчанию. */
	Default struct {
		Data []string
	}
)

var connection *sql.DB

/** Подключение к БД. */
func ConnectionDB(dbUrl string, dbName string, userName string, userPassword string) {
	// Конфигурация.
	var connectionString = fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbUrl, userName, userPassword, dbName)

	var err error
	connection, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Ошибка при открытии БД -->%v\n", err)
		panic("Database error")
	}
}

/** Закрытие БД. */
func CloseDB() {
	err := connection.Close()
	if err != nil {
		fmt.Printf("Ошибка при закрытии БД -->%v\n", err)
		panic("Database error")
	}
}

func GetDataBaseConnection() *sql.DB {
	return connection
}

/** Инициализация схемы. */
func (schema Schema) Setup() {
	for _, query := range schema.Data {
		_, err := connection.Exec(query)
		if err != nil {
			fmt.Printf("Ошибка при инициализации схемы БД -->%v\n", err)
			panic("Database error")
		}
	}
}

/** Инициализация данных по умолчанию. */
func (defaultData Default) Setup() {
	for _, query := range defaultData.Data {
		_, err := connection.Exec(query)

		if err != nil {
			fmt.Printf("Ошибка при инициализации данных по умолчанию БД -->%v\n", err)
			panic("Database error")
		}
	}
}