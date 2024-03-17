package models

import "time"

// Order represents the model for an order
type Order struct {
    OrderID      uint      `json:"orderId" gorm:"primary_key"`
    CustomerName string    `json:"customerName"`
    OrderedAt    time.Time `json:"orderedAt"`
    Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

// Item represents the model for an item in an order
type Item struct {
    LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
    ItemCode    string `json:"itemCode"`
    Description string `json:"description"`
    Quantity    uint   `json:"quantity"`
    OrderID     uint   `json:"-"`
}
