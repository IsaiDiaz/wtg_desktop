-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2024-06-21 23:24:22.26

-- tables
-- Table: Adelanto
CREATE TABLE Adelanto (
    adelant_id integer NOT NULL CONSTRAINT Adelanto_pk PRIMARY KEY AUTOINCREMENT,
    monto_adelanto decimal(10,2) NOT NULL,
    fecha date NOT NULL,
    proyecto_empleado_id integer NOT NULL,
    CONSTRAINT Adelanto_Proyecto_empleado FOREIGN KEY (proyecto_empleado_id)
    REFERENCES Proyecto_empleado (proyecto_empleado_id)
);

-- Table: Dispositivo_movil
CREATE TABLE Dispositivo_movil (
    dispositivo_movil_id integer NOT NULL CONSTRAINT Dispositivo_movil_pk PRIMARY KEY AUTOINCREMENT,
    dispositivo_uuid varchar(100) NOT NULL
);

-- Table: Empleado
CREATE TABLE Empleado (
    empleado_id integer NOT NULL CONSTRAINT Empleado_pk PRIMARY KEY AUTOINCREMENT,
    nombre varchar(100) NOT NULL,
    ci varchar(30) NOT NULL,
    fecha_ingreso date NOT NULL,
    estatus_empleo boolean NOT NULL
);

-- Table: Empleado_dispositivo
CREATE TABLE Empleado_dispositivo (
    empleado_dispositivo_id integer NOT NULL CONSTRAINT Empleado_dispositivo_pk PRIMARY KEY AUTOINCREMENT,
    empleado_id integer NOT NULL,
    dispositivo_movil_id integer NOT NULL,
    CONSTRAINT Empleado_dispositivo_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id),
    CONSTRAINT Empleado_dispositivo_Dispositivo_movil FOREIGN KEY (dispositivo_movil_id)
    REFERENCES Dispositivo_movil (dispositivo_movil_id)
);

-- Table: Empleado_tarjeta
CREATE TABLE Empleado_tarjeta (
    empleado_tarjeta_id integer NOT NULL CONSTRAINT Empleado_tarjeta_pk PRIMARY KEY AUTOINCREMENT,
    empleado_id integer NOT NULL,
    tarjeta_id integer NOT NULL,
    CONSTRAINT Empleado_tarjeta_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id),
    CONSTRAINT Empleado_tarjeta_Tarjeta_RFID FOREIGN KEY (tarjeta_id)
    REFERENCES Tarjeta_RFID (tarjeta_id)
);

-- Table: Proyecto
CREATE TABLE Proyecto (
    proyecto_id integer NOT NULL CONSTRAINT Proyecto_pk PRIMARY KEY AUTOINCREMENT,
    nombre varchar(100) NOT NULL,
    activo boolean NOT NULL DEFAULT false,
    fecha_inicio date,
    fecha_fin date
);

-- Table: Proyecto_empleado
CREATE TABLE Proyecto_empleado (
    proyecto_empleado_id integer NOT NULL CONSTRAINT Proyecto_empleado_pk PRIMARY KEY AUTOINCREMENT,
    presente_rfid boolean NOT NULL,
    presente_dispositivo_movil boolean NOT NULL,
    activo boolean NOT NULL,
    proyecto_id integer NOT NULL,
    empleado_id integer NOT NULL,
    CONSTRAINT Proyecto_empleado_Proyecto FOREIGN KEY (proyecto_id)
    REFERENCES Proyecto (proyecto_id),
    CONSTRAINT Proyecto_empleado_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id)
);

-- Table: Registro_empleado_dispositivo
CREATE TABLE Registro_empleado_dispositivo (
    registro_empleado_dispositivo_id integer NOT NULL CONSTRAINT Registro_empleado_dispositivo_pk PRIMARY KEY AUTOINCREMENT,
    fecha_hora datetime NOT NULL,
    tipo boolean NOT NULL,
    empleado_id integer NOT NULL,
    jornada_laboral_id integer NOT NULL,
    CONSTRAINT Registro_empleado_dispositivo_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id),
    CONSTRAINT Registro_empleado_dispositivo_jornada_laboral FOREIGN KEY (jornada_laboral_id)
    REFERENCES jornada_laboral (jornada_laboral_id)
);

-- Table: Registro_empleado_rfid
CREATE TABLE Registro_empleado_rfid (
    registro_empleado_rfid_id integer NOT NULL CONSTRAINT Registro_empleado_rfid_pk PRIMARY KEY AUTOINCREMENT,
    fecha_hora datetime NOT NULL,
    tipo boolean NOT NULL,
    empleado_id integer NOT NULL,
    jornada_laboral_id integer NOT NULL,
    CONSTRAINT Registro_empleado_rfid_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id),
    CONSTRAINT Registro_empleado_rfid_jornada_laboral FOREIGN KEY (jornada_laboral_id)
    REFERENCES jornada_laboral (jornada_laboral_id)
);

-- Table: Salario
CREATE TABLE Salario (
    salario_id integer NOT NULL CONSTRAINT Salario_pk PRIMARY KEY AUTOINCREMENT,
    monto_salario decimal(10,2) NOT NULL
);

-- Table: Salario_empleado
CREATE TABLE Salario_empleado (
    salario_empleado_id integer NOT NULL CONSTRAINT Salario_empleado_pk PRIMARY KEY AUTOINCREMENT,
    fecha date NOT NULL,
    salario_id integer NOT NULL,
    empleado_id integer NOT NULL,
    actual boolean NOT NULL,
    CONSTRAINT Salario_empleado_Salario FOREIGN KEY (salario_id)
    REFERENCES Salario (salario_id),
    CONSTRAINT Salario_empleado_Empleado FOREIGN KEY (empleado_id)
    REFERENCES Empleado (empleado_id)
);

-- Table: Tarjeta_RFID
CREATE TABLE Tarjeta_RFID (
    tarjeta_id integer NOT NULL CONSTRAINT Tarjeta_RFID_pk PRIMARY KEY AUTOINCREMENT,
    codigo_tarjeta varchar(100) NOT NULL
);

-- Table: jornada_laboral
CREATE TABLE jornada_laboral (
    jornada_laboral_id integer NOT NULL CONSTRAINT jornada_laboral_pk PRIMARY KEY AUTOINCREMENT,
    fecha_hora_inicio datetime NOT NULL,
    fecha_hora_fin datetime,
    observaciones text,
    proyecto_id integer NOT NULL,
    CONSTRAINT jornada_laboral_Proyecto FOREIGN KEY (proyecto_id)
    REFERENCES Proyecto (proyecto_id)
);

-- End of file.

