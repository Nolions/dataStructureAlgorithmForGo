package sort

func selectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		minIndex := i // 紀錄陣列中最小值索引
		for j := i + 1; j < len(data); j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}
