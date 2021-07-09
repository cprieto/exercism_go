// Package flatten works with lists
package flatten

// Flatten a given interface int a list
func Flatten(input interface{}) []interface{} {
	output := make([]interface{}, 0)

	switch elem := input.(type) {
	case []interface{}:
		for _, val := range elem {
			output = append(output, Flatten(val)...)
		}
	case int:
		output = append(output, elem)
	}

	return output
}
