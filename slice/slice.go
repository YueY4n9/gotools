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

func IsExist[T comparable](slice []T, ele T) bool {
	for _, e := range slice {
		if e == ele {
			return true
		}
	}
	return false
}

func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// Subtract return new(slice1 - slice2)
func Subtract[T comparable](slice1, slice2 []T) []T {
	res := make([]T, 0)
	m := make(map[T]bool)
	for _, e := range slice2 {
		m[e] = true
	}
	for _, e := range slice1 {
		if !m[e] {
			res = append(res, e)
		}
	}
	return res
}

// Intersect 取交集
func Intersect[T comparable](slice1, slice2 []T) []T {
	hashSet := make(map[T]bool)
	res := make([]T, 0)
	for _, v := range slice1 {
		hashSet[v] = true
	}
	for _, v := range slice2 {
		if hashSet[v] {
			res = append(res, v)
		}
	}
	return res
}
