package _slice

func ChunkSlice[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func RemoveDuplication[T comparable](slice []T) []T {
	encountered := map[T]bool{}
	result := make([]T, 0)
	for _, e := range slice {
		if !encountered[e] {
			encountered[e] = true
			result = append(result, e)
		}
	}
	return result
}

func RemoveElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func ToSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, e := range slice {
		set[e] = struct{}{}
	}
	return set
}
