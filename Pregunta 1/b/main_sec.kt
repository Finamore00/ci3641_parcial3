//Pequeño cliente de prueba para ambas clases implementadas
fun main() {
    val pila = Pila<Int>()
    for (i in 0 until 10) {
        pila.agregar(i)
    }
    println("Contenidos de la pila: ${pila}")
    for (i in 0 until 5) {
        println("Elemento removido: ${pila.remover()}")
    }
    println("Contenidos de la pila; ${pila}")
    //Vaciamos la pila
    println("Vaciando pila...")
    while (!pila.vacia()) {
        pila.remover()
    }
    println("Indicador de pila vacia: ${pila.vacia()}")
    println("Intentando remover elemento de pila vacía...")
    try {
        pila.remover()
    } catch (e: NoSuchElementException) {
        println("Epa, la pila está vacía compadre ;)")
    }
    println("")

    val cola = Cola<Int>()
    for (i in 0 until 10) {
        cola.agregar(i)
    }
    println("Contenidos de la cola: ${cola}")
    for (i in 0 until 5) {
        println("Elemento removido: ${cola.remover()}")
    }
    println("Contenidos de la cola: ${cola}")
    println("Vaciando cola...")
    while (!cola.vacia()) {
        cola.remover()
    }
    println("Intentando remover elemento de cola vacía...")
    try {
        cola.remover()
    } catch(e: NoSuchElementException) {
        println("Epa, la cola está vacía compadre ;)")
    }

}