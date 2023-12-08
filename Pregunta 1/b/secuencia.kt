/**
Define la interface secuencia, la cual debe tener métodos para agregar y
remover y objetos, y alguna función que indique si la secuencia está 
vacía
*/
public interface Secuencia<T> {
    fun agregar(elem: T)
    fun remover(): T
    fun vacia(): Boolean
}

/*
Clase que implementa una pila e implementa la interfaz secuencia. Los elementos
son agregados y removidos bajo método LIFO.
*/
public class Pila<T>(): Secuencia<T> {
    private var contents: MutableList<T> = mutableListOf()

    override fun agregar(elem: T) {
        contents.add(elem)
    }

    override fun remover(): T {
        return contents.removeLast() //La lista mutable ya lanza excepción si la lista está vacía :)
    }

    override fun vacia(): Boolean = !contents.isNotEmpty()

    override fun toString(): String {
        var str = "["
        for (elem in contents) {
            str += "${elem}, "
        }
        str = str.subSequence(0, str.length-2).toString()
        str += "]"
        return str
    }
}


/*
Clase que implementa una cola e implementa la interfaz secuencia. Los elementos
son agregados y removidos bajo método FIFO.
*/
public class Cola<T>(): Secuencia<T> {
    private var contents: MutableList<T> = mutableListOf()

    override fun agregar(elem: T) {
        contents.add(elem)
    }

    override fun remover(): T {
        return contents.removeFirst()
    }

    override fun vacia(): Boolean = !contents.isNotEmpty()

    override fun toString(): String {
        var str: String = "["
        for (elem in contents) {
            str += "${elem.toString()}, "
        }
        str = str.subSequence(0, str.length-2).toString()
        str += "]"
        return str
    }
}