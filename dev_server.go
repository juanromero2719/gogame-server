package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Crear instancia de Echo
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configurar CORS para desarrollo
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Configurar rutas (las mismas que en api/index.go)
	setupRoutes(e)

	// Configurar puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Mensaje de bienvenida
	fmt.Printf("🚀 Servidor de desarrollo GoGame iniciado\n")
	fmt.Printf("📡 Servidor ejecutándose en: http://localhost:%s\n", port)
	fmt.Printf("🏥 Health check: http://localhost:%s/api/health\n", port)
	fmt.Printf("🎮 Game status: http://localhost:%s/api/game/status\n", port)
	fmt.Printf("ℹ️  Game info: http://localhost:%s/api/game/info\n", port)
	fmt.Printf("⏹️  Presiona Ctrl+C para detener el servidor\n\n")

	// Iniciar servidor
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func setupRoutes(e *echo.Echo) {
	// Ruta de health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status":      "ok",
			"message":     "GoGame API está funcionando correctamente",
			"environment": "development",
		})
	})

	// Ruta de ejemplo para el juego
	e.GET("/api/game/status", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status":      "running",
			"game":        "GoGame",
			"version":     "1.0.0",
			"environment": "development",
		})
	})

	// Ruta para información del juego
	e.GET("/api/game/info", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"name":        "GoGame",
			"description": "Un juego desarrollado en Go",
			"author":      "Tu nombre",
			"version":     "1.0.0",
			"environment": "development",
		})
	})

	// Ruta raíz para desarrollo
	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, `
<!DOCTYPE html>
<html>
<head>
    <title>GoGame - Desarrollo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        h1 { color: #333; }
        .endpoint { background: #f8f9fa; padding: 15px; margin: 10px 0; border-radius: 5px; border-left: 4px solid #007bff; }
        .method { color: #28a745; font-weight: bold; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎮 GoGame - Servidor de Desarrollo</h1>
        <p>Tu API está funcionando correctamente en modo desarrollo.</p>
        
        <h2>📡 Endpoints disponibles:</h2>
        
        <div class="endpoint">
            <div class="method">GET</div>
            <strong><a href="/api/health">/api/health</a></strong>
            <p>Verificar el estado de la API</p>
        </div>
        
        <div class="endpoint">
            <div class="method">GET</div>
            <strong><a href="/api/game/status">/api/game/status</a></strong>
            <p>Estado actual del juego</p>
        </div>
        
        <div class="endpoint">
            <div class="method">GET</div>
            <strong><a href="/api/game/info">/api/game/info</a></strong>
            <p>Información del juego</p>
        </div>
        
        <h2>🚀 Despliegue en Vercel</h2>
        <p>Una vez que hagas push a tu repositorio, tu API estará disponible en Vercel con las mismas rutas.</p>
    </div>
</body>
</html>
		`)
	})
}
