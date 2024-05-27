package _map

func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	m := make(map[K]V)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}

// Intersect 取交集
func Intersect[K comparable, V any](m1, m2 map[K]V) map[K]V {
	intersection := make(map[K]V)
	for k1 := range m1 {
		if v2, ok := m2[k1]; ok {
			intersection[k1] = v2
		}
	}
	return intersection
}
