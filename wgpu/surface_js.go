//go:build js

package wgpu

import "syscall/js"

type Surface struct{} // no JS equivalent

func GetPreferredFormat(adapter *Adapter) TextureFormat {
	// TODO(kai): need to set enum from string
	return TextureFormat(js.Global().Get("navigator").Get("gpu").Call("getPreferredCanvasFormat").Int())
}
