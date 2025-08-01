# login-service

## 🧾 Descripción
Este microservicio se encarga de la autenticación y registro de usuarios.

**login-service** utilizando Go, Gin y JWT. Arquitectura de microservicios, puede ser integrado con otros servicios.

## 🚀 Funcionalidades actuales

- Registro de usuarios con validaciones robustas
- Inicio de sesión con autenticación basada en JWT
- Cierre de sesión (logout) seguro
- Protección de rutas privadas mediante middleware
- Prevención de acceso a rutas públicas si ya hay sesión activa
- Mensajes flash para errores y acciones exitosas
- Página 404 personalizada
- Backend con Go (Gin), Frontend con HTML + Bootstrap
- Base de datos MySQL
- Contenedores Docker listos para desarrollo y despliegue

## 🛠️ Tecnologías utilizadas

- Go (Gin framework)
- MySQL
- Docker & Docker Compose
- JWT (Autenticación segura)
- Bootstrap 5
- HTML5, CSS3

## 📂 Estructura del proyecto

```bash
login-service/
├── cmd/               # Punto de entrada principal
├── config/            # Conexión a la base de datos
├── controllers/       # Controladores HTTP
├── dto/               # Data Transfer Objects (LoginInput, RegisterInput)
├── middlewares/       # Middlewares (Auth, ErrorHandler)
├── models/            # Modelos de base de datos (User, etc.)
├── routes/            # Rutas públicas y protegidas
├── templates/         # HTML templates con Bootstrap
├── utils/             # Funciones auxiliares (JWT, Flash Messages, etc.)
├── validations/       # Validaciones personalizadas
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Endpoints

### Públicos
- `GET /signup` – Página de registro
- `POST /signup` – Registro de usuario
- `GET /login` – Página de inicio de sesión
- `POST /login` – Inicio de sesión
- `GET /logout` – Cierre de sesión

### Protegidos (requieren token JWT)
- `GET /dashboard` – Dashboard del usuario autenticado

## 🔒 Seguridad

- JWT tokens firmados y almacenados en cookies HttpOnly
- Hashing seguro con `bcrypt` para contraseñas
- Validaciones backend para email y password

## 🛠️ Cómo correr el proyecto

### Variables de entorno esperadas
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_HOST
- DB_PORT
- JWT_SECRET

### Local
1. Clona el repositorio
```bash
git clone https://github.com/<tu-usuario>/login-service.git
cd login-service
```

2. Crea tu archivo `.env` basado en este formato:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=user
DB_PASSWORD=password
DB_NAME=login-service
JWT_SECRET=clave_secreta
PORT=8181
```

3. Corre el proyecto en entorno local:
```bash
go run cmd/main.go
```

### Docker
3.1 Corre el proyecto con docker:
```bash
docker-compose up --build
```
Accede en `http://localhost:8181`

## 📅 Roadmap

- [x] Registro e inicio de sesión con JWT
- [ ] Angular
- [ ] Documentacion
- [ ] Pruebas
- [ ] Pipeline

## 🤝 Contribuciones

Bienvenidas ideas, reportes y PRs! El objetivo es construir una herramienta útil para todos los usuarios.
