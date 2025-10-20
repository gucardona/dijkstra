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