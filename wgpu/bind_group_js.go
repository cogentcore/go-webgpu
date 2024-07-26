//go:build js

package wgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// BufferBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbindinglayout
type BufferBindingLayout struct {
	Type             opt.T[GPUBufferBindingType]
	HasDynamicOffset opt.T[bool]
	MinBindingSize   opt.T[GPUSize64]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Type.Specified {
		result["type"] = g.Type.Value.ToJS()
	}
	if g.HasDynamicOffset.Specified {
		result["hasDynamicOffset"] = g.HasDynamicOffset.Value
	}
	if g.MinBindingSize.Specified {
		result["minBindingSize"] = g.MinBindingSize.Value.ToJS()
	}
	return result
}

// SamplerBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpusamplerbindinglayout
type SamplerBindingLayout struct {
	Type opt.T[GPUSamplerBindingType]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g SamplerBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Type.Specified {
		result["type"] = g.Type.Value.ToJS()
	}
	return result
}

// TextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gputexturebindinglayout
type TextureBindingLayout struct {
	SampleType    opt.T[GPUTextureSampleType]
	ViewDimension opt.T[GPUTextureViewDimension]
	Multisampled  opt.T[bool]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g TextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.SampleType.Specified {
		result["sampleType"] = g.SampleType.Value.ToJS()
	}
	if g.ViewDimension.Specified {
		result["viewDimension"] = g.ViewDimension.Value.ToJS()
	}
	if g.Multisampled.Specified {
		result["multisampled"] = g.Multisampled.Value
	}
	return result
}

// StorageTextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustoragetexturebindinglayout
type StorageTextureBindingLayout struct {
	Access        opt.T[GPUStorageTextureAccess]
	Format        GPUTextureFormat
	ViewDimension opt.T[GPUTextureViewDimension]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StorageTextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Access.Specified {
		result["access"] = g.Access.Value.ToJS()
	}
	result["format"] = g.Format.ToJS()
	if g.ViewDimension.Specified {
		result["viewDimension"] = g.ViewDimension.Value.ToJS()
	}
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
	Visibility      GPUShaderStageFlags
	Buffer          opt.T[GPUBufferBindingLayout]
	Sampler         opt.T[GPUSamplerBindingLayout]
	Texture         opt.T[GPUTextureBindingLayout]
	StorageTexture  opt.T[GPUStorageTextureBindingLayout]
	ExternalTexture opt.T[GPUExternalTextureBindingLayout]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BindGroupLayoutEntry) ToJS() any {
	result := make(map[string]any)
	result["binding"] = g.Binding.ToJS()
	result["visibility"] = g.Visibility.ToJS()
	if g.Buffer.Specified {
		result["buffer"] = g.Buffer.Value.ToJS()
	}
	if g.Sampler.Specified {
		result["sampler"] = g.Sampler.Value.ToJS()
	}
	if g.Texture.Specified {
		result["texture"] = g.Texture.Value.ToJS()
	}
	if g.StorageTexture.Specified {
		result["storageTexture"] = g.StorageTexture.Value.ToJS()
	}
	if g.ExternalTexture.Specified {
		result["externalTexture"] = g.ExternalTexture.Value.ToJS()
	}
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
	Offset opt.T[GPUSize64]
	Size   opt.T[GPUSize64]
}

var _ BindingResource = BufferBinding{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BufferBinding) ToJS() any {
	result := make(map[string]any)
	result["buffer"] = g.Buffer.ToJS()
	if g.Offset.Specified {
		result["offset"] = g.Offset.Value.ToJS()
	}
	if g.Size.Specified {
		result["size"] = g.Size.Value.ToJS()
	}
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
