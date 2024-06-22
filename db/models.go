package db

import (
	"time"

	"gorm.io/gorm"
)

// Adelanto model
type Adelanto struct {
	gorm.Model
	AdelantID          uint      `gorm:"primaryKey"`
	MontoAdelanto      float64   `gorm:"type:decimal(10,2);not null"`
	Fecha              time.Time `gorm:"type:date;not null"`
	ProyectoEmpleadoID uint      `gorm:"not null"`
}

// DispositivoMovil model
type DispositivoMovil struct {
	gorm.Model
	DispositivoMovilID uint   `gorm:"primaryKey"`
	DispositivoUUID    string `gorm:"size:100;not null"`
}

// Empleado model
type Empleado struct {
	gorm.Model
	EmpleadoID    uint      `gorm:"primaryKey"`
	Nombre        string    `gorm:"size:100;not null"`
	CI            string    `gorm:"size:30;not null"`
	FechaIngreso  time.Time `gorm:"type:date;not null"`
	EstatusEmpleo bool      `gorm:"not null"`
}

// EmpleadoDispositivo model
type EmpleadoDispositivo struct {
	gorm.Model
	EmpleadoDispositivoID uint `gorm:"primaryKey"`
	EmpleadoID            uint `gorm:"not null"`
	DispositivoMovilID    uint `gorm:"not null"`
}

// EmpleadoTarjeta model
type EmpleadoTarjeta struct {
	gorm.Model
	EmpleadoTarjetaID uint `gorm:"primaryKey"`
	EmpleadoID        uint `gorm:"not null"`
	TarjetaID         uint `gorm:"not null"`
}

// Proyecto model
type Proyecto struct {
	gorm.Model
	ProyectoID  uint      `gorm:"primaryKey"`
	Nombre      string    `gorm:"size:100;not null"`
	Activo      bool      `gorm:"not null;default:false"`
	FechaInicio time.Time `gorm:"type:date"`
	FechaFin    time.Time `gorm:"type:date"`
}

// ProyectoEmpleado model
type ProyectoEmpleado struct {
	gorm.Model
	ProyectoEmpleadoID       uint `gorm:"primaryKey"`
	PresenteRFID             bool `gorm:"not null"`
	PresenteDispositivoMovil bool `gorm:"not null"`
	Activo                   bool `gorm:"not null"`
	ProyectoID               uint `gorm:"not null"`
	EmpleadoID               uint `gorm:"not null"`
}

// RegistroEmpleadoDispositivo model
type RegistroEmpleadoDispositivo struct {
	gorm.Model
	RegistroEmpleadoDispositivoID uint      `gorm:"primaryKey"`
	FechaHora                     time.Time `gorm:"type:datetime;not null"`
	Tipo                          bool      `gorm:"not null"`
	EmpleadoID                    uint      `gorm:"not null"`
	JornadaLaboralID              uint      `gorm:"not null"`
}

// RegistroEmpleadoRFID model
type RegistroEmpleadoRFID struct {
	gorm.Model
	RegistroEmpleadoRFIDID uint      `gorm:"primaryKey"`
	FechaHora              time.Time `gorm:"type:datetime;not null"`
	Tipo                   bool      `gorm:"not null"`
	EmpleadoID             uint      `gorm:"not null"`
	JornadaLaboralID       uint      `gorm:"not null"`
}

// Salario model
type Salario struct {
	gorm.Model
	SalarioID    uint    `gorm:"primaryKey"`
	MontoSalario float64 `gorm:"type:decimal(10,2);not null"`
}

// SalarioEmpleado model
type SalarioEmpleado struct {
	gorm.Model
	SalarioEmpleadoID uint      `gorm:"primaryKey"`
	Fecha             time.Time `gorm:"type:date;not null"`
	SalarioID         uint      `gorm:"not null"`
	EmpleadoID        uint      `gorm:"not null"`
	Actual            bool      `gorm:"not null"`
}

// TarjetaRFID model
type Tarjeta struct {
	gorm.Model
	TarjetaID     uint   `gorm:"primaryKey"`
	CodigoTarjeta string `gorm:"size:100;not null"`
}

// JornadaLaboral model
type JornadaLaboral struct {
	gorm.Model
	JornadaLaboralID uint       `gorm:"primaryKey"`
	FechaHoraInicio  time.Time  `gorm:"type:datetime;not null"`
	FechaHoraFin     *time.Time `gorm:"type:datetime"`
	Observaciones    string     `gorm:"type:text"`
	ProyectoID       uint       `gorm:"not null"`
}
