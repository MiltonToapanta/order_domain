package resolvers

import (
	"list_order_service/config"
	"list_order_service/models"

	"github.com/graphql-go/graphql"
)

// Resolver para obtener todas las órdenes
func GetOrdersResolver(params graphql.ResolveParams) (interface{}, error) {
	var orders []models.Order

	// Obtener todas las órdenes de la base de datos
	if err := config.DB.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// Resolver para obtener una orden por ID
func GetOrderByIDResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	var order models.Order

	// Buscar la orden por ID
	if err := config.DB.First(&order, id).Error; err != nil {
		return nil, err
	}

	return order, nil
}
