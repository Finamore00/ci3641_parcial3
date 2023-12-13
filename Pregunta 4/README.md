# Pregunta 4

Implementación en Go de simulador de tablas de métodos virtuales. Pregunta 4 del tercer parcial de Lenguajes de Programación I (CI3641), trimestre Sep-Dic 2023

## Para ejecutar

* Instalar Golang versión 1.12.x o superior
* Para compilar el código fuente a un binario, ejecutar `go build -o <nombre_binario> main/main.go`
* Para simplemente ejecutar el código, ejecutar `go run main/main.go`

## A considerar

* El programa fue implementado, específicamente, con Go versión 1.21.5. En pro de minimizar posibles errores de compatibilidad se insta a usar esta versión.
* Si se desea compilar, asegurarse de que <nombre_binario> no sea igual al nombre de alguno de los subdirectorios (Llamar únicamente a `go build main/main.go` crea por defecto un binario llamado "main", por lo que también debe evitarse compilar el código de esta manera).

## Pruebas unitarias

Para ejecutar las pruebas unitarias creadas, ejecutar el comando `go test ./simulatorFuncs` desde este directorio.