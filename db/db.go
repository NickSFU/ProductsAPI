package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NickSFU/ProductsAPI/models"
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

func GetProducts() []models.Product {
	var arr []models.Product
	result, err := db.Query("SELECT * FROM products;")
	if err != nil {
		log.Fatalf("Ошибка получения массива продуктов: %v", err)
	}
	defer result.Close()
	for result.Next() {
		var p models.Product
		err := result.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.Measure)
		if err != nil {
			log.Fatalf("Ошибка присваивания структуры: %v", err)
		}
		arr = append(arr, p)
	}
	return arr
}
func GetProductByID(id int) models.Product {
	var p models.Product
	result, err := db.Query("SELECT * FROM products WHERE (id = $1);", id)
	if err != nil {
		log.Fatalf("Ошибка получения продукта по id: %v", err)
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.Measure)
		if err != nil {
			log.Fatalf("Ошибка присваивания структуры: %v", err)
		}
	}
	return p
}
