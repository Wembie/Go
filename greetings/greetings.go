package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hellos( names []string)(map[string]string,error){
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil{
			return nil, err		
		}
		messages[name] = message
	}
	return messages, nil
}

//Saluda y ya que mas quiere
func Hello( name string ) (string , error){
	if name == ""{
		return "", errors.New("nombre vacio")
	}
	/*if name == ""{
		return name, errors.New("nombre vacio")
	}*/
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) //Error para probar test
	return message, nil
}

func randomFormat() string {
	formats := [] string{
		"Hola, %v! Que se dice el/la soci@",
		"%v, Cantalas pues mij@",
		"Oe %v",
	}
	return formats[rand.Intn(len(formats))]
}