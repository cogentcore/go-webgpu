//go:build js

package wgpu

import (
	"syscall/js"
)

// BufferBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbindinglayout
type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   uint64
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBindingLayout) toJS() any {
	result := make(map[string]any)
	result["type"] = toJS(g.Type)
	result["hasDynamicOffset"] = g.HasDynamicOffset
	result["minBindingSize"] = g.MinBindingSize
	return result
}

// SamplerBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpusamplerbindinglayout
type SamplerBindingLayout struct {
	Type SamplerBindingType
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g SamplerBindingLayout) toJS() any {
	result := make(map[string]any)
	result["type"] = toJS(g.Type)
	return result
}

// TextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gputexturebindinglayout
type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g TextureBindingLayout) toJS() any {
	result := make(map[string]any)
	result["sampleType"] = toJS(g.SampleType)
	result["viewDimension"] = toJS(g.ViewDimension)
	result["multisampled"] = g.Multisampled
	return result
}

// StorageTextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustoragetexturebindinglayout
type StorageTextureBindingLayout struct {
	Access        StorageTextureAccess
	Format        TextureFormat
	ViewDimension TextureViewDimension
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StorageTextureBindingLayout) toJS() any {
	result := make(map[string]any)
	result["access"] = toJS(g.Access)
	result["format"] = toJS(g.Format)
	result["viewDimension"] = toJS(g.ViewDimension)
	return result
}

// ExternalTextureBindingLayout as described:
type ExternalTextureBindingLayout struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ExternalTextureBindingLayout) toJS() any {
	return g.jsValue
}

// BindGroupLayoutEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutentry
type BindGroupLayoutEntry struct {
	Binding         uint32
	Visibility      ShaderStage
	Buffer          BufferBindingLayout
	Sampler         SamplerBindingLayout
	Texture         TextureBindingLayout
	StorageTexture  StorageTextureBindingLayout
	ExternalTexture ExternalTextureBindingLayout
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayoutEntry) toJS() any {
	result := make(map[string]any)
	result["binding"] = g.Binding
	result["visibility"] = toJS(g.Visibility)
	result["buffer"] = g.Buffer.toJS()
	result["sampler"] = g.Sampler.toJS()
	result["texture"] = g.Texture.toJS()
	result["storageTexture"] = g.StorageTexture.toJS()
	result["externalTexture"] = g.ExternalTexture.toJS()
	return result
}

// BindGroupLayoutDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutdescriptor
type BindGroupLayoutDescriptor struct {
	Entries []BindGroupLayoutEntry
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayoutDescriptor) toJS() any {
	return map[string]any{
		"entries": mapSlice(g.Entries, func(entry BindGroupLayoutEntry) any {
			return entry.toJS()
		}),
	}
}

// BindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#gpubindgrouplayout
type BindGroupLayout struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayout) toJS() any {
	return g.jsValue
}

// BufferBinding as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbinding
type BufferBinding struct {
	Buffer Buffer
	Offset uint64
	Size   uint64
}

var _ BindingResource = BufferBinding{}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBinding) toJS() any {
	result := make(map[string]any)
	result["buffer"] = g.Buffer.toJS()
	result["offset"] = g.Offset
	result["size"] = g.Size
	return result
}

func (g BufferBinding) _isBindingResource() {}

// BindingResource as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpubindingresource
type BindingResource interface {
	_isBindingResource()
	toJS() any
}

// BindGroupEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupentry
type BindGroupEntry struct {
	Binding  uint32
	Resource BindingResource
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupEntry) toJS() any {
	return map[string]any{
		"binding":  g.Binding,
		"resource": g.Resource.toJS(),
	}
}

// BindGroupDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupdescriptor
type BindGroupDescriptor struct {
	Layout  BindGroupLayout
	Entries []BindGroupEntry
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupDescriptor) toJS() any {
	return map[string]any{
		"layout": g.Layout.toJS(),
		"entries": mapSlice(g.Entries, func(entry BindGroupEntry) any {
			return entry.toJS()
		}),
	}
}

// BindGroup as described:
// https://gpuweb.github.io/gpuweb/#gpubindgroup
type BindGroup struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroup) toJS() any {
	return g.jsValue
}
