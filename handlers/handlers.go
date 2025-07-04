package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NickSFU/ProductsAPI/db"
	"github.com/NickSFU/ProductsAPI/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := db.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetId(w http.ResponseWriter, r *http.Request, entype string) (int, error) {
	path := r.URL.Path
	var prefix string
	switch entype {
	case "product":
		prefix = "/product/"
	case "measure":
		prefix = "/measure/"
	default:
		return 0, fmt.Errorf("Неизвестный тип сущности: %s", entype)
	}
	idStr := strings.TrimPrefix(path, prefix)
	if idStr == "" {
		return -1, nil
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("Неверный формат ID: %s", idStr)
	}
	return id, nil
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(w, r, "product")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	if id == -1 {
		products := db.GetProducts()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		product, err := db.GetProductByID(id)
		if err != nil {
			http.Error(w, "Ошибка получения продукта: "+err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
	// var product models.Product
	// var products []models.Product
	// id, err := GetId(w, r, "product")
	// if err != nil {
	// 	http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// if id == -1 {
	// 	products = db.GetProducts()
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(products)
	// } else {
	// 	product, err = db.GetProductByID(id)
	// 	if err != nil {
	// 		http.Error(w, "Ошибка получения продукта: "+err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(product)
	// }

}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Ошибка парсинга:", http.StatusBadRequest)
		return
	}
	newid := db.PostProduct(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      newid,
		"message": "Продукт успешно создан",
	})
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	id, err := GetId(w, r, "product")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id == -1 {
		http.Error(w, "Ошибка получения ID: пустой id ", http.StatusBadRequest)
		return
	}
	p, err = db.GetProductByID(id)
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Ошибка парсинга:", http.StatusBadRequest)
		return
	}

	db.PutProduct(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Продукт успешно изменен",
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(w, r, "product")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id == -1 {
		http.Error(w, "Ошибка получения ID: пустой id ", http.StatusBadRequest)
		return
	}
	db.DeleteProduct(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Продукт успешно удален",
	})
}

func GetMeasures(w http.ResponseWriter, r *http.Request) {
	m := db.GetMeasures()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func GetMeasureByID(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(w, r, "measure")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	if id == -1 {
		measure := db.GetMeasures()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(measure)
	} else {
		measure, err := db.GetMeasureByID(id)
		if err != nil {
			http.Error(w, "Ошибка получения единицы измерения: "+err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(measure)
	}

}

func PostMeasure(w http.ResponseWriter, r *http.Request) {
	var m models.Measure
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Ошибка парсинга:", http.StatusBadRequest)
		return
	}
	newid := db.PostMeasure(m)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      newid,
		"message": "Единица измерения успешно создана",
	})
}

func PutMeasure(w http.ResponseWriter, r *http.Request) {
	var p models.Measure
	id, err := GetId(w, r, "measure")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id == -1 {
		http.Error(w, "Ошибка получения ID: пустой id ", http.StatusBadRequest)
		return
	}
	p, err = db.GetMeasureByID(id)
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Ошибка парсинга:", http.StatusBadRequest)
		return
	}

	db.PutMeasure(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Единица измерения успешно изменена",
	})
}

func DeleteMeasure(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(w, r, "measure")
	if err != nil {
		http.Error(w, "Ошибка получения ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id == -1 {
		http.Error(w, "Ошибка получения ID: пустой id ", http.StatusBadRequest)
		return
	}
	db.DeleteMeasure(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Единица измерения успешно удалена",
	})
}
