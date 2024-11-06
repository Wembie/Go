package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// la <- se utiliza como lo que se va a devolver al canal
	/*canal := make(chan int)
	canal <- 15
	valor := <- canal*/
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/", //No existe
		"https://graph.microsoft.com",
	}
	//Ya que con esa toma api por api
	/*for _, api := range apis {
		checkApi(api)
	}
	//Congruencia 
	for _, api := range apis {
		go checkApi(api)
	}

	//time.Sleep(2* time.Second)*/
	//CHANNELSSSSS
	ch := make(chan string)
	//go es para hacer runtime
	for _, api := range apis {
		go checkApi(api, ch)
		//fmt.Println(<-ch) //no
	}

	/*fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/

	for i := 0; i < len(apis); i++{
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Tomo %v segundos\n", elapsed.Seconds())
}

/*func checkApi(api string){
	if _, err := http.Get(api); err != nil{
		fmt.Printf("ERROR: !%s esta caido pa!\n", api)
		return
	}
	fmt.Printf("SUCCESS: !%s esta mela pa!\n", api)
}*/

func checkApi(api string, ch chan string){
	if _, err := http.Get(api); err != nil{
		ch <- fmt.Sprintf("ERROR: !%s esta caido pa!\n", api)
		return
	}
	 ch <- fmt.Sprintf("SUCCESS: !%s esta mela pa!\n", api)
}