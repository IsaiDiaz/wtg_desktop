package db

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./db/wtg.db"), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	if !DB.Migrator().HasTable(&Adelanto{}) {
		//CreateDatabase("./db/wtg.db")
		migrateTables()
	} else {
		fmt.Println("Base de datos ya existe")
	}
}

func migrateTables() {
	if !DB.Migrator().HasTable(&Adelanto{}) {
		DB.AutoMigrate(
			&Adelanto{},
			&DispositivoMovil{},
			&Empleado{},
			&EmpleadoDispositivo{},
			&EmpleadoTarjeta{},
			&Proyecto{},
			&ProyectoEmpleado{},
			&RegistroEmpleadoDispositivo{},
			&RegistroEmpleadoRFID{},
			&Salario{},
			&SalarioEmpleado{},
			&Tarjeta{},
			&JornadaLaboral{},
		)
		fmt.Println("Tablas migradas")
	} else {
		fmt.Println("Tablas ya migradas")
	}
}

func CreateDatabase(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error al crear la base de datos")
		return
	}
	defer db.Close()

	sqlFile, err := os.ReadFile("./db/wtg_db_create.sql")
	if err != nil {
		fmt.Println("Error al leer el archivo sql")
		return
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		fmt.Println("Error al ejecutar el archivo sql")
		return
	}

	fmt.Println("Base de datos creada")
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("Error al cerrar la base de datos")
		return
	}
	sqlDB.Close()
	fmt.Println("Base de datos cerrada")
}
