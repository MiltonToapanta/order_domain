package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	ID               uint    `json:"id" gorm:"primary_key"`
	IDProducto       int     `json:"id_producto"`
	PrecioIndividual float64 `json:"precio_individual"`
	Cantidad         int     `json:"cantidad"`
	PrecioTotal      float64 `json:"precio_total"`
}

// Función para crear una nueva orden en la base de datos
func CreateOrder(db *gorm.DB, order *Order) (*Order, error) {
	// Guardar la orden en la base de datos
	if err := db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
