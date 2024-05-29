package utilities

import "reflect"

// Merge merges two structs of the same type such that non-zero values from the updated struct take precedence.
func Merge[T any](existing T, updated T) T {
	existingVal := reflect.ValueOf(&existing).Elem()
	updatedVal := reflect.ValueOf(updated)

	for i := 0; i < existingVal.NumField(); i++ {
		existingField := existingVal.Field(i)
		updatedField := updatedVal.Field(i)

		// Check if the updated field is non-zero
		if !isZeroValue(updatedField) {
			existingField.Set(updatedField)
		}
	}

	return existing
}

// isZeroValue checks if a reflect.Value is the zero value for its type
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
