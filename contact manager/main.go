package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"` //Para que en el json sea minuscula y ya :v
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//Guardar contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts) //codificar el slice de contacto en un formato json y escribirlo en un archivo
	if err != nil {
		return err
	}
	return nil
}

//Cargar contactos desde un archivo json
func loadContactsFromFile( contacts *[]Contact) error{
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contacts []Contact
	//Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}
	//Crear instancia de bufio
	reader := bufio.NewReader(os.Stdin)
	for {
		//Mostrar opciones usuario
		fmt.Print("--===[ GESTOR DE CONTACTOS ]===--:\n",
			"1. Agregar Contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Eligue una opcion: ")
		var option int 
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcios: ", err)
			return
		}

		switch option{
		case 1:
			//Agregar un contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n') //variable, error
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			//Agrega contacto al slice
			contacts = append(contacts, c)
			//Guarda en el archivo json
			if err := saveContactsToFile(contacts); err != nil{
				fmt.Println("Error al guardar el contacto: ", err)
			}
			fmt.Println("Contacto agregado correctamente")
		case 2:
			//Mostrar todos los contactos
			fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
			for index, contact := range contacts{
				fmt.Printf("%d. Nombre: %sEmail: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		case 3:
			fmt.Println("S-A-L-I-E-N-D-OOOOOOOOOOOOOOOOO")
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}