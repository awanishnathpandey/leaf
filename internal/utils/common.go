package utils

func ConvertToPointerSlice(input []string) []*string {
	result := make([]*string, len(input))
	for i, val := range input {
		result[i] = &val
	}
	return result
}
