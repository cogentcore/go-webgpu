//go:build js

package wgpu

import "syscall/js"

// Adapter as described:
// https://gpuweb.github.io/gpuweb/#gpuadapter
type Adapter struct {
	jsValue js.Value
}

func (g Adapter) Release() {} // no-op

type Surface struct{} // no-op
