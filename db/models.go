package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Serializadores
type CustomTime struct {
	time.Time
}

// Layout for the custom time format
const ctLayout = "2006-01-02"

// UnmarshalJSON parses the JSON string into CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "null" {
		ct.Time = time.Time{}
		return nil
	}

	str = str[1 : len(str)-1]
	t, err := time.Parse(ctLayout, str)
	if err != nil {
		return fmt.Errorf("error parsing time %q: %v", str, err)
	}
	ct.Time = t
	return nil
}

// MarshalJSON formats the CustomTime into JSON string
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(ctLayout) + `"`), nil
}

// Scan implements the Scanner interface for database deserialization
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		ct.Time = v
		return nil
	case string:
		t, err := time.Parse(ctLayout, v)
		if err != nil {
			return err
		}
		ct.Time = t
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}

// Value implements the Valuer interface for database serialization
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time.Format(ctLayout), nil
}

// Adelanto model
type Adelanto struct {
	AdelantID          uint      `gorm:"primaryKey"`
	MontoAdelanto      float64   `gorm:"type:decimal(10,2);not null"`
	Fecha              time.Time `gorm:"type:date;not null"`
	ProyectoEmpleadoID uint      `gorm:"not null"`
}

// DispositivoMovil model
type DispositivoMovil struct {
	DispositivoMovilID uint   `gorm:"primaryKey"`
	DispositivoUUID    string `gorm:"size:100;not null"`
}

// Empleado model
type Empleado struct {
	EmpleadoID      uint            `gorm:"primaryKey" json:"empleado_id"`
	Nombre          string          `gorm:"size:100;not null" json:"nombre"`
	CI              string          `gorm:"size:30;not null" json:"ci"`
	FechaIngreso    CustomTime      `gorm:"type:date;not null" json:"fecha_ingreso"`
	EstatusEmpleo   bool            `gorm:"not null" json:"estatus_empleo"`
	EmpleadoTarjeta EmpleadoTarjeta `gorm:"foreignKey:EmpleadoID;references:EmpleadoID"`
}

// EmpleadoDispositivo model
type EmpleadoDispositivo struct {
	EmpleadoDispositivoID uint `gorm:"primaryKey"`
	EmpleadoID            uint `gorm:"not null"`
	DispositivoMovilID    uint `gorm:"not null"`
}

// EmpleadoTarjeta model
type EmpleadoTarjeta struct {
	EmpleadoTarjetaID uint `gorm:"primaryKey"`
	EmpleadoID        uint `gorm:"not null"`
	TarjetaID         uint `gorm:"not null"`
}

// Proyecto model
type Proyecto struct {
	ProyectoID  uint      `gorm:"primaryKey"`
	Nombre      string    `gorm:"size:100;not null"`
	Activo      bool      `gorm:"not null;default:false"`
	FechaInicio time.Time `gorm:"type:date"`
	FechaFin    time.Time `gorm:"type:date"`
}

// ProyectoEmpleado model
type ProyectoEmpleado struct {
	ProyectoEmpleadoID       uint `gorm:"primaryKey"`
	PresenteRFID             bool `gorm:"not null"`
	PresenteDispositivoMovil bool `gorm:"not null"`
	Activo                   bool `gorm:"not null"`
	ProyectoID               uint `gorm:"not null"`
	EmpleadoID               uint `gorm:"not null"`
}

// RegistroEmpleadoDispositivo model
type RegistroEmpleadoDispositivo struct {
	RegistroEmpleadoDispositivoID uint      `gorm:"primaryKey"`
	FechaHora                     time.Time `gorm:"type:datetime;not null"`
	Tipo                          bool      `gorm:"not null"`
	EmpleadoID                    uint      `gorm:"not null"`
	JornadaLaboralID              uint      `gorm:"not null"`
}

// RegistroEmpleadoRFID model
type RegistroEmpleadoRFID struct {
	RegistroEmpleadoRFIDID uint      `gorm:"primaryKey"`
	FechaHora              time.Time `gorm:"type:datetime;not null"`
	Tipo                   bool      `gorm:"not null"`
	EmpleadoID             uint      `gorm:"not null"`
	JornadaLaboralID       uint      `gorm:"not null"`
}

// Salario model
type Salario struct {
	SalarioID    uint    `gorm:"primaryKey"`
	MontoSalario float64 `gorm:"type:decimal(10,2);not null"`
}

// SalarioEmpleado model
type SalarioEmpleado struct {
	SalarioEmpleadoID uint      `gorm:"primaryKey"`
	Fecha             time.Time `gorm:"type:date;not null"`
	SalarioID         uint      `gorm:"not null"`
	EmpleadoID        uint      `gorm:"not null"`
	Actual            bool      `gorm:"not null"`
}

// TarjetaRFID model
type Tarjeta struct {
	TarjetaID       uint            `gorm:"primaryKey" json:"tarjeta_id"`
	CodigoTarjeta   string          `gorm:"size:100;not null" json:"codigo_tarjeta"`
	EmpleadoTarjeta EmpleadoTarjeta `gorm:"foreignKey:TarjetaID;references:TarjetaID"`
}

// JornadaLaboral model
type JornadaLaboral struct {
	JornadaLaboralID uint       `gorm:"primaryKey"`
	FechaHoraInicio  time.Time  `gorm:"type:datetime;not null"`
	FechaHoraFin     *time.Time `gorm:"type:datetime"`
	Observaciones    string     `gorm:"type:text"`
	ProyectoID       uint       `gorm:"not null"`
}
