package resolvers

import (
	"update_order_service/config"
	"update_order_service/models"

	"github.com/graphql-go/graphql"
)

// Resolver para actualizar una orden por ID
func UpdateOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	var order models.Order

	// Primero verificar si la orden existe
	if err := config.DB.First(&order, id).Error; err != nil {
		return nil, err
	}

	// Actualizar campos si se proporcionan
	if idProducto, ok := params.Args["id_producto"]; ok {
		order.IDProducto = idProducto.(int)
	}
	if precioIndividual, ok := params.Args["precio_individual"]; ok {
		order.PrecioIndividual = precioIndividual.(float64)
	}
	if cantidad, ok := params.Args["cantidad"]; ok {
		order.Cantidad = cantidad.(int)
	}
	if precioTotal, ok := params.Args["precio_total"]; ok {
		order.PrecioTotal = precioTotal.(float64)
	}

	// Guardar cambios en la base de datos
	if err := config.DB.Save(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

// Resolver para actualizar múltiples órdenes con los mismos valores
func UpdateAllOrdersResolver(params graphql.ResolveParams) (interface{}, error) {
	updates := make(map[string]interface{})

	// Construir el mapa de actualizaciones
	if idProducto, ok := params.Args["id_producto"]; ok {
		updates["id_producto"] = idProducto
	}
	if precioIndividual, ok := params.Args["precio_individual"]; ok {
		updates["precio_individual"] = precioIndividual
	}
	if cantidad, ok := params.Args["cantidad"]; ok {
		updates["cantidad"] = cantidad
	}
	if precioTotal, ok := params.Args["precio_total"]; ok {
		updates["precio_total"] = precioTotal
	}

	// Actualizar todas las órdenes
	result := config.DB.Model(&models.Order{}).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{
		"updated_count": result.RowsAffected,
		"message":       "Orders updated successfully",
	}, nil
}
