package errors

import "net/http"

var (
	ErrNotFound       = New(http.StatusNotFound, "Recurso no encontrado")
	ErrUnauthorized   = New(http.StatusUnauthorized, "No autorizado")
	ErrForbidden      = New(http.StatusForbidden, "Acceso denegado")
	ErrBadRequest     = New(http.StatusBadRequest, "Solicitud inválida")
	ErrInternalServer = New(http.StatusInternalServerError, "Error interno del servidor")
	ErrConflict       = New(http.StatusConflict, "Conflicto con el estado actual del recurso")
)
