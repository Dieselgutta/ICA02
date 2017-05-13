package algorithms

// Deres kode her
func Bubble_sort_modified(list []int) {
	length := len(list)
	for {
		unsorted := true
		for index := 0; index < length-1; index++ {
			if list[index] > list[index+1] {
				temp := list[index+1]
				list[index+1] = list[index]
				list[index] = temp
				unsorted = false
			}
		}
		length -= 1
		if unsorted == true {
			break
		}
	}
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
