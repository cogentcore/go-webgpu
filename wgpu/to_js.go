//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// toJS converts the given value to a type that can be passed as
// an argument to JavaScript. It should only be called with pointers
// and enums.
func toJS[T comparable](v T) any {
	var zero T
	if v == zero {
		return js.Undefined()
	}
	if tj, ok := any(v).(interface{ toJS() any }); ok {
		return tj.toJS()
	}
	if s, ok := any(v).(fmt.Stringer); ok {
		ss := s.String()
		if ss == "undefined" {
			return js.Undefined()
		}
		return ss
	}
	return v
}

func (g Color) toJS() any {
	return []any{g.R, g.G, g.B, g.A}
}

func (g *Extent3D) toJS() any {
	return []any{g.Width, g.Height, g.DepthOrArrayLayers}
}

func (g *RequestAdapterOptions) toJS() any {
	result := make(map[string]any)
	result["powerPreference"] = toJS(g.PowerPreference)
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

func (g *TextureViewDescriptor) toJS() any {
	return map[string]any{
		"label":           g.Label,
		"format":          toJS(g.Format),
		"dimension":       toJS(g.Dimension),
		"baseMipLevel":    g.BaseMipLevel,
		"mipLevelCount":   g.MipLevelCount,
		"baseArrayLayer":  g.BaseArrayLayer,
		"arrayLayerCount": g.ArrayLayerCount,
		"aspect":          toJS(g.Aspect),
	}

}

func (g *CommandEncoderDescriptor) toJS() any {
	return map[string]any{"label": g.Label}
}

func (g BufferDescriptor) toJS() any {
	return map[string]any{
		"size":             g.Size,
		"usage":            toJS(g.Usage),
		"mappedAtCreation": g.MappedAtCreation,
	}
}

func (g *ImageCopyBuffer) toJS() any {
	return map[string]any{
		"layout": toJS(g.Layout),
		"buffer": toJS(g.Buffer),
	}
}

func (g *ImageCopyTexture) toJS() any {
	return map[string]any{
		"texture":  toJS(g.Texture),
		"mipLevel": g.MipLevel,
		"origin":   toJS(g.Origin),
		"aspect":   toJS(g.Aspect),
	}
}

func (g *TextureDataLayout) toJS() any {
	return map[string]any{
		"offset":       g.Offset,
		"bytesPerRow":  g.BytesPerRow,
		"rowsPerImage": g.RowsPerImage,
	}
}

func (g *Origin3D) toJS() any {
	return []any{g.X, g.Y, g.Z}
}

func (g *RenderPassColorAttachment) toJS() any {
	result := make(map[string]any)
	result["view"] = g.View.jsValue
	result["loadOp"] = toJS(g.LoadOp)
	result["storeOp"] = toJS(g.StoreOp)
	result["clearValue"] = toJS(g.ClearValue)
	result["resolveTarget"] = toJS(g.ResolveTarget)
	return result
}
