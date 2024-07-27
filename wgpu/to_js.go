//go:build js

package wgpu

import "syscall/js"

// ToJS converts a type to one that can be passed as an argument
// to JavaScript.

func (g Color) ToJS() any {
	return []any{g.R, g.G, g.B, g.A}
}

func (g *Extent3D) ToJS() any {
	return []any{g.Width, g.Height, g.DepthOrArrayLayers}
}

func (g *RequestAdapterOptions) ToJS() js.Value {
	result := make(map[string]any)
	result["powerPreference"] = g.PowerPreference.String()
	result["forceFallbackAdapter"] = g.ForceFallbackAdapter
	return js.ValueOf(result)
}

func (g *DeviceDescriptor) ToJS() js.Value {
	result := make(map[string]any)
	result["label"] = g.Label
	result["requiredFeatures"] = mapSlice(g.RequiredFeatures, func(f FeatureName) any { return f })
	// result["requiredLimits"] = // TODO(kai): convert requiredLimits to JS
	return js.ValueOf(result)
}

func (g *TextureViewDescriptor) ToJS() js.Value {
	return js.ValueOf(map[string]any{
		"label":           g.Label,
		"format":          g.Format.String(),
		"dimension":       g.Dimension.String(),
		"baseMipLevel":    g.BaseMipLevel,
		"mipLevelCount":   g.MipLevelCount,
		"baseArrayLayer":  g.BaseArrayLayer,
		"arrayLayerCount": g.ArrayLayerCount,
		"aspect":          g.Aspect.String(),
	})

}

func (g BufferDescriptor) ToJS() any {
	return map[string]any{
		"size":             g.Size,
		"usage":            g.Usage.String(),
		"mappedAtCreation": g.MappedAtCreation,
	}
}

func (g *ImageCopyBuffer) ToJS() any {
	return map[string]any{
		"layout": g.Layout.ToJS(),
		"buffer": g.Buffer.ToJS(),
	}
}

func (g *ImageCopyTexture) ToJS() any {
	return map[string]any{
		"texture":  g.Texture.ToJS(),
		"mipLevel": g.MipLevel,
		"origin":   g.Origin.ToJS(),
		"aspect":   g.Aspect.String(),
	}
}

func (g *TextureDataLayout) ToJS() any {
	return map[string]any{
		"offset":       g.Offset,
		"bytesPerRow":  g.BytesPerRow,
		"rowsPerImage": g.RowsPerImage,
	}
}

func (g *Origin3D) ToJS() any {
	return []any{g.X, g.Y, g.Z}
}
