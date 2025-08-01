# FinanzApp

Aplicación web para la gestión de finanzas personales.

## 🧾 Descripción

**FinanzApp** es una aplicación que permite a los usuarios llevar el control de sus ingresos y egresos, gestionar categorías de gastos, visualizar reportes, y mantener sus datos respaldados.

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
finanzapp/
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

## 🔒 Seguridad

- JWT tokens firmados y almacenados en cookies HttpOnly
- Hashing seguro con `bcrypt` para contraseñas
- Validaciones backend para email y password

## 🛠️ Setup inicial

1. Clona el repositorio
```bash
git clone https://github.com/<tu-usuario>/finanzapp.git
cd finanzapp
```

2. Crea tu archivo `.env` basado en este formato:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=user
DB_PASSWORD=password
DB_NAME=finanzapp
JWT_SECRET=clave_secreta
PORT=8181
```

3. Corre el proyecto en entorno local:
```bash
go run cmd/main.go
```

3.1 Corre el proyecto con docker:
```bash
docker-compose up --build
```
Accede en `http://localhost:8181`

## 📅 Roadmap

- [x] Registro e inicio de sesión con JWT
- [ ] Manejo de transacciones (ingresos y egresos)
- [ ] Categorización y reportes
- [ ] Visualización con gráficos
- [ ] Gestión de grupos y conciliación
- [ ] Exportación/importación de datos
- [ ] Modo móvil / PWA

## 🤝 Contribuciones

Bienvenidas ideas, reportes y PRs! El objetivo es construir una herramienta útil para todos los usuarios.
