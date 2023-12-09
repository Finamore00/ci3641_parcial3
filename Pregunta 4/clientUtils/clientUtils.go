package clientUtils

import (
	"errors"
	"fmt"
	"strings"
)

/*
Da la bienvenida al programa
*/
func PrintWelcome() {
	fmt.Println("¡Bienvenido al simulador de tablas virtuales de métodos!")
	fmt.Println("Para información sobre comandos disponibles y su uso, ingresar 'help'")
}

/*
Función que recibe un comando de declaración de clase y extrae la información relevante
del mismo (Nombre de la clase a crear, nombre de su superclase en caso de tenerla, y listado
de métodos). En caso de que el comando esté mal formateado o alguno de los argumentos sea
inválido, se retorna error.
*/
func ParseClassCmd(s string) (string, string, []string, error) {
	retVals := struct {
		name       string
		superClass string
		methods    []string
		err        error
	}{}
	words := strings.Split(s, " ")

	if len(words) < 2 {
		retVals.err = errors.New("número de argumentos incorrecto")
		return retVals.name, retVals.superClass, retVals.methods, retVals.err
	}

	retVals.name = words[1]
	if words[2] == ":" {
		retVals.superClass = words[3]
		retVals.methods = words[4:]
	} else {
		retVals.methods = words[2:]
	}

	return retVals.name, retVals.superClass, retVals.methods, retVals.err

}

/*
Función que recibe un string con el comando de descripción de clase ingresado por el
usuario y extrae la información relevante del mismo (Nombre de la clase a mostrar).
En caso de que el comando esté mal formateado o tenga errores retorna error.
*/
func ParseShowClassCmd(s string) (string, error) {
	retVals := struct {
		name string
		err  error
	}{}

	words := strings.Split(s, " ")

	if len(words) != 2 {
		retVals.err = errors.New("numero de argumentos incorrecto")
	}

	retVals.name = words[1]
	return retVals.name, retVals.err

}

/*
Muestra al usuario el uso del programa
*/
func PrintHelp() {
	fmt.Println("COMANDOS")
	fmt.Println("\ti. CLASS: Define una nueva clase")
	fmt.Println("\t\tUso: CLASS <nombre> [: <superclase>] [<metodos>...]")
	fmt.Println("\tii. DESCRIBIR: Muestra la tabla de métodos virtuales de una clase")
	fmt.Println("\t\tUso: Describir <clase>")
	fmt.Println("\tiii. SALIR: Culmina la ejecución del programa")
	fmt.Println("\tiv. help: Muestra el manual de uso")
}

/*
Indica que se cometió algún error en el ingreso del comando de descripcion de clase
*/
func PrintShowCmdErr() {
	fmt.Println("Comando de descripción de clase incorrecto o mal formateado.")
	fmt.Println("Para ver uso del comando, ingresar 'help'")
}

/*
Indica que se cometió algún error en el ingreso del comando de declaración de clase
*/
func PrintClassCmdErr() {
	fmt.Println("Comando de creación de clase incorrecto o mal formateado.")
	fmt.Println("Para ver uso del comando, ingresar 'help'")
}

/*
Indica al usuario que la clase indicada como superclase en el proceso de creación no existe.
*/
func PrintClassNotFoundErr(class string) {
	fmt.Printf("No existe ninguna clase registrada de nombre %v\n", class)
}

/*
Indica al usuario que está intentando declarar una clase ya existente
*/
func PrintRepeatedClassErr(class string) {
	fmt.Printf("Ya existe una clase de nombre %v\n", class)
}

/*
Muestra un error genérico de error en el ingreso del comando
*/
func PrintGenericCmdError() {
	fmt.Println("El comando ingresado es desconocido. Para ver información de comandos, ver 'help'")
}
