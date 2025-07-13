# Deployment en Vercel - GoGame

## Configuración realizada

### 1. Estructura del proyecto
- ✅ Creado `/api/index.go` - Punto de entrada para Vercel
- ✅ Configurado `vercel.json` - Configuración de routing y headers
- ✅ Actualizado `go.mod` - Dependencias optimizadas
- ✅ Creado `.vercelignore` - Archivos excluidos del deployment

### 2. Archivos clave

#### `api/index.go`
- Handler principal para funciones serverless (package handler)
- Configuración de CORS
- Rutas de ejemplo del juego
- Compatibilidad con Echo framework

#### `vercel.json`
- Runtime Go 1.x
- Configuración de rewrites
- Headers CORS

## Rutas disponibles

Una vez desplegado, tu API estará disponible en:

### 🏥 Health Check
```
GET https://tu-proyecto.vercel.app/api/health
```

### 🎮 Estado del juego
```
GET https://tu-proyecto.vercel.app/api/game/status
```

### ℹ️ Información del juego
```
GET https://tu-proyecto.vercel.app/api/game/info
```

## Pasos para desplegar

1. **Commit y push de los cambios**
   ```bash
   git add .
   git commit -m "Configuración para Vercel"
   git push origin main
   ```

2. **En Vercel Dashboard**
   - Ve a tu proyecto en Vercel
   - Haz clic en "Deploy" o espera el auto-deploy
   - Vercel detectará automáticamente los cambios

3. **Verificar el deployment**
   - Accede a `https://tu-proyecto.vercel.app/api/health`
   - Deberías ver: `{"status":"ok","message":"GoGame API está funcionando correctamente"}`

## Solución de problemas

### ✅ Error resuelto: "Please change `package main` to `package handler`"
**Problema**: Error durante el build en Vercel
**Solución**: Cambié `package main` a `package handler` en `api/index.go`
**Motivo**: Vercel requiere que las funciones serverless usen `package handler`

### Si sigues viendo 404:
1. Verifica que el archivo `api/index.go` esté en la raíz del proyecto
2. Asegúrate de que `vercel.json` esté configurado correctamente
3. Revisa los logs en Vercel Dashboard

### Si hay errores de build:
1. Verifica que Go 1.21 esté especificado en `go.mod`
2. Ejecuta `go mod tidy` localmente
3. Asegúrate de que todas las dependencias estén correctas
4. **IMPORTANTE**: Vercel requiere `package handler` en lugar de `package main` para funciones serverless

### Para debugging:
- Revisa los logs en Vercel Dashboard > Functions
- Utiliza `vercel dev` para testing local
- Verifica que las rutas empiecen con `/api/`

## Próximos pasos

1. **Agregar más rutas**: Modifica `setupRoutes()` en `api/index.go`
2. **Conectar base de datos**: Agrega variables de entorno en Vercel
3. **Autenticación**: Implementa middleware de auth
4. **Testing**: Crea tests para tus endpoints

## Variables de entorno

Si necesitas variables de entorno:
1. Ve a Vercel Dashboard > Settings > Environment Variables
2. Agrega las variables necesarias
3. Redeploya el proyecto

¡Tu API GoGame ahora debería estar funcionando en Vercel! 🚀 