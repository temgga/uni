package main

func siftUp(data []int, i int) {
	hole := data[i]
	for i > 0 {
		parent := (i - 1) / 2
		if hole >= data[parent] {
			break
		}
		data[i] = data[parent]
		//передвигаем родителя вниз
		//родитель теперь на позиции i
		//а дырка на позиции parent
		i = parent //дырка там где был родитель
	}
	data[i] = hole //кладем наш элем в дырку
}
