package schema

import (
	"update_order_service/resolvers"

	"github.com/graphql-go/graphql"
)

// Definir el tipo de orden para GraphQL
var orderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":                &graphql.Field{Type: graphql.Int},
			"id_producto":       &graphql.Field{Type: graphql.Int},
			"precio_individual": &graphql.Field{Type: graphql.Float},
			"cantidad":          &graphql.Field{Type: graphql.Int},
			"precio_total":      &graphql.Field{Type: graphql.Float},
		},
	},
)

// Tipo para la respuesta de actualización masiva
var updateResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UpdateResponse",
		Fields: graphql.Fields{
			"updated_count": &graphql.Field{Type: graphql.Int},
			"message":       &graphql.Field{Type: graphql.String},
		},
	},
)

// Query básico para verificar el servicio
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"health": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return "Update service is running", nil
				},
			},
		},
	},
)

// Mutaciones para actualizar
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Actualizar una orden específica por ID
			"updateOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"id":                &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"id_producto":       &graphql.ArgumentConfig{Type: graphql.Int},
					"precio_individual": &graphql.ArgumentConfig{Type: graphql.Float},
					"cantidad":          &graphql.ArgumentConfig{Type: graphql.Int},
					"precio_total":      &graphql.ArgumentConfig{Type: graphql.Float},
				},
				Resolve: resolvers.UpdateOrderResolver,
			},
			// Actualizar todas las órdenes con los mismos valores
			"updateAllOrders": &graphql.Field{
				Type: updateResponseType,
				Args: graphql.FieldConfigArgument{
					"id_producto":       &graphql.ArgumentConfig{Type: graphql.Int},
					"precio_individual": &graphql.ArgumentConfig{Type: graphql.Float},
					"cantidad":          &graphql.ArgumentConfig{Type: graphql.Int},
					"precio_total":      &graphql.ArgumentConfig{Type: graphql.Float},
				},
				Resolve: resolvers.UpdateAllOrdersResolver,
			},
		},
	},
)

// Crear el esquema de GraphQL con mutaciones para actualizar
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
