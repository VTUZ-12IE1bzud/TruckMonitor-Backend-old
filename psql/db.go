package psql

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"io/ioutil"
	"TruckMonitor-Backend/env"
)

type DB struct {
	DataBase *sql.DB
}

/** Подключение к БД. */
func Connection() *DB {
	// Конфигурация.
	var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBDatabase)

	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Ошибка при открытии БД -->%v\n", err)
		panic("Database error")
	}
	return &DB{connection}
}

/** Закрытие БД. */
func (db *DB) Close() {
	err := db.DataBase.Close()
	if err != nil {
		fmt.Printf("Ошибка при закрытии БД -->%v\n", err)
		panic("Database error")
	}
}

/** Инициализация данных. */
func (db *DB) SchemeInitialization() {
	_, err := db.DataBase.Exec(getInitializationData())
	if err != nil {
		fmt.Printf("Ошибка при инициализации схемы БД -->%v\n", err)
		panic("Database error")
	}
}

func getInitializationData() string {
	data, err := ioutil.ReadFile("db.sql")
	if err != nil {
		fmt.Println("Оишбка чтения файла db.sql", err)
		panic("Config Error")
	}
	return string(data)
}
