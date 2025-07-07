package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"update_order_service/config"
	"update_order_service/schema"

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

	// Crear el servidor GraphQL con el esquema para actualizar
	handler := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)

	// Levantar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	fmt.Println("✏️  Servidor GraphQL (Solo UPDATE) ejecutándose en puerto", port)
	fmt.Println("⚠️  ADVERTENCIA: Este servicio solo puede ACTUALIZAR órdenes")
	fmt.Println("GraphiQL disponible en: http://localhost:" + port + "/graphql")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
