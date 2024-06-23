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

type CustomDateTime struct {
	time.Time
}

// Layout for the custom time format
const cdtLayout = "2006-01-02 15:04:05"

// UnmarshalJSON parses the JSON string into CustomTime
func (cdt *CustomDateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "null" {
		cdt.Time = time.Time{}
		return nil
	}

	str = str[1 : len(str)-1]
	t, err := time.Parse(cdtLayout, str)
	if err != nil {
		return fmt.Errorf("error parsing time %q: %v", str, err)
	}
	cdt.Time = t
	return nil
}

// MarshalJSON formats the CustomTime into JSON string
func (cdt CustomDateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cdt.Time.Format(cdtLayout) + `"`), nil
}

// Scan implements the Scanner interface for database deserialization
func (cdt *CustomDateTime) Scan(value interface{}) error {
	if value == nil {
		cdt.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		cdt.Time = v
		return nil
	case string:
		t, err := time.Parse(cdtLayout, v)
		if err != nil {
			return err
		}
		cdt.Time = t
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}

// Value implements the Valuer interface for database serialization
func (cdt CustomDateTime) Value() (driver.Value, error) {
	return cdt.Time.Format(cdtLayout), nil
}

// Adelanto model
type Adelanto struct {
	AdelantID          uint       `gorm:"primaryKey" json:"adelanto_id"`
	MontoAdelanto      float64    `gorm:"type:decimal(10,2);not null" json:"monto_adelanto"`
	Fecha              CustomTime `gorm:"type:date;not null" json:"fecha"`
	ProyectoEmpleadoID uint       `gorm:"not null" json:"proyecto_empleado_id"`
}

// DispositivoMovil model
type DispositivoMovil struct {
	DispositivoMovilID uint   `gorm:"primaryKey" json:"dispositivo_movil_id"`
	DispositivoUUID    string `gorm:"size:100;not null" json:"dispositivo_uuid"`
	EmpleadoID         uint   `gorm:"not null" json:"empleado_id"`
}

// Empleado model
type Empleado struct {
	EmpleadoID                  uint                          `gorm:"primaryKey" json:"empleado_id"`
	Nombre                      string                        `gorm:"size:100;not null" json:"nombre"`
	CI                          string                        `gorm:"size:30;not null" json:"ci"`
	FechaIngreso                CustomTime                    `gorm:"type:date;not null" json:"fecha_ingreso"`
	EstatusEmpleo               bool                          `gorm:"not null;default:true" json:"estatus_empleo"`
	Tarjeta                     Tarjeta                       `gorm:"foreignKey:EmpleadoID" json:"tarjeta"`
	DispositivoMovil            DispositivoMovil              `gorm:"foreignKey:EmpleadoID" json:"dispositivo_movil"`
	ProyectoEmpleado            []ProyectoEmpleado            `gorm:"foreignKey:EmpleadoID" json:"proyecto_empleado"`
	SalarioEmpleado             []SalarioEmpleado             `gorm:"foreignKey:EmpleadoID" json:"salario_empleado"`
	RegistroEmpleadoDispositivo []RegistroEmpleadoDispositivo `gorm:"foreignKey:EmpleadoID" json:"registro_empleado_dispositivo"`
	RegistroEmpleadoRFID        []RegistroEmpleadoRFID        `gorm:"foreignKey:EmpleadoID" json:"registro_empleado_rfid"`
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
	ProyectoID       uint               `gorm:"primaryKey" json:"proyecto_id"`
	Nombre           string             `gorm:"size:100;not null" json:"nombre"`
	Activo           bool               `gorm:"not null;default:false" json:"activo"`
	FechaInicio      CustomTime         `gorm:"type:date" json:"fecha_inicio"`
	FechaFin         CustomTime         `gorm:"type:date" json:"fecha_fin"`
	ProyectoEmpleado []ProyectoEmpleado `gorm:"foreignKey:ProyectoID" json:"proyecto_empleado"`
	JornadaLaboral   []JornadaLaboral   `gorm:"foreignKey:ProyectoID" json:"jornada_laboral"`
}

// ProyectoEmpleado model
type ProyectoEmpleado struct {
	ProyectoEmpleadoID       uint       `gorm:"primaryKey" json:"proyecto_empleado_id"`
	PresenteRFID             bool       `gorm:"not null" json:"presente_rfid"`
	PresenteDispositivoMovil bool       `gorm:"not null" json:"presente_dispositivo_movil"`
	Activo                   bool       `gorm:"not null" json:"activo"`
	ProyectoID               uint       `gorm:"not null" json:"proyecto_id"`
	EmpleadoID               uint       `gorm:"not null" json:"empleado_id"`
	Adelanto                 []Adelanto `gorm:"foreignKey:ProyectoEmpleadoID" json:"adelanto"`
}

// RegistroEmpleadoDispositivo model
type RegistroEmpleadoDispositivo struct {
	RegistroEmpleadoDispositivoID uint           `gorm:"primaryKey" json:"registro_empleado_dispositivo_id"`
	FechaHora                     CustomDateTime `gorm:"type:datetime;not null" json:"fecha_hora"`
	Tipo                          bool           `gorm:"not null" json:"tipo"`
	EmpleadoID                    uint           `gorm:"not null" json:"empleado_id"`
	JornadaLaboralID              uint           `gorm:"not null" json:"jornada_laboral_id"`
}

// RegistroEmpleadoRFID model
type RegistroEmpleadoRFID struct {
	RegistroEmpleadoRFIDID uint      `gorm:"primaryKey" json:"registro_empleado_rfid_id"`
	FechaHora              time.Time `gorm:"type:datetime;not null" json:"fecha_hora"`
	Tipo                   bool      `gorm:"not null" json:"tipo"`
	EmpleadoID             uint      `gorm:"not null" json:"empleado_id"`
	JornadaLaboralID       uint      `gorm:"not null" json:"jornada_laboral_id"`
}

// Salario model
type Salario struct {
	SalarioID       uint              `gorm:"primaryKey" json:"salario_id"`
	MontoSalario    float64           `gorm:"type:decimal(10,2);not null" json:"monto_salario"`
	SalarioEmpleado []SalarioEmpleado `gorm:"foreignKey:SalarioID" json:"salario_empleado"`
}

// SalarioEmpleado model
type SalarioEmpleado struct {
	SalarioEmpleadoID uint       `gorm:"primaryKey" json:"salario_empleado_id"`
	Fecha             CustomTime `gorm:"type:date;not null" json:"fecha"`
	SalarioID         uint       `gorm:"not null" json:"salario_id"`
	EmpleadoID        uint       `gorm:"not null" json:"empleado_id"`
	Actual            bool       `gorm:"not null" json:"actual"`
}

// TarjetaRFID model
type Tarjeta struct {
	TarjetaID     uint   `gorm:"primaryKey" json:"tarjeta_id"`
	CodigoTarjeta string `gorm:"size:100;not null" json:"codigo_tarjeta"`
	EmpleadoID    uint   `gorm:"not null" json:"empleado_id"`
}

// JornadaLaboral model
type JornadaLaboral struct {
	JornadaLaboralID            uint                          `gorm:"primaryKey" json:"jornada_laboral_id"`
	FechaHoraInicio             CustomDateTime                `gorm:"type:datetime;not null" json:"fecha_hora_inicio"`
	FechaHoraFin                *CustomDateTime               `gorm:"type:datetime" json:"fecha_hora_fin"`
	Observaciones               string                        `gorm:"type:text" json:"observaciones"`
	ProyectoID                  uint                          `gorm:"not null" json:"proyecto_id"`
	RegistroEmpleadoDispositivo []RegistroEmpleadoDispositivo `gorm:"foreignKey:JornadaLaboralID" json:"registro_empleado_dispositivo"`
	RegistroEmpleadoRFID        []RegistroEmpleadoRFID        `gorm:"foreignKey:JornadaLaboralID" json:"registro_empleado_rfid"`
}
