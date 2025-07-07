package resolvers

import (
	"create_order_service/config"
	"create_order_service/models"

	"github.com/graphql-go/graphql"
)

// Resolver para la mutación createOrder
func CreateOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	// Crear la orden usando los parámetros pasados
	order := models.Order{
		IDProducto:       params.Args["id_producto"].(int),
		PrecioIndividual: params.Args["precio_individual"].(float64),
		Cantidad:         params.Args["cantidad"].(int),
		PrecioTotal:      params.Args["precio_total"].(float64), // Usamos el precio total proporcionado por el cliente
	}

	// Crear la orden en la base de datos
	return models.CreateOrder(config.DB, &order)
}
