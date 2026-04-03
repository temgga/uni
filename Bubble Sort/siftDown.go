package main

func siftdown(data []int, i int) {
	n := len(data)
	hole := data[i]
	for {
		left := 2*i + 1
		if left >= n { //если не существ выходим
			break
		}
		small := left
		if right := left + 1; right < n && data[right] < data[left] {
			//если правый существ и < левого берем правого
			small = right
		}
		if hole <= data[small] { //порядок нарушен в таком сл
			break
		}
		data[i] = data[small] //наименьш верх вместо дырки
		i = small             //переход к месту где тепер дырка
	}
	data[i] = hole //кладем известный элемент в дырку
}
