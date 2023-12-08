KC=kotlinc
FLAGS = -include-runtime -d
TARGET_SEC = secuencia.jar
TARGET_GR = grafo.jar

all:
	$(KC) secuencia.kt main_sec.kt $(FLAGS) $(TARGET_SEC)
	$(KC) secuencia.kt grafo.kt main_grafo.kt $(FLAGS) $(TARGET_GR)

secuencia:
	$(KC) secuencia.kt main_sec.kt $(FLAGS) $(TARGET_SEC)

grafo:
	$(KC) secuencia.kt grafo.kt main_grafo.kt $(FLAGS) $(TARGET_GR)

clean:
	rm ./{secuencia,grafo}.jar