package main

import (
	"net/http"

	"github.com/NickSFU/ProductsAPI/db"
	"github.com/NickSFU/ProductsAPI/handlers"
)

func main() {
	//fmt.Println("Hello, GitHub!")
	db.Init()
	//config.InsertBaseData()
	//config.DeleteAllData()
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetProductByID(w, r)
		case "POST":
			handlers.PostProduct(w, r)
		case "PUT":
			handlers.PutProduct(w, r)
		case "DELETE":
			handlers.DeleteProduct(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)

}
