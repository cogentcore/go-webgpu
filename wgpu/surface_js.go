//go:build js

package wgpu

type Surface struct{} // no JS equivalent

func (g Surface) GetPreferredFormat(adapter *Adapter) TextureFormat {
	return TextureFormatBGRA8Unorm
	// TODO(kai): need to set enum from string
	// return TextureFormat(js.Global().Get("navigator").Get("gpu").Call("getPreferredCanvasFormat").Int())
}

func (g Surface) GetCapabilities(adapter *Adapter) (ret SurfaceCapabilities) {
	// Based on https://developer.mozilla.org/en-US/docs/Web/API/GPUCanvasContext/configure
	ret.Formats = []TextureFormat{TextureFormatBGRA8Unorm, TextureFormatRGBA8Unorm, TextureFormatRGBA16Float}
	ret.AlphaModes = []CompositeAlphaMode{CompositeAlphaModeOpaque, CompositeAlphaModePreMultiplied}
	ret.PresentModes = []PresentMode{PresentModeImmediate}
	return
}

func (g Surface) Release() {} // no-op
