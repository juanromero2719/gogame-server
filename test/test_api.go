package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// URL base - cambiar por tu URL de Vercel después del deployment
	baseURL := "https://tu-proyecto.vercel.app"
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	}

	fmt.Printf("🧪 Probando API en: %s\n\n", baseURL)

	// Lista de endpoints para probar
	endpoints := []string{
		"/api/health",
		"/api/game/status",
		"/api/game/info",
	}

	// Probar cada endpoint
	for _, endpoint := range endpoints {
		testEndpoint(baseURL, endpoint)
	}
}

func testEndpoint(baseURL, endpoint string) {
	url := baseURL + endpoint

	fmt.Printf("🔍 Probando: %s\n", url)

	// Crear cliente HTTP con timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Hacer request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ Error: %v\n\n", err)
		return
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error leyendo respuesta: %v\n\n", err)
		return
	}

	// Mostrar resultado
	if resp.StatusCode == 200 {
		fmt.Printf("✅ Status: %d\n", resp.StatusCode)
		fmt.Printf("📄 Respuesta: %s\n\n", string(body))
	} else {
		fmt.Printf("❌ Status: %d\n", resp.StatusCode)
		fmt.Printf("📄 Respuesta: %s\n\n", string(body))
	}
}
