package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	connStr := "user=myuser password=mypassword dbname=productsdb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка открытия соединения: %v", err)
	}
	fmt.Println("DB подключен:", db != nil)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL,
					quantity INTEGER,
					unit_cost NUMERIC NOT NULL,
					measure INT NOT NULL
					);`)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы products: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS measures (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL
					);`)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы measures: %v", err)
	}
}

func InsertBaseData() {
	_, err := db.Exec(`INSERT INTO measures (id ,name) VALUES (1, 'kilogramm'), (2, 'liter');`)
	if err != nil {
		log.Fatalf("Ошибка добавления теста: %v", err)
	}

	_, err = db.Exec(`INSERT INTO products (name, quantity, unit_cost ,measure) VALUES ('Tomatoes', 250, 350, 1), ('Coca-Cola', 50, 100, 2);`)
	if err != nil {
		log.Fatalf("Ошибка добавления теста: %v", err)
	}
}

func DeleteAllData() {
	_, err := db.Exec("DELETE FROM products;")
	if err != nil {
		log.Fatalf("Ошибка удаления теста: %v", err)
	}
	_, err = db.Exec("DELETE FROM measures;")
	if err != nil {
		log.Fatalf("Ошибка удаления теста: %v", err)
	}
}

func readTest() {
	result, err := db.Query("SELECT * FROM products",
		"test", 10, 25, "kg",
	)
	if err != nil {
		log.Fatalf("Ошибка добавления теста: %v", err)
	}
	fmt.Println(result)
}
