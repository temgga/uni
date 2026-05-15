package main

import (
	"fmt"
	"sort"
)

// Задача 2: Выбор монет
// Жадный алгоритм: на каждом шаге берём монету максимального номинала, не превышающего остаток

func coinChange(denominations []int, sum int) ([]int, int) {
	// Сортируем номиналы по убыванию — O(k log k)
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))

	result := make([]int, 0)
	remaining := sum

	for _, d := range denominations {
		for remaining >= d {
			result = append(result, d)
			remaining -= d
		}
	}

	return result, len(result)
}

func main() {
	denominations := []int{1, 5, 10, 25}
	S := 99

	fmt.Println("=== Выбор монет (жадный алгоритм) ===")
	fmt.Printf("Номиналы: %v\n", denominations)
	fmt.Printf("Целевая сумма S = %d\n\n", S)

	coins, count := coinChange(denominations, S)

	// Подсчёт по номиналам
	freq := make(map[int]int)
	for _, c := range coins {
		freq[c]++
	}
	// Выводим в порядке убывания
	sorted := []int{25, 10, 5, 1}
	for _, d := range sorted {
		if freq[d] > 0 {
			fmt.Printf("  Монета %2d: %d шт.\n", d, freq[d])
		}
	}

	fmt.Printf("\nИтого монет: %d\n", count)
	fmt.Printf("Набранная сумма: %d\n", S)
}
