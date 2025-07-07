package main

import (
	"create_order_service/config"
	"create_order_service/schema"
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
	// Conectar a la base de datos (esto ahora también ejecuta la migración automática)
	config.Connect()
	// Crear el servidor GraphQL con el esquema y habilitar la interfaz gráfica de GraphiQL
	handler := handler.New(&handler.Config{
		Schema:   &schema.Schema, // Usar el esquema con solo la mutación
		Pretty:   true,
		GraphiQL: true, // Habilitar la interfaz de GraphiQL para pruebas
	})
	http.Handle("/graphql", handler)
	// Levantar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Println("Servidor GraphQL ejecutándose en puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
