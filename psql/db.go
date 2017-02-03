package psql

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"TruckMonitor-Backend/config"
	"io/ioutil"
)

type DB struct {
	DataBase *sql.DB
}

/** Подключение к БД. */
func Connection() *DB {
	// Конфигурация.
	conf := config.GetConfiguration()
	var connectionString = fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", conf.DBUrl, conf.DBUser, conf.DBPassword, conf.DBName)

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
