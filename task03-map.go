package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	numbers := make([]int, 0)
	for key := range input {
		numbers = append(numbers, key)
	}
	sort.Ints(numbers)

	for _, val := range numbers {
		result = append(result, input[val])
	}
	return
}
