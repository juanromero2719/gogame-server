# 🚀 Desarrollo Local - GoGame

## Para ejecutar tu proyecto en local:

### 1. Ejecutar el servidor de desarrollo
```bash
go run dev_server.go
```

### 2. Abrir en el navegador
- **Página principal**: http://localhost:8080
- **Health check**: http://localhost:8080/api/health
- **Game status**: http://localhost:8080/api/game/status
- **Game info**: http://localhost:8080/api/game/info

### 3. Para detener el servidor
Presiona `Ctrl+C` en la terminal

## Diferencias entre local y producción:

### Local (dev_server.go):
- ✅ Incluye página web con documentación
- ✅ Respuestas incluyen `"environment": "development"`
- ✅ Usa `package main` para poder ejecutarse directamente

### Producción (api/index.go):
- ✅ Optimizado para Vercel serverless
- ✅ Usa `package handler` requerido por Vercel
- ✅ Solo endpoints JSON, sin página web

## Comandos útiles:

### Actualizar dependencias:
```bash
go mod tidy
```

### Compilar para testing:
```bash
go build -o dev_server.exe dev_server.go
./dev_server.exe
```

### Probar endpoints desde terminal:
```bash
# Windows PowerShell
curl http://localhost:8080/api/health

# O usar el archivo de test
cd test
go run test_api.go http://localhost:8080
```

¡Tu servidor local está listo! 🎮 