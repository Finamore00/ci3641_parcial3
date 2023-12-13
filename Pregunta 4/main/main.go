/*
Programa cliente para el simulador de tablas de métodos virtuales. Pregunta
4 del 3er Parcial de Lenguajes de Programación I (CI3641), Trimestre Sep-Dic
2023

Autor: Santiago Finamore
Prof. Ricardo Monascal
*/

package main

import (
	"bufio"
	"fmt"
	"methodTableSim/clientUtils"
	"methodTableSim/simulatorFuncs"
	"os"
	"strings"
)

func main() {
	classTable := map[string]*simulatorFuncs.Class{}
	inputReader := bufio.NewReader(os.Stdin)
	exitFlag := false

	clientUtils.PrintWelcome()

	for !exitFlag {
		fmt.Print("Ingrese el comando a ejecutar: ")
		input, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Println("Se produjo un error leyendo la entrada estándar. Abortando...")
			os.Exit(1)
		}

		input = strings.ToLower(input[:len(input)-1]) //Remove trailing \n

		switch {
		case strings.HasPrefix(input, "class "):
			name, superClassName, methods, err := clientUtils.ParseClassCmd(input)
			var superClass *simulatorFuncs.Class = nil

			if err != nil {
				clientUtils.PrintClassCmdErr()
				continue
			}

			_, classCheck := classTable[name]
			if classCheck {
				clientUtils.PrintRepeatedClassErr(name)
				continue
			}

			if superClassName != "" {
				_, superClassCheck := classTable[superClassName]
				if !superClassCheck {
					clientUtils.PrintClassNotFoundErr(superClassName)
					continue
				}

				superClass = classTable[superClassName]
			}

			if !clientUtils.AllDifferent(methods) {
				clientUtils.PrintRepeatedMethodErr()
				continue
			}

			newClass := simulatorFuncs.NewClass(name, superClass, methods)
			classTable[name] = &newClass

		case strings.HasPrefix(input, "describir "):
			name, err := clientUtils.ParseShowClassCmd(input)

			if err != nil {
				clientUtils.PrintShowCmdErr()
			}

			class, classExists := classTable[name]

			if !classExists {
				clientUtils.PrintClassNotFoundErr(name)
				continue
			}

			simulatorFuncs.ShowClass(class)
		case input == "help":
			clientUtils.PrintHelp()
		case input == "salir":
			exitFlag = true
		default:
			clientUtils.PrintGenericCmdError()
		}
	}
}
