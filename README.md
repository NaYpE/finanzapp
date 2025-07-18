# FinanzApp

Aplicación web para la gestión de finanzas personales.

## 🚀 Descripción

FinanzApp es una aplicación que permite a los usuarios llevar el control de sus ingresos y egresos, gestionar categorías de gastos, visualizar reportes, y mantener sus datos respaldados.

## 🧩 Tecnologías utilizadas

- Lenguaje: Go (Gin Framework)
- Base de datos: MySQL
- Autenticación: JWT
- ORM: GORM
- Frontend: Angular (futuro)
- Docker
- CI/CD con GitHub Actions (futuro)

## 📂 Estructura del proyecto

```bash
finanzapp/
├── cmd/
├── config/
├── controllers/
├── dto/
├── handlers/
├── middleware/
├── models/
├── repositories/
├── routes/
├── server/
├── services/
├── tests/
├── utils/
├── go.mod
├── .env
└── README.md
```

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

3. Corre el proyecto:
```bash
go run cmd/main.go
```

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
