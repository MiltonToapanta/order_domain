package schema

import (
	"delete_order_service/resolvers"

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

// Tipo para la respuesta de borrado masivo
var deleteResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DeleteResponse",
		Fields: graphql.Fields{
			"deleted_count": &graphql.Field{Type: graphql.Int},
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
					return "Delete service is running", nil
				},
			},
		},
	},
)

// Mutaciones para borrar
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Borrar una orden específica por ID
			"deleteOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: resolvers.DeleteOrderResolver,
			},
			// Borrar todas las órdenes
			"deleteAllOrders": &graphql.Field{
				Type:    deleteResponseType,
				Resolve: resolvers.DeleteAllOrdersResolver,
			},
		},
	},
)

// Crear el esquema de GraphQL con mutaciones para borrar
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
