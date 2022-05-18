package homework

func reverse(input []int64) (result []int64) {
	sliceLen := len(input)
	sliceLen--
	for i := sliceLen; i >= 0; i-- {
		result = append(result, input[i])
	}

	return
}
