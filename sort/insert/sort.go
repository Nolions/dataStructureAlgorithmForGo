package insert

func sort(data []int) {
	for i := 0; i < len(data); i++ {
		v := data[i]

		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > v {
				data[j+1] = data[j]
			} else {
				break
			}
		}

		data[j+1] = v
	}
}
