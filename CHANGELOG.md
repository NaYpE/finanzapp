# 📦 CHANGELOG

Historial de cambios para FinanzApp.

# Changelog

## [v0.1.0] - 2025-07-16
### Added
- Registro de usuarios con validación de correo y contraseña
- Inicio de sesión con JWT y manejo de sesiones en cookies
- Middleware de autenticación para rutas protegidas
- Middleware que redirige si el usuario logueado accede a rutas públicas
- Mensajes flash con cookies (éxito y error), compatibles con Bootstrap
- Página 404 personalizada
- Página principal con diseño responsivo
- Logout funcional y seguro
- Protección contra acceso no autorizado

### Changed
- Actualización de validaciones de entrada con regex
- Mejora de mensajes de error en login y signup
- Refactorización del controlador `Login` y `Register`

### Fixed
- Corrección en la eliminación de cookies tras logout
- Manejo adecuado de errores al parsear JWT

---

Próxima versión: CRUD de transacciones financieras y visualización con gráficos.