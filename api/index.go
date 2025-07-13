package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configurar CORS para permitir requests desde cualquier origen
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Rutas de la API
	setupRoutes()
}

func setupRoutes() {
	// Ruta de health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status":  "ok",
			"message": "GoGame API está funcionando correctamente",
		})
	})

	// Ruta de ejemplo para el juego
	e.GET("/api/game/status", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status":  "running",
			"game":    "GoGame",
			"version": "1.0.0",
		})
	})

	// Ruta para información del juego
	e.GET("/api/game/info", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"name":        "GoGame",
			"description": "Un juego desarrollado en Go",
			"author":      "Tu nombre",
			"version":     "1.0.0",
		})
	})
}

// Handler principal para Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Crear un contexto de Echo
	req := &http.Request{
		Method:        r.Method,
		URL:           r.URL,
		Header:        r.Header,
		Body:          r.Body,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
	}

	// Ajustar la URL para que funcione con Echo
	if !strings.HasPrefix(req.URL.Path, "/api") {
		req.URL.Path = "/api" + req.URL.Path
	}

	// Procesar la request con Echo
	e.ServeHTTP(w, req)
}

// Punto de entrada alternativo para desarrollo local
func main() {
	// Esta función no se ejecutará en Vercel, solo para desarrollo local
	e.Logger.Fatal(e.Start(":8080"))
}
