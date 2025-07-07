package main

import (
	"fmt"
	"list_order_service/config"
	"list_order_service/schema"
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
		Schema:   &schema.Schema, // Usar el esquema solo con queries
		Pretty:   true,
		GraphiQL: true, // Habilitar la interfaz de GraphiQL para pruebas
	})

	http.Handle("/graphql", handler)

	// Levantar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "4002"
	}

	fmt.Println("Servidor GraphQL (Solo Lectura) ejecutándose en puerto", port)
	fmt.Println("GraphiQL disponible en: http://localhost:" + port + "/graphql")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
