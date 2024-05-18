package utilities

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	merged := make(map[K]V)

	for _, m := range maps {
		for key, value := range m {
			merged[key] = value
		}
	}

	return merged
}
