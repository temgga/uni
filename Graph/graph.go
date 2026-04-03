package graph

import "fmt"

// объясвляем новый тип Graph он представляет неориентированный граф с использованием списка смежности.
type Graph struct {
	adjList map[int][]int // поле adjList, где ключ - вершина, значение - срез соседей
}

// AddVertex проверяет, существует ли уже вершина в карте, если нет то создает запись с пустым срезом
func (g *Graph) AddVertex(v int) {
	if g.adjList == nil {
		g.adjList = make(map[int][]int)
	}
	if _, exists := g.adjList[v]; !exists {
		g.adjList[v] = []int{}
	}
}

// AddEdge сначала вызывает AddVertex для обеих вершин(чтобы они точно были в графе) затем добавляет
// каждую вершину в список соседей другой, это создает двунаправленную связь
func (g *Graph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// стек представлен обычным спрезом []int, мы кладем начальную вершину в стек
func (g *Graph) dfsIterative(start int, visited map[int]bool) {
	stack := []int{start} // инициализируем стек с начальной вершиной

	for len(stack) > 0 {
		// Извлекаем вершину из конца стека (LIFO)
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Если вершина уже посещена, пропускаем
		if visited[v] {
			continue
		}
		// Помечаем как посещённую
		visited[v] = true

		// Добавляем всех непосещённых соседей в стек
		for _, neighbor := range g.adjList[v] {
			if !visited[neighbor] { //благодаря проверке избегаем повторного добавления отработанных вершин
				stack = append(stack, neighbor)
			}
		}
	}
}

// здесь мы создаем пустую карту
func (g *Graph) CountComponents() int {
	visited := make(map[int]bool)
	count := 0

	// Проходим по всем вершинам графа
	for vertex := range g.adjList {
		if !visited[vertex] {
			// обнаружена новая вершина
			g.dfsIterative(vertex, visited)
			count++
		}
	}
	return count
}

func main() {
	g := &Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(4, 5)
	g.AddVertex(6)

	fmt.Println("Количество компонент связности:", g.CountComponents()) // 3
}
