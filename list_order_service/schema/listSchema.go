package schema

import (
	"list_order_service/resolvers" // Importar el archivo de resolvers

	"github.com/graphql-go/graphql"
)

// Definir el tipo de orden para GraphQL
var orderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":                &graphql.Field{Type: graphql.Int},   // ID de la orden
			"id_producto":       &graphql.Field{Type: graphql.Int},   // ID del producto
			"precio_individual": &graphql.Field{Type: graphql.Float}, // Precio individual del producto
			"cantidad":          &graphql.Field{Type: graphql.Int},   // Cantidad del producto
			"precio_total":      &graphql.Field{Type: graphql.Float}, // Precio total proporcionado por el cliente
		},
	},
)

// Definir el Query raíz para listar órdenes
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			// Consulta para obtener todas las órdenes
			"orders": &graphql.Field{
				Type:    graphql.NewList(orderType), // Lista de órdenes
				Resolve: resolvers.GetOrdersResolver,
			},
			// Consulta para obtener una orden específica por ID
			"order": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: resolvers.GetOrderByIDResolver,
			},
			// Campo de salud del servicio
			"health": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return "Order service is running", nil
				},
			},
		},
	},
)

// Crear el esquema de GraphQL solo con Query (sin mutaciones)
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery, // Solo consultas para listar órdenes
})
