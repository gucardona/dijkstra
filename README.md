# Visualizador do Algoritmo de Dijkstra

Ferramenta de visualização passo a passo do algoritmo de Dijkstra para caminho mais curto que gera um GIF animado com saída detalhada no console.

## Estrutura do Projeto

```
dijkstra-visualizer/
├── cmd/
│   └── main.go
├── internal/
│   ├── graph/
│   │   ├── graph.go
│   │   └── priority_queue.go
│   ├── algorithm/
│   │   └── dijkstra.go
│   └── visualizer/
│       ├── dot.go
│       └── media.go
├── output/
│   └── dijkstra_animation.gif
├── temp_frames/
│   └── step_*.png
├── go.mod
└── README.md
```

## Requisitos

- **Go** 1.25+
- **Graphviz** (comando dot para geração de PNG)
- **ImageMagick** (comando convert/magick para criação de GIF)

### Instalação das Dependências

**macOS:**
```bash
brew install graphviz imagemagick
```

**Ubuntu/Debian:**
```bash
sudo apt-get install graphviz imagemagick
```

**Windows:**
- Baixe o Graphviz em https://graphviz.org/download/
- Baixe o ImageMagick em https://imagemagick.org/script/download.php
- Adicione ambos ao PATH do sistema

## Instalação

```bash
git clone <url-do-repositorio>
cd dijkstra-visualizer
go mod init dijkstra-visualizer
```

## Uso

```bash
go run cmd/main.go
```

O programa irá:
1. Executar o algoritmo de Dijkstra com saída detalhada no console
2. Gerar frames de visualização (arquivos PNG)
3. Criar GIF animado em `output/dijkstra_animation.gif`
4. Manter os frames PNG em `temp_frames/` para visualização individual
5. Remover apenas os arquivos DOT temporários

## Saída

### Saída no Console
O programa exibe:
- Cada passo da execução do algoritmo
- Vértice atual sendo processado
- Arestas sendo exploradas com comparação de distâncias
- Distâncias e caminhos atualizados
- Vértices visitados e não visitados
- Caminhos mais curtos finais com rotas completas

### Arquivos Gerados
- `output/dijkstra_animation.gif` - Animação completa
- `temp_frames/step_000.png`, `step_001.png`, etc. - Frames individuais

## Personalização

### Modificar o Grafo

Edite a função `createSampleGraph()` em `cmd/main.go`:

```go
func createSampleGraph() *graph.Graph {
    g := graph.NewGraph()
    g.AddEdge("A", "B", 7)
    g.AddEdge("B", "C", 10)
    return g
}
```

### Ajustar Velocidade da Animação

Altere o parâmetro de delay em `cmd/main.go`:

```go
visualizer.CreateGIF(tempDir, outputGIF, 800)  // 800ms entre frames
```

### Alterar Vértice Inicial

Modifique a variável `start` em `cmd/main.go`:

```go
start := "A"  // Altere para qualquer vértice do seu grafo
```

## Visualização

### Cores do Grafo
- **Amarelo** (#FFD700): Vértice atual sendo processado
- **Verde Claro** (#90EE90): Vértices já visitados
- **Azul Claro** (#ADD8E6): Vértices não visitados

### Cores das Arestas
- **Laranja Avermelhado** (#FF4500): Aresta sendo explorada
- **Verde** (#32CD32): Aresta no caminho mais curto
- **Preto** (#000000): Arestas normais

### Tabela
Uma tabela dinâmica abaixo do grafo mostra:
- **Vertex**: Nome do vértice
- **Distance**: Distância atual mais curta do início
- **Previous**: Vértice anterior no caminho mais curto
- **Visited**: Se o vértice já foi visitado

## Detalhes do Algoritmo

1. **Inicialização**: Define todas as distâncias como infinito exceto o vértice inicial (0)
2. **Fila de Prioridade**: Sempre processa o vértice com distância mínima
3. **Relaxamento**: Atualiza distâncias através do vértice atual se encontrar caminho mais curto
4. **Término**: Continua até que todos os vértices alcançáveis sejam visitados

## Solução de Problemas

### "dot: command not found"
Instale o Graphviz e certifique-se de que está no PATH.

### "convert: command not found" ou "magick: command not found"
Instale o ImageMagick e certifique-se de que está no PATH.

### Nenhum arquivo PNG gerado
Verifique se o Graphviz está instalado corretamente e acessível pela linha de comando.

### Falha na criação do GIF
Verifique se a instalação do ImageMagick suporta o formato GIF.


## Step-by-step da execução

========================================
   DIJKSTRA'S ALGORITHM VISUALIZATION
========================================
Start vertex: A


╔════════════════════════════════════════╗
║        ALGORITHM STEP-BY-STEP          ║
╚════════════════════════════════════════╝

[Step 0] Initial state
--------------------------------------------------
→ Current vertex: A
  Distance from start: 0
  Unvisited: A, B, C, D, E, F

  Current distances:
    ► A: distance=0, previous=-
      B: distance=∞, previous=-
      C: distance=∞, previous=-
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=∞, previous=-

[Step 1] Visiting vertex A
--------------------------------------------------
→ Current vertex: A
  Distance from start: 0

  Visited: A
  Unvisited: B, C, D, E, F

  Current distances:
    ► A: distance=0, previous=-
      B: distance=∞, previous=-
      C: distance=∞, previous=-
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=∞, previous=-

[Step 2] Updated distance to B
--------------------------------------------------
→ Current vertex: A
  Distance from start: 0

  Exploring edges:
    A → B (weight: 7)
      Old distance to B: 7
      New distance to B: 7
      ✓ Updated! New path through A

  Visited: A
  Unvisited: B, C, D, E, F

  Current distances:
    ► A: distance=0, previous=-
      B: distance=7, previous=A
      C: distance=∞, previous=-
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=∞, previous=-

[Step 3] Updated distance to C
--------------------------------------------------
→ Current vertex: A
  Distance from start: 0

  Exploring edges:
    A → B (weight: 7)
      Old distance to B: 7
      New distance to B: 7
      ✓ Updated! New path through A
    A → C (weight: 9)
      Old distance to C: 9
      New distance to C: 9
      ✓ Updated! New path through A

  Visited: A
  Unvisited: B, C, D, E, F

  Current distances:
    ► A: distance=0, previous=-
      B: distance=7, previous=A
      C: distance=9, previous=A
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=∞, previous=-

[Step 4] Updated distance to F
--------------------------------------------------
→ Current vertex: A
  Distance from start: 0

  Exploring edges:
    A → B (weight: 7)
      Old distance to B: 7
      New distance to B: 7
      ✓ Updated! New path through A
    A → C (weight: 9)
      Old distance to C: 9
      New distance to C: 9
      ✓ Updated! New path through A
    A → F (weight: 14)
      Old distance to F: 14
      New distance to F: 14
      ✓ Updated! New path through A

  Visited: A
  Unvisited: B, C, D, E, F

  Current distances:
    ► A: distance=0, previous=-
      B: distance=7, previous=A
      C: distance=9, previous=A
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=14, previous=A

[Step 5] Visiting vertex B
--------------------------------------------------
→ Current vertex: B
  Distance from start: 7

  Visited: A, B
  Unvisited: C, D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ► B: distance=7, previous=A
      C: distance=9, previous=A
      D: distance=∞, previous=-
      E: distance=∞, previous=-
      F: distance=14, previous=A

[Step 6] Updated distance to D
--------------------------------------------------
→ Current vertex: B
  Distance from start: 7

  Exploring edges:
    B → C (weight: 10)
      Old distance to C: 9
      New distance to C: 17
    B → D (weight: 15)
      Old distance to D: 22
      New distance to D: 22
      ✓ Updated! New path through B

  Visited: A, B
  Unvisited: C, D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ► B: distance=7, previous=A
      C: distance=9, previous=A
      D: distance=22, previous=B
      E: distance=∞, previous=-
      F: distance=14, previous=A

[Step 7] Visiting vertex C
--------------------------------------------------
→ Current vertex: C
  Distance from start: 9

  Visited: A, B, C
  Unvisited: D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ► C: distance=9, previous=A
      D: distance=22, previous=B
      E: distance=∞, previous=-
      F: distance=14, previous=A

[Step 8] Updated distance to D
--------------------------------------------------
→ Current vertex: C
  Distance from start: 9

  Exploring edges:
    C → D (weight: 11)
      Old distance to D: 20
      New distance to D: 20
      ✓ Updated! New path through C

  Visited: A, B, C
  Unvisited: D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ► C: distance=9, previous=A
      D: distance=20, previous=C
      E: distance=∞, previous=-
      F: distance=14, previous=A

[Step 9] Updated distance to F
--------------------------------------------------
→ Current vertex: C
  Distance from start: 9

  Exploring edges:
    C → D (weight: 11)
      Old distance to D: 20
      New distance to D: 20
      ✓ Updated! New path through C
    C → F (weight: 2)
      Old distance to F: 11
      New distance to F: 11
      ✓ Updated! New path through C

  Visited: A, B, C
  Unvisited: D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ► C: distance=9, previous=A
      D: distance=20, previous=C
      E: distance=∞, previous=-
      F: distance=11, previous=C

[Step 10] Visiting vertex F
--------------------------------------------------
→ Current vertex: F
  Distance from start: 11

  Visited: A, B, C, F
  Unvisited: D, E

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ✓ C: distance=9, previous=A
      D: distance=20, previous=C
      E: distance=∞, previous=-
    ► F: distance=11, previous=C

[Step 11] Updated distance to E
--------------------------------------------------
→ Current vertex: F
  Distance from start: 11

  Exploring edges:
    F → E (weight: 9)
      Old distance to E: 20
      New distance to E: 20
      ✓ Updated! New path through F

  Visited: A, B, C, F
  Unvisited: D, E

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ✓ C: distance=9, previous=A
      D: distance=20, previous=C
      E: distance=20, previous=F
    ► F: distance=11, previous=C

[Step 12] Visiting vertex E
--------------------------------------------------
→ Current vertex: E
  Distance from start: 20

  Visited: A, B, C, E, F
  Unvisited: D

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ✓ C: distance=9, previous=A
      D: distance=20, previous=C
    ► E: distance=20, previous=F
    ✓ F: distance=11, previous=C

[Step 13] Visiting vertex D
--------------------------------------------------
→ Current vertex: D
  Distance from start: 20

  Visited: A, B, C, D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ✓ C: distance=9, previous=A
    ► D: distance=20, previous=C
    ✓ E: distance=20, previous=F
    ✓ F: distance=11, previous=C

[Step 14] Algorithm complete
--------------------------------------------------

  Visited: A, B, C, D, E, F

  Current distances:
    ✓ A: distance=0, previous=-
    ✓ B: distance=7, previous=A
    ✓ C: distance=9, previous=A
    ✓ D: distance=20, previous=C
    ✓ E: distance=20, previous=F
    ✓ F: distance=11, previous=C

╔════════════════════════════════════════╗
║           FINAL RESULTS                ║
╚════════════════════════════════════════╝

Shortest paths from start vertex:
==================================================

  A:
    Distance: 0
    Path: A

  B:
    Distance: 7
    Path: A → B

  C:
    Distance: 9
    Path: A → C

  D:
    Distance: 20
    Path: A → C → D

  E:
    Distance: 20
    Path: A → C → F → E

  F:
    Distance: 11
    Path: A → C → F


Generated 15 visualization steps
