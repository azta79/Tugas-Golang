package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "Tugas-kedua-api/models"
    "Tugas-kedua-api/database"
    "github.com/gorilla/mux"
)

// init inisialisasi koneksi database
func init() {
    database.InitDB()
}

// CreateOrder menangani pembuatan order baru
func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)
    // Membuat order baru dengan memasukkan rekaman ke tabel `orders` dan `items`
    database.Db.Create(&order)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}

// GetOrders mengambil semua pesanan
func GetOrders(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var orders []models.Order
    database.Db.Preload("Items").Find(&orders)
    json.NewEncoder(w).Encode(orders)
}

// GetOrder mengambil pesanan berdasarkan ID
func GetOrder(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    inputOrderID := params["orderId"]

    var order models.Order
    database.Db.Preload("Items").First(&order, inputOrderID)
    json.NewEncoder(w).Encode(order)
}

// UpdateOrder memperbarui pesanan yang ada
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
    var updatedOrder models.Order
    json.NewDecoder(r.Body).Decode(&updatedOrder)
    database.Db.Save(&updatedOrder)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedOrder)
}

// DeleteOrder menghapus pesanan yang ada
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    inputOrderID := params["orderId"]
    // Mengkonversi `orderId` dari string ke uint64
    id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
    // Mengkonversi uint64 menjadi uint
    idToDelete := uint(id64)

    database.Db.Where("order_id = ?", idToDelete).Delete(&models.Item{})
    database.Db.Where("order_id = ?", idToDelete).Delete(&models.Order{})
    w.WriteHeader(http.StatusNoContent)
}
