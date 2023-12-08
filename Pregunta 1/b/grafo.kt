/*Clase que implementa un grafo dirigido*/

class Grafo(val nodeCount: Int) {
    private var adyancencyLs = Array<MutableList<Int>>(nodeCount, {mutableListOf()})

    fun agregarArco(beg: Int, end: Int) {
        adyancencyLs[beg].add(end)
        return
    }

    fun adyacentes(node: Int): Iterable<Int> {
        if (node >= nodeCount || node < 0)  {
            throw NoSuchElementException("No hay un nodo de valor ${node} en el grafo.")
        }
        return adyancencyLs[node].asIterable()
    }
}

abstract class Busqueda(val g: Grafo) {
    private val visited = Array<Boolean>(g.nodeCount, {false})
    abstract val remaining: Secuencia<Int>

    fun buscar(D: Int, H: Int): Int {
        remaining.agregar(D)
        visited[D] = true
        var explored: Int = 0

        while (!remaining.vacia()) {
            val curr = remaining.remover()
            explored++
            //Si se consigue el nodo se retorna el número de nodos visitado
            if (curr == H) {
                return explored
            }
            for (neighbor in g.adyacentes(curr)) {
                if (!visited[neighbor]) {
                    remaining.agregar(neighbor)
                    visited[neighbor] = true
                }
            }
        }

        //Si no se alcanzó el nodo se retorna -1
        return -1
    }
}

class BFS(g: Grafo): Busqueda(g) {
    override val remaining = Cola<Int>()
}

class DFS(g: Grafo): Busqueda(g) {
    override val remaining = Pila<Int>()
}