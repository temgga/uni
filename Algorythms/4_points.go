package main

import (
	"fmt"
	"sort"
)

// Задача 4: Покрытие отрезков минимальным количеством точек
// Жадный алгоритм: сортируем отрезки по правому концу,
// ставим точку на правом конце первого непокрытого отрезка

type Segment struct {
	id    int
	left  int
	right int
}

func minPoints(segments []Segment) []int {
	// Сортируем по правому концу — O(n log n)
	sort.Slice(segments, func(i, j int) bool {
		if segments[i].right == segments[j].right {
			return segments[i].left < segments[j].left
		}
		return segments[i].right < segments[j].right
	})

	points := []int{}
	lastPoint := -1 << 30 // -∞

	for _, seg := range segments {
		if lastPoint < seg.left || lastPoint > seg.right {
			// Текущая точка не покрывает этот отрезок — ставим новую
			lastPoint = seg.right
			points = append(points, lastPoint)
		}
	}

	return points
}

func main() {
	segments := []Segment{
		{1, 1, 4},
		{2, 2, 5},
		{3, 5, 6},
		{4, 7, 9},
	}

	fmt.Println("=== Покрытие отрезков минимальным количеством точек ===")
	fmt.Println("Отрезки:")
	for _, s := range segments {
		fmt.Printf("  Отрезок %d: [%d, %d]\n", s.id, s.left, s.right)
	}

	points := minPoints(segments)

	fmt.Printf("\nМинимальное покрывающее множество точек: %v\n", points)
	fmt.Printf("Количество точек: %d\n", len(points))

	// Проверка покрытия
	fmt.Println("\nПроверка:")
	for _, s := range segments {
		covered := false
		for _, p := range points {
			if p >= s.left && p <= s.right {
				covered = true
				fmt.Printf("  Отрезок [%d,%d] покрыт точкой %d\n", s.left, s.right, p)
				break
			}
		}
		if !covered {
			fmt.Printf("  Отрезок [%d,%d] НЕ ПОКРЫТ!\n", s.left, s.right)
		}
	}
}
