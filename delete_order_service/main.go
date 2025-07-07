package main

import (
	"delete_order_service/config"
	"delete_order_service/schema"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Conectar a la base de datos
	config.Connect()

	// Crear el servidor GraphQL con el esquema para borrar
	handler := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)

	// Levantar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "4004"
	}

	fmt.Println("🗑️  Servidor GraphQL (Solo DELETE) ejecutándose en puerto", port)
	fmt.Println("⚠️  ADVERTENCIA: Este servicio solo puede BORRAR órdenes")
	fmt.Println("GraphiQL disponible en: http://localhost:" + port + "/graphql")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
