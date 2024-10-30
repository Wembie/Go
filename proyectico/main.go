package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct{
	nombre string
	desc string
	completado bool
}

type ListaTareas struct{
	tareas []Tarea
}
//Agregar tarea
func (l *ListaTareas) agregarTarea( t Tarea ){
	l.tareas = append(l.tareas, t)
}

//Marcar commpletado
func (l *ListaTareas) marcarCompletado( index int ){
	l.tareas[index].completado = true
}

//Editar tarea
func (l *ListaTareas) editarTarea( index int, t Tarea ){
	l.tareas[index] = t
}

//Eliminar tarea
func (l *ListaTareas) eliminarTarea( index int ){
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}
func main() {
	//instancia
	lista := ListaTareas{}

	//instacia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for{
		var opcion int
		fmt.Println("Selecciona opcion:\n",
		"1. Agregar Tarea\n",
		"2. Marcar tarea completada\n",
		"3. Editar tarea\n",
		"4. Eliminar tarea\n",
		"5. Salir.")
		fmt.Print("Ingrese opcion: ")
		fmt.Scanln(&opcion)
		switch opcion{
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n') //variable, error
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int 
			fmt.Print("Ingrese el indice de la tarea a completar: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea completada correctamente")
		case 3:
			var index int 
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea a editar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea a editar: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea a editar: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int 
			fmt.Print("Ingrese el indice de la tarea a eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea actualizada correctamente")
		case 5:
			fmt.Println("S-A-L-I-E-N-D-OOOOOOOOOOOOOOOOO")
			return
		default:
			fmt.Println("Opcion invalida")
		}
		//Listar todas las tareas:
		fmt.Println("Lista de tareas:")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++")
		for i, t := range lista.tareas{
			fmt.Printf("%d. Nombre: %s - Descripcion: %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	
}	