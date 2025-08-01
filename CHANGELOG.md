# 📦 CHANGELOG

Historial de cambios para login-service.

# Changelog

## [v0.1.0] - 2025-07-16
### Added
- Registro de usuarios con validación de correo y contraseña
- Confirmación de contraseña en registro
- Verificación de existencia de usuario antes de registrarlo
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

### Improved
- Limpieza de errores en controladores
- División por capas (controllers, routes, dto, models, etc.)
- Redireccionamientos inteligentes según sesión activa

### Removed
- Dependencias innecesarias

### To Do
- Documentación inicial del proyecto
- Página para "Olvido Contraseña"
- Mejoar el front/Angular