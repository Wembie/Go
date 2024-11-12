package database

import (
	"database/sql"
	"log"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Connect()(*sql.DB, error){
	//CARGAR VARIABLES DE ENTORNO DESDEL EL .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Cadena de conexión a la base de datos MySQL
	//usuario:contraseña@tpc(ip:puerto)/nombreDB
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil{
		return nil, err
	}
	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil{
		return nil, err
	}
	log.Println("Conexion a la DB (MySQL) Exitosa")
	return db, nil
}