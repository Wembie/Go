package handlers

import (
	"database/sql"
	"mysql/models"
	"fmt"
	"log"
)

// ListContacts lista todos los contactos desde la base de datos
func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact;"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("---------------------------------------------------------------------------")
	for rows.Next() {
		// Instancia de modelo contact
		contact := models.Contact{}

		var valueEmail, valuePhone sql.NullString // Para manejar valor null
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &valuePhone)
		if err != nil {
			log.Fatal(err)
		}

		// Verificar si el valor es null o no
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electrónico"
		}
		if valuePhone.Valid {
			contact.Phone = valuePhone.String
		} else {
			contact.Phone = "Sin Telefono"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------------------")
	}
}

// GetContactByID obtiene un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo contact
	contact := models.Contact{}
	var valueEmail, valuePhone sql.NullString // Para manejar valor null

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &valuePhone)
	if err != nil {
		//verificar que exista el registro sino....
		if err == sql.ErrNoRows {
			log.Fatalf("no se encontró ningún contacto con el ID %d", contactID)
		}
	}

	// Verificar si el valor es null o no
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electrónico"
	}
	if valuePhone.Valid {
		contact.Phone = valuePhone.String
	} else {
		contact.Phone = "Sin Telefono"
	}

	fmt.Println("\nLISTA DE UN CONTACTO:")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("---------------------------------------------------------------------------")
}

// CreateContact registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para actualizar un contacto
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con éxito")
}

// DeleteContact elimina un contacto de la base de datos
func DeleteContact(db *sql.DB, contactID int) {
	// Sentencia SQL para eliminar un contacto por su ID
	query := "DELETE FROM contact WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con éxito")
}