package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	/*t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Mañana uwu")
	} else if hora < 17 {
		fmt.Println("Tarde uwu")
	} else {
		fmt.Println("Noche uwu")
	}*/
	if t := time.Now(); t.Hour() < 12 { //Solo se puede usar t dentro del if
		fmt.Println("Mañana uwu")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde uwu")
	} else {
		fmt.Println("Noche uwu")
	}
	switch h := time.Now(); {
	case h.Hour() < 12:
		fmt.Println("Mañana uwu")
	case h.Hour() < 17:
		fmt.Println("Tarde uwu")
	default:
		fmt.Println("Noche uwu")
	}
	/*os := runtime.GOOS //Que sistema opertivo tiene
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
		//break //en Go no es necesario el break
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> Ni idea que vrgs esta usando [pa]")
	}*/
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
		//break //en Go no es necesario el break
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> Ni idea que vrgs esta usando [pa]")
	}

	/*for { //infinito

	}*/
	var i int
	for i <= 10{
		fmt.Println(i)
		i++
	}
	{
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	}
	{
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	}
	{
	count := 5
	for i := 0; i < count; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	}

}