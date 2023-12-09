package simulatorFuncs

import (
	"fmt"
)

type Class struct {
	Name       string
	Methods    map[string]string
	SuperClass *Class
}

/*
Función que construye una nueva clase y su tabla de métodos virtuales. Recibe
como argumentos el nombre de la clase a crear, una referencia a su superclase
y una lista de strings con los nombres de sus métodos. En el caso de que la
clase a crear no herede de ninguna otra clase, el argumento superClass es nil.
*/
func NewClass(name string, superClass *Class, methods []string) Class {
	methodList := map[string]string{}
	for _, m := range methods {
		methodList[m] = name
	}

	currSuperClass := superClass
	for currSuperClass != nil {
		//Check all methods of the superClass
		for k, v := range currSuperClass.Methods {
			_, ok := methodList[k]
			//If method not already shadowed, add to the class under creation
			if !ok {
				methodList[k] = v
			}
		}
		currSuperClass = currSuperClass.SuperClass //Climb to next superclass in hierarchy
	}

	return Class{
		Name:       name,
		Methods:    methodList,
		SuperClass: superClass,
	}

}

/*
Muestra la tabla virtual de métodos de una clase
*/
func ShowClass(class *Class) {
	for method, supClass := range class.Methods {
		fmt.Printf("%v -> %v :: %v\n", method, supClass, method)
	}
}
