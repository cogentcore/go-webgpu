//go:build js

package wgpu

import "syscall/js"

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

// await is a helper function roughly equivalent to await in JS.
func await(promise js.Value) js.Value {
	result := make(chan js.Value)
	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		result <- args[0]
		return nil
	}))
	return <-result
}

func (g Color) ToJS() any {
	return []any{g.R, g.G, g.B, g.A}
}

func (g *Extent3D) ToJS() any {
	return []any{g.Width, g.Height, g.DepthOrArrayLayers}
}

// no-ops
func SetLogLevel(level LogLevel) {}
func GetVersion() Version        { return 0 }
