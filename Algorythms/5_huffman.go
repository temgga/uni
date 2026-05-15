package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Задача 5: Кодирование Хаффмана
// Жадный алгоритм: на каждом шаге объединяем два узла с наименьшими частотами

// --- Узел дерева Хаффмана ---
type HuffNode struct {
	char  rune
	freq  int
	left  *HuffNode
	right *HuffNode
}

// --- Приоритетная очередь (мин-куча) ---
type PriorityQueue []*HuffNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// --- Построение дерева Хаффмана ---
// Сложность: O(n log n), где n — размер алфавита
func buildHuffmanTree(freq map[rune]int) *HuffNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for ch, f := range freq {
		heap.Push(pq, &HuffNode{char: ch, freq: f})
	}

	for pq.Len() > 1 {
		// Извлекаем два узла с минимальной частотой
		left := heap.Pop(pq).(*HuffNode)
		right := heap.Pop(pq).(*HuffNode)
		// Создаём новый внутренний узел
		parent := &HuffNode{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}
		heap.Push(pq, parent)
	}

	return heap.Pop(pq).(*HuffNode)
}

// --- Обход дерева и генерация кодов ---
func generateCodes(node *HuffNode, prefix string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		// Листовой узел — это символ
		codes[node.char] = prefix
		return
	}
	generateCodes(node.left, prefix+"0", codes)
	generateCodes(node.right, prefix+"1", codes)
}

func main() {
	text := "abracadabra"

	fmt.Println("=== Кодирование Хаффмана ===")
	fmt.Printf("Текст: %s\n\n", text)

	// 1. Подсчёт частот — O(|T|)
	freq := make(map[rune]int)
	for _, ch := range text {
		freq[ch]++
	}

	fmt.Println("Частоты символов:")
	// Для детерминированного вывода сортируем
	chars := make([]rune, 0, len(freq))
	for ch := range freq {
		chars = append(chars, ch)
	}
	sort.Slice(chars, func(i, j int) bool { return freq[chars[i]] > freq[chars[j]] })
	for _, ch := range chars {
		fmt.Printf("  '%c': %d\n", ch, freq[ch])
	}

	// 2. Построение дерева Хаффмана — O(n log n)
	root := buildHuffmanTree(freq)

	// 3. Генерация кодов
	codes := make(map[rune]string)
	generateCodes(root, "", codes)

	fmt.Println("\nКоды Хаффмана:")
	for _, ch := range chars {
		fmt.Printf("  '%c': %s\n", ch, codes[ch])
	}

	// 4. Вычисление средней длины кода
	// Lavg = Σ (fc/|T|) * lc
	total := len(text)
	lavg := 0.0
	for _, ch := range chars {
		pc := float64(freq[ch]) / float64(total)
		lc := float64(len(codes[ch]))
		lavg += pc * lc
	}

	fmt.Printf("\nСредняя длина кода Lavg = %.4f бит/символ\n", lavg)
	fmt.Printf("Длина закодированного текста: %d бит\n", func() int {
		bits := 0
		for _, ch := range text {
			bits += len(codes[ch])
		}
		return bits
	}())
	fmt.Printf("Без сжатия (8 бит/символ): %d бит\n", total*8)
}
