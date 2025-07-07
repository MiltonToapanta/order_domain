package resolvers

import (
	"delete_order_service/config"
	"delete_order_service/models"

	"github.com/graphql-go/graphql"
)

// Resolver para borrar una orden por ID
func DeleteOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	var order models.Order

	// Primero verificar si la orden existe
	if err := config.DB.First(&order, id).Error; err != nil {
		return nil, err
	}

	// Borrar la orden
	if err := config.DB.Delete(&order).Error; err != nil {
		return nil, err
	}

	// Retornar la orden borrada
	return order, nil
}

// Resolver para borrar todas las órdenes
func DeleteAllOrdersResolver(params graphql.ResolveParams) (interface{}, error) {
	var count int64

	// Contar cuántas órdenes hay antes de borrar
	config.DB.Model(&models.Order{}).Count(&count)

	// Borrar todas las órdenes
	if err := config.DB.Delete(&models.Order{}, "1 = 1").Error; err != nil {
		return nil, err
	}

	// Retornar el número de órdenes borradas
	return map[string]interface{}{
		"deleted_count": count,
		"message":       "All orders deleted successfully",
	}, nil
}
