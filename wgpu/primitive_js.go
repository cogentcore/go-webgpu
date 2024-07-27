//go:build js

package wgpu

// mapSlice can be used to transform one slice into another by providing a
// function to do the mapping.
func mapSlice[S, T any](slice []S, fn func(S) T) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Color as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpucolor
type Color struct {
	R float64
	G float64
	B float64
	A float64
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Color) ToJS() any {
	return []any{g.R, g.G, g.B, g.A}
}
