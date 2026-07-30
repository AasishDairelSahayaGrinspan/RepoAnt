

# 🐜 RepoAnt

Una herramienta CLI para eliminar repositorios de GitHub: selecciona uno o varios repositorios y elimínalos con una sola confirmación.

## Características

- 🎨 **Interfaz moderna**: Interfaz de terminal elegante con colores en degradado
- 🔍 **Selección interactiva**: Navega y selecciona repositorios usando las teclas de flecha
- 🔒 **Repositorios protegidos**: Configura repositorios que nunca podrán eliminarse
- 🔐 **Almacenamiento seguro de tokens**: El PAT de GitHub se almacena localmente con permisos restringidos
- ⚠️ **Avisos de seguridad**: Advertencias claras y solicitudes de confirmación antes de la eliminación
- 📦 **Eliminación masiva**: Elimina múltiples repositorios a la vez con salvaguardas adicionales

## Instalación

RepoAnt funciona en **Windows**, **Linux** y **macOS**.

### Homebrew (macOS / Linux)

```bash
brew install AasishDairelSahayaGrinspan/repoant/repoant
```

Alternativa (dos pasos):

```bash
brew tap AasishDairelSahayaGrinspan/repoant
brew install repoant
```

### macOS / Linux

```bash
# Compilar desde el código fuente
go mod tidy
go build -o repoant .

# Instalar globalmente (opcional)
sudo mv repoant /usr/local/bin/
```

### Windows

```bash
# Compilar desde el código fuente
go mod tidy
go build -o repoant.exe .

# Agregar a PATH (opcional)
# Mueve repoant.exe a un directorio en tu PATH
```

### Binarios precompilados

Descarga los binarios precompilados desde la página de lanzamientos.

## Sitio web (Página de marketing)

Se incluye un sitio web de marketing estilo repositorio en `website/`.

Páginas incluidas:

- `website/index.html` - Página de destino con contador de estrellas en vivo y galería de simulación de escenarios
- `website/features.html` - Página de análisis detallado de características optimizada para SEO
- `website/documentation.html` - Documentación completa para configuración, comandos, escenarios y despliegue

Guía del mantenedor:

- `docs/website-marketing-and-deployment.md` - Planificación, operaciones y solución de problemas

### Vista previa local

```bash
cd website
python3 -m http.server 8080
```

Luego abre `http://localhost:8080` en tu navegador.

La página obtiene las estrellas de GitHub en vivo desde:

`https://api.github.com/repos/AasishDairelSahayaGrinspan/repoant`

### Recursos multimedia

- Capturas de escenario:
	- `website/assets/scenario-login.svg`
	- `website/assets/scenario-list.svg`
	- `website/assets/scenario-single-delete.svg`
	- `website/assets/repoant-screenshot.svg` (eliminación múltiple)
	- `website/assets/scenario-protected.svg`
	- `website/assets/scenario-token-missing.svg`
- Espacio para GIF: `website/assets/repoant-demo.gif`

Reemplaza `repoant-demo.gif` con un GIF real grabado desde la terminal para marketing en producción.

### Despliegue en GitHub Pages

GitHub Pages está configurado usando `.github/workflows/pages.yml`.

Comportamiento del despliegue:

- Se activa con pushes a `main`
- Publica el directorio `website/`
- Sirve el contenido en `https://aasishdairelsahayagrinspan.github.io/RepoAnt/`

Si este es el primer despliegue, asegúrate de que en Configuración -> Pages del repositorio se utilice `GitHub Actions` como fuente de compilación.

## Uso

### Inicio de sesión

Almacena tu token de acceso personal de GitHub (requiere los permisos `repo` y `delete_repo`):

```bash
repoant login
```

### Listar repositorios

Visualiza todos tus repositorios de GitHub:

```bash
repoant list
```

### Eliminar repositorio (único)

Selecciona y elimina UN repositorio de forma interactiva:

```bash
repoant delete
```

**Navegación:**
- ↑↓ Teclas de flecha para navegar
- Enter para seleccionar
- Ctrl+C para cancelar

### Eliminar múltiples repositorios

Selecciona y elimina múltiples repositorios a la vez:

```bash
repoant delete --multi
# o
repoant delete -m
```

**Navegación:**
- ↑↓ Teclas de flecha para navegar  
- SPACE (Espacio) para alternar selección
- Enter para confirmar
- Ctrl+C para cancelar

⚠️ La eliminación múltiple requiere escribir `DELETE <cantidad>` para confirmar.

### Administrar repositorios protegidos

Ver repositorios protegidos:
```bash
repoant protect
```

Agregar un repositorio a la lista protegida:
```bash
repoant protect add owner/repo
```

Eliminar un repositorio de la lista protegida:
```bash
repoant protect remove owner/repo
```

### Versión

Verifica la versión de la CLI:
```bash
repoant version
```

## Repositorios protegidos

Los repositorios protegidos no aparecerán en la lista de selección para eliminar. Puedes administrarlos con el comando `protect` o editar manualmente `~/.protected-repos`:

```text
# Repositorios protegidos (uno por línea, formato: owner/repo)
myusername/important-repo
myusername/production-app
```

## Token de GitHub

La CLI requiere un token de acceso personal de GitHub con los siguientes permisos:
- `repo` - Control total de repositorios privados
- `delete_repo` - Eliminar repositorios

Crea un token en: https://github.com/settings/tokens

El token se almacena en `~/.repoant-token` con permisos `0600` (lectura solo para ti).

## Estructura del proyecto

```
repoant/
├── main.go                          # Punto de entrada
├── go.mod                           # Definición del módulo Go
├── cmd/
│   ├── root.go                      # Comando raíz
│   ├── login.go                     # Comando de inicio de sesión
│   ├── list.go                      # Comando para listar
│   ├── delete.go                    # Comando para eliminar
│   ├── protect.go                   # Comando para proteger
│   └── version.go                   # Comando de versión
├── internal/
│   ├── config/
│   │   └── config.go                # Almacenamiento de tokens
│   ├── github/
│   │   └── client.go                # Cliente de la API de GitHub
│   ├── protected/
│   │   └── protected.go             # Manejo de repositorios protegidos
│   └── ui/
│       └── ui.go                    # Componentes de UI con colores
└── .protected-repos.example         # Ejemplo de archivo de repositorios protegidos
```

## Autor

por @aasishdairel

## Licencia

MIT

## Compatibilidad multiplataforma

RepoAnt está construido con Go y funciona sin problemas en:

- **macOS** (Intel y Apple Silicon)
- **Linux** (x86_64, ARM64)
- **Windows** (x86_64)

La aplicación utiliza bibliotecas multiplataforma para:
- Interacciones de la interfaz de terminal
- Operaciones del sistema de archivos
- Salida de colores
- Solicitudes HTTP

Todas las funciones funcionan de manera idéntica en todas las plataformas.
