# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/Wembie/greetings
```
## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"log"
	"github.com/Wembie/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //bandera de formado, para no incluir fecha y hora

	names := []string{"Wumpux","Toe","BO6", "Wembie"}
	messages, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	/*message, err := greetings.Hello("Wembie")
	//message, err := greetings.Hello("")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)*/
	fmt.Println(messages)
}

```
Este ejemplo importa el paquete github.com/Wembie/greetings y llama a la función Hello para obtener un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
