fun main() {
    val g = Grafo(7)
    g.agregarArco(0, 2)
    g.agregarArco(2, 1)
    g.agregarArco(1, 3)
    g.agregarArco(4, 3)
    g.agregarArco(3, 5)
    g.agregarArco(4, 5)
    g.agregarArco(0, 4)

    val bfs = BFS(g)
    val dfs = DFS(g)
    println("Nodos explorados para ir de 0 a 5 (BFS): ${bfs.buscar(0, 5)}")
    println("Nodos explorados para ir a nodo inalcanzable 6 (BFS): ${bfs.buscar(0, 6)}")
    println("Nodos explorados para ir de 0 a 5 (DFS): ${dfs.buscar(0, 5)}")
    println("Nodos explorados para ir a nodo inalcanzable 6 (DFS): ${dfs.buscar(0, 6)}")
}