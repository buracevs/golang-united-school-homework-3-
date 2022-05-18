package homework

func average(input [15]float32) (result float32) {
	var sum float32
	totalItems := 0
	for _, item := range input {
		sum += item
		totalItems++
	}
	result = sum / float32(totalItems)
	return
}
