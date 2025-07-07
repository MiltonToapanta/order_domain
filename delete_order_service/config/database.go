package config

import (
	"fmt"
	"log"
	"os"

	"delete_order_service/models" // Importar el modelo

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// Conectar a la base de datos MySQL
func Connect() {
	var err error
	// Aquí usamos las variables de entorno para la conexión a la DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Conectar a la base de datos MySQL
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	fmt.Println("Conectado a la base de datos MySQL")

	// Migraciones automáticas: crea la tabla `order` si no existe
	err = DB.AutoMigrate(&models.Order{}).Error
	if err != nil {
		log.Fatal("Error al realizar migración:", err)
	}
	fmt.Println("Tabla 'orders' creada o actualizada automáticamente")
}
