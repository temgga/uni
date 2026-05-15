package main

import (
	"fmt"
	"sort"
)

// Задача 3: Расписание (Activity Selection Problem)
// Жадный алгоритм: сортируем занятия по времени окончания,
// выбираем занятие если оно начинается не раньше конца предыдущего

type Activity struct {
	id    int
	start int
	end   int
}

func activitySelection(activities []Activity) []Activity {
	// Сортируем по времени окончания — O(n log n)
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	})

	selected := []Activity{activities[0]}
	lastEnd := activities[0].end

	for i := 1; i < len(activities); i++ {
		a := activities[i]
		if a.start >= lastEnd {
			selected = append(selected, a)
			lastEnd = a.end
		}
	}

	return selected
}

func main() {
	activities := []Activity{
		{1, 1, 3},
		{2, 2, 5},
		{3, 4, 6},
		{4, 6, 7},
	}

	fmt.Println("=== Задача о расписании (Activity Selection) ===")
	fmt.Println("Занятия:")
	for _, a := range activities {
		fmt.Printf("  Занятие %d: [%d, %d]\n", a.id, a.start, a.end)
	}

	result := activitySelection(activities)

	fmt.Printf("\nВыбранные занятия (максимальное непересекающееся подмножество):\n")
	for _, a := range result {
		fmt.Printf("  Занятие %d: [%d, %d]\n", a.id, a.start, a.end)
	}
	fmt.Printf("Количество занятий: %d\n", len(result))
}
