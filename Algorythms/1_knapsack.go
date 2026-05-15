package main

import (
	"fmt"
	"sort"
)

// Задача 1: Задача о рюкзаке (дробная)
// Жадный алгоритм: берём предметы в порядке убывания удельной ценности (vi/pi)

type Item struct {
	id    int
	weight float64
	value  float64
	ratio  float64 // vi / pi
}

func fractionalKnapsack(items []Item, capacity float64) float64 {
	// Сортируем по убыванию удельной ценности — O(n log n)
	sort.Slice(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	totalValue := 0.0
	remaining := capacity

	for _, item := range items {
		if remaining <= 0 {
			break
		}
		if item.weight <= remaining {
			// Берём предмет целиком
			totalValue += item.value
			remaining -= item.weight
			fmt.Printf("  Предмет %d: берём целиком (вес=%.1f, ценность=%.1f)\n",
				item.id, item.weight, item.value)
		} else {
			// Берём часть предмета
			fraction := remaining / item.weight
			totalValue += item.value * fraction
			fmt.Printf("  Предмет %d: берём долю %.4f (вес=%.1f, ценность=%.2f)\n",
				item.id, fraction, remaining, item.value*fraction)
			remaining = 0
		}
	}

	return totalValue
}

func main() {
	items := []Item{
		{1, 10, 40, 40.0 / 10},
		{2, 20, 90, 90.0 / 20},
		{3, 35, 120, 120.0 / 35},
		{4, 40, 100, 100.0 / 40},
		{5, 50, 80, 80.0 / 50},
	}
	capacity := 135.0

	fmt.Println("=== Задача о рюкзаке (дробная) ===")
	fmt.Printf("Грузоподъёмность: %.1f\n", capacity)
	fmt.Println("Предметы (отсортированные по vi/pi):")
	fmt.Println("Выбор:")

	result := fractionalKnapsack(items, capacity)
	fmt.Printf("Максимальная ценность: %.4f\n", result)
}
