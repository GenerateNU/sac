package utilities

// MergeMaps merges multiple maps into a single map.
// The keys of maps earlier in the list will take precedence over keys in later maps.
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	merged := make(map[K]V)

	for index := len(maps) - 1; index >= 0; index-- {
		for k, v := range maps[index] {
			merged[k] = v
		}
	}

	return merged
}
