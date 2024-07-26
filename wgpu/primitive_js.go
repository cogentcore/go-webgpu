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

// BufferDynamicOffset as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpubufferdynamicoffset
type BufferDynamicOffset uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferDynamicOffset) ToJS() any {
	return uint32(g)
}

// Size64 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusize64
type Size64 uint64

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Size64) ToJS() any {
	return uint64(g)
}

// IntegerCoordinate as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuintegercoordinate
type IntegerCoordinate uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g IntegerCoordinate) ToJS() any {
	return uint32(g)
}

// Index32 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuindex32
type Index32 uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Index32) ToJS() any {
	return uint32(g)
}

// Size32 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusize32
type Size32 uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g Size32) ToJS() any {
	return uint32(g)
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

// FlagsConstant as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuflagsconstant
type FlagsConstant uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g FlagsConstant) ToJS() any {
	return uint32(g)
}

// StencilValue as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpustencilvalue
type StencilValue uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StencilValue) ToJS() any {
	return uint32(g)
}

// DepthBias as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpudepthbias
type DepthBias int32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g DepthBias) ToJS() any {
	return int32(g)
}

// SampleMask as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusamplemask
type SampleMask uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g SampleMask) ToJS() any {
	return uint32(g)
}
