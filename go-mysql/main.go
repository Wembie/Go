package main

import (
	"log"
	"mysql/database"
	"mysql/handlers"
	"mysql/models"
	"fmt"
	"bufio"
	"strings"
	"os"

	_ "github.com/go-sql-driver/mysql" //la _ se colocar para usar el paquete de forma indirecta y no directa como las demas
	//seria como una dependencia practicamente
)

func main() {
	db, err := database.Connect()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	/*handlers.ListContacts(db)

	handlers.GetContactByID(db, 1)
	//handlers.GetContactByID(db, 69) //no se encuentra

	newContact := models.Contact{
		Name:  "Nuevo Usuario",
		Email: "nuevo@example.com",
		Phone: "123456789",
	}
	handlers.CreateContact(db, newContact)

	handlers.ListContacts(db)

	newEditContact := models.Contact{
		Id: 6,
		Name:  "Wembie",
		Email: "Wembitos@example.com",
		Phone: "987654321",
	}
	handlers.UpdateContact(db, newEditContact)

	handlers.ListContacts(db)
	handlers.DeleteContact(db, 6)
	handlers.ListContacts(db)*/

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		// Ejecutar la opción seleccionada
		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updatedContact := inputContactDetails(option)
			handlers.UpdateContact(db, updatedContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
			handlers.ListContacts(db)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}

}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails( option int) models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4{
		fmt.Print("Ingrese el ID del contacto que quiere editar: ")
		var idContact int
		fmt.Scanln(&idContact)
		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}