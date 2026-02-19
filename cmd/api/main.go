package main

import (
	"granja-avicola/internal/handlers"
	"granja-avicola/internal/repository"
	"granja-avicola/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Conexión a MariaDB (ajusta tus credenciales)
	dsn := "admin_user:admin_password@tcp(host.docker.internal:3306)/gestion_avicola?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falló la conexión a la base de datos:", err)
	}

	// 2. Inicializar Capas (Inyección de dependencias)
	loteRepo := repository.NewLoteRepository(db)
	loteService := services.NewLoteService(loteRepo)
	loteHandler := handlers.NewLoteHandler(loteService)

	// 3. Router
	r := gin.Default()

	// Configurar CORS para que ngx-admin pueda conectar
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// 4. Rutas
	api := r.Group("/api/v1")
	{
		api.GET("/lotes", loteHandler.GetAll)
		api.POST("/lotes", loteHandler.Create)
	}

	// 5. Run
	r.Run(":8080")
}