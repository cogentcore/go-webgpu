//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// enumToJS converts the given non-bit-flag enum value to a type that
// can be passed as an argument to JavaScript. Bit flag enums should be
// passed as a uint.
func enumToJS(s fmt.Stringer) any {
	ss := s.String()
	if ss == "undefined" {
		return js.Undefined()
	}
	return ss
}

// pointerToJS converts the given pointer value to a type that can be
// passed as an argument to JavaScript. It must implement a toJS method.
func pointerToJS[T any, P interface {
	*T
	toJS() any
}](v P) any {
	if v == nil {
		return js.Undefined()
	}
	return v.toJS()
}

func (g Color) toJS() any {
	return []any{g.R, g.G, g.B, g.A}
}

func (g Extent3D) toJS() any {
	return []any{g.Width, g.Height, g.DepthOrArrayLayers}
}

func (g Origin3D) toJS() any {
	return []any{g.X, g.Y, g.Z}
}

func (g *RequestAdapterOptions) toJS() any {
	result := make(map[string]any)
	result["powerPreference"] = enumToJS(g.PowerPreference)
	result["forceFallbackAdapter"] = g.ForceFallbackAdapter
	return result
}

func (g *DeviceDescriptor) toJS() any {
	result := make(map[string]any)
	result["label"] = g.Label
	result["requiredFeatures"] = mapSlice(g.RequiredFeatures, func(f FeatureName) any { return f })
	// result["requiredLimits"] = // TODO(kai): convert requiredLimits to JS
	return result
}

func (g *TextureDescriptor) toJS() any {
	return map[string]any{
		"label":         g.Label,
		"usage":         uint32(g.Usage),
		"dimension":     enumToJS(g.Dimension),
		"size":          g.Size.toJS(),
		"format":        enumToJS(g.Format),
		"mipLevelCount": g.MipLevelCount,
		"sampleCount":   g.SampleCount,
	}
}

func (g *TextureViewDescriptor) toJS() any {
	return map[string]any{
		"label":           g.Label,
		"format":          enumToJS(g.Format),
		"dimension":       enumToJS(g.Dimension),
		"baseMipLevel":    g.BaseMipLevel,
		"mipLevelCount":   g.MipLevelCount,
		"baseArrayLayer":  g.BaseArrayLayer,
		"arrayLayerCount": g.ArrayLayerCount,
		"aspect":          enumToJS(g.Aspect),
	}
}

func (g *CommandEncoderDescriptor) toJS() any {
	return map[string]any{"label": g.Label}
}

func (g *CommandBufferDescriptor) toJS() any {
	return map[string]any{"label": g.Label}
}

func (g BufferDescriptor) toJS() any {
	return map[string]any{
		"size":             g.Size,
		"usage":            uint32(g.Usage),
		"mappedAtCreation": g.MappedAtCreation,
	}
}

func (g *ImageCopyBuffer) toJS() any {
	return map[string]any{
		"layout": g.Layout.toJS(),
		"buffer": pointerToJS(g.Buffer),
	}
}

func (g *ImageCopyTexture) toJS() any {
	return map[string]any{
		"texture":  pointerToJS(g.Texture),
		"mipLevel": g.MipLevel,
		"origin":   g.Origin.toJS(),
		"aspect":   enumToJS(g.Aspect),
	}
}

func (g *TextureDataLayout) toJS() any {
	return map[string]any{
		"offset":       g.Offset,
		"bytesPerRow":  g.BytesPerRow,
		"rowsPerImage": g.RowsPerImage,
	}
}

func (g *RenderPassColorAttachment) toJS() any {
	result := make(map[string]any)
	result["view"] = g.View.jsValue
	result["loadOp"] = enumToJS(g.LoadOp)
	result["storeOp"] = enumToJS(g.StoreOp)
	result["clearValue"] = g.ClearValue.toJS()
	result["resolveTarget"] = pointerToJS(g.ResolveTarget)
	return result
}
