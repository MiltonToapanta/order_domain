package schema

import (
	"create_order_service/resolvers" // Importar el archivo de resolvers

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

// Definir un Query raíz básico (requerido por GraphQL)
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"health": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return "Service is running", nil
				},
			},
		},
	},
)

// Definir la mutación para crear una nueva orden
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: orderType, // El tipo de dato que vamos a devolver
				Args: graphql.FieldConfigArgument{
					"id_producto":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},   // Aceptar ID del producto
					"precio_individual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)}, // Aceptar precio individual
					"cantidad":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},   // Aceptar cantidad
					"precio_total":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)}, // Aceptar precio total desde el cliente
				},
				// Usar el resolver de la mutación
				Resolve: resolvers.CreateOrderResolver,
			},
		},
	},
)

// Crear el esquema de GraphQL con Query y Mutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,    // Ahora incluimos el Query raíz
	Mutation: rootMutation, // La mutación para crear órdenes
})
