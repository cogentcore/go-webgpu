//go:build js

package wgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
)

// BufferBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbindinglayout
type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   Size64
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBindingLayout) ToJS() any {
	result := make(map[string]any)
	result["type"] = g.Type.String()
	result["hasDynamicOffset"] = g.HasDynamicOffset
	result["minBindingSize"] = g.MinBindingSize
	return result
}

// SamplerBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpusamplerbindinglayout
type SamplerBindingLayout struct {
	Type SamplerBindingType
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g SamplerBindingLayout) ToJS() any {
	result := make(map[string]any)
	result["type"] = g.Type.String()
	return result
}

// TextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gputexturebindinglayout
type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g TextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	result["sampleType"] = g.SampleType.String()
	result["viewDimension"] = g.ViewDimension.String()
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

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StorageTextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	result["access"] = g.Access.String()
	result["format"] = g.Format.String()
	result["viewDimension"] = g.ViewDimension.String()
	return result
}

// ExternalTextureBindingLayout as described:
type ExternalTextureBindingLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ExternalTextureBindingLayout) ToJS() any {
	return g.jsValue
}

// BindGroupLayoutEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutentry
type BindGroupLayoutEntry struct {
	Binding         Index32
	Visibility      ShaderStage
	Buffer          BufferBindingLayout
	Sampler         SamplerBindingLayout
	Texture         TextureBindingLayout
	StorageTexture  StorageTextureBindingLayout
	ExternalTexture ExternalTextureBindingLayout
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayoutEntry) ToJS() any {
	result := make(map[string]any)
	result["binding"] = g.Binding.ToJS()
	result["visibility"] = g.Visibility.String()
	result["buffer"] = g.Buffer.ToJS()
	result["sampler"] = g.Sampler.ToJS()
	result["texture"] = g.Texture.ToJS()
	result["storageTexture"] = g.StorageTexture.ToJS()
	result["externalTexture"] = g.ExternalTexture.ToJS()
	return result
}

// BindGroupLayoutDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutdescriptor
type BindGroupLayoutDescriptor struct {
	Entries []BindGroupLayoutEntry
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayoutDescriptor) ToJS() any {
	return map[string]any{
		"entries": gog.Map(g.Entries, func(entry BindGroupLayoutEntry) any {
			return entry.ToJS()
		}),
	}
}

// BindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#gpubindgrouplayout
type BindGroupLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayout) ToJS() any {
	return g.jsValue
}

// BufferBinding as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbinding
type BufferBinding struct {
	Buffer Buffer
	Offset Size64
	Size   Size64
}

var _ BindingResource = BufferBinding{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBinding) ToJS() any {
	result := make(map[string]any)
	result["buffer"] = g.Buffer.ToJS()
	result["offset"] = g.Offset.ToJS()
	result["size"] = g.Size.ToJS()
	return result
}

func (g BufferBinding) _isBindingResource() {}

// BindingResource as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpubindingresource
type BindingResource interface {
	_isBindingResource()
	ToJS() any
}

// BindGroupEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupentry
type BindGroupEntry struct {
	Binding  Index32
	Resource BindingResource
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupEntry) ToJS() any {
	return map[string]any{
		"binding":  g.Binding.ToJS(),
		"resource": g.Resource.ToJS(),
	}
}

// BindGroupDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupdescriptor
type BindGroupDescriptor struct {
	Layout  BindGroupLayout
	Entries []BindGroupEntry
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupDescriptor) ToJS() any {
	return map[string]any{
		"layout": g.Layout.ToJS(),
		"entries": gog.Map(g.Entries, func(entry BindGroupEntry) any {
			return entry.ToJS()
		}),
	}
}

// BindGroup as described:
// https://gpuweb.github.io/gpuweb/#gpubindgroup
type BindGroup struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroup) ToJS() any {
	return g.jsValue
}
