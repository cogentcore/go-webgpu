//go:build js

package wgpu

import "syscall/js"

type Surface struct{} // no JS equivalent

func (g Surface) GetPreferredFormat(adapter *Adapter) TextureFormat {
	// TODO(kai): need to set enum from string
	return TextureFormat(js.Global().Get("navigator").Get("gpu").Call("getPreferredCanvasFormat").Int())
}

func (g Surface) GetCapabilities(adapter *Adapter) (ret SurfaceCapabilities) {
	// TODO(kai): get capabilities
	return
}
