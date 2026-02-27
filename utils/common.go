package utils

import (
	"fmt"
)

// ConvertToPtr attempts to safely type assert a cached item
// to a pointer of type T.
func ConvertToPtr[T any](item any) (*T, error) {
	if val, ok := item.(*T); ok {
		return val, nil
	} else if val, ok := item.(T); ok {
		return &val, nil
	}
	return nil, fmt.Errorf("variable is not of the expected type")
}
