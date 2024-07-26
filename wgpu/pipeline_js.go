//go:build js

package wgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// PipelineLayoutDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpupipelinelayoutdescriptor
type PipelineLayoutDescriptor struct {
	BindGroupLayouts []BindGroupLayout
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g PipelineLayoutDescriptor) ToJS() any {
	return map[string]any{
		"bindGroupLayouts": gog.Map(g.BindGroupLayouts, func(layout BindGroupLayout) any {
			return layout.ToJS()
		}),
	}
}

// PipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#gpupipelinelayout
type PipelineLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g PipelineLayout) ToJS() any {
	return g.jsValue
}

// VertexAttribute as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexattribute
type VertexAttribute struct {
	Format         VertexFormat
	Offset         Size64
	ShaderLocation Index32
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g VertexAttribute) ToJS() any {
	return map[string]any{
		"format":         g.Format.String(),
		"offset":         g.Offset.ToJS(),
		"shaderLocation": g.ShaderLocation.ToJS(),
	}
}

// VertexBufferLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexbufferlayout
type VertexBufferLayout struct {
	ArrayStride Size64
	StepMode    opt.T[GPUVertexStepMode]
	Attributes  []VertexAttribute
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g VertexBufferLayout) ToJS() any {
	result := make(map[string]any)
	result["arrayStride"] = g.ArrayStride.ToJS()
	if g.StepMode.Specified {
		result["stepMode"] = g.StepMode.Value.ToJS()
	}
	result["attributes"] = gog.Map(g.Attributes, func(attrib VertexAttribute) any {
		return attrib.ToJS()
	})
	return result
}

// VertexState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexstate
type VertexState struct {
	Module     ShaderModule
	EntryPoint string
	Buffers    []VertexBufferLayout
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g VertexState) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
		"buffers": gog.Map(g.Buffers, func(layout VertexBufferLayout) any {
			return layout.ToJS()
		}),
	}
}

// PrimitiveState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuprimitivestate
type PrimitiveState struct {
	Topology         opt.T[GPUPrimitiveTopology]
	StripIndexFormat opt.T[GPUIndexFormat]
	FrontFace        opt.T[GPUFrontFace]
	CullMode         opt.T[GPUCullMode]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g PrimitiveState) ToJS() any {
	result := make(map[string]any)
	if g.Topology.Specified {
		result["topology"] = g.Topology.Value.ToJS()
	}
	if g.StripIndexFormat.Specified {
		result["stripIndexFormat"] = g.StripIndexFormat.Value.ToJS()
	}
	if g.FrontFace.Specified {
		result["frontFace"] = g.FrontFace.Value.ToJS()
	}
	if g.CullMode.Specified {
		result["cullMode"] = g.CullMode.Value.ToJS()
	}
	return result
}

// StencilFaceState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustencilfacestate
type StencilFaceState struct {
	Compare     opt.T[GPUCompareFunction]
	FailOp      opt.T[GPUStencilOperation]
	DepthFailOp opt.T[GPUStencilOperation]
	PassOp      opt.T[GPUStencilOperation]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StencilFaceState) ToJS() any {
	result := make(map[string]any)
	if g.Compare.Specified {
		result["compare"] = g.Compare.Value.ToJS()
	}
	if g.FailOp.Specified {
		result["failOp"] = g.FailOp.Value.ToJS()
	}
	if g.DepthFailOp.Specified {
		result["depthFailOp"] = g.DepthFailOp.Value.ToJS()
	}
	if g.PassOp.Specified {
		result["passOp"] = g.PassOp.Value.ToJS()
	}
	return result
}

// DepthStencilState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpudepthstencilstate
type DepthStencilState struct {
	Format              TextureFormat
	DepthWriteEnabled   bool
	DepthCompare        CompareFunction
	StencilFront        opt.T[GPUStencilFaceState]
	StencilBack         opt.T[GPUStencilFaceState]
	StencilReadMask     opt.T[GPUStencilValue]
	StencilWriteMask    opt.T[GPUStencilValue]
	DepthBias           opt.T[GPUDepthBias]
	DepthBiasSlopeScale opt.T[float32]
	DepthBiasClamp      opt.T[float32]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g DepthStencilState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.String()
	result["depthWriteEnabled"] = g.DepthWriteEnabled
	result["depthCompare"] = g.DepthCompare.String()
	if g.StencilFront.Specified {
		result["stencilFront"] = g.StencilFront.Value.ToJS()
	}
	if g.StencilBack.Specified {
		result["stencilBack"] = g.StencilBack.Value.ToJS()
	}
	if g.StencilReadMask.Specified {
		result["stencilReadMask"] = g.StencilReadMask.Value.ToJS()
	}
	if g.StencilWriteMask.Specified {
		result["stencilWriteMask"] = g.StencilWriteMask.Value.ToJS()
	}
	if g.DepthBias.Specified {
		result["depthBias"] = g.DepthBias.Value.ToJS()
	}
	if g.DepthBiasSlopeScale.Specified {
		result["depthBiasSlopeScale"] = g.DepthBiasSlopeScale.Value
	}
	if g.DepthBiasClamp.Specified {
		result["depthBiasClamp"] = g.DepthBiasClamp.Value
	}
	return result
}

// MultisampleState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpumultisamplestate
type MultisampleState struct {
	Count                  opt.T[GPUSize32]
	Mask                   opt.T[GPUSampleMask]
	AlphaToCoverageEnabled opt.T[bool]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g MultisampleState) ToJS() any {
	result := make(map[string]any)
	if g.Count.Specified {
		result["count"] = g.Count.Value.ToJS()
	}
	if g.Mask.Specified {
		result["mask"] = g.Mask.Value.ToJS()
	}
	if g.AlphaToCoverageEnabled.Specified {
		result["alphaToCoverageEnabled"] = g.AlphaToCoverageEnabled.Value
	}
	return result
}

// BlendComponent as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendcomponent
type BlendComponent struct {
	Operation opt.T[GPUBlendOperation]
	SrcFactor opt.T[GPUBlendFactor]
	DstFactor opt.T[GPUBlendFactor]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BlendComponent) ToJS() any {
	result := make(map[string]any)
	if g.Operation.Specified {
		result["operation"] = g.Operation.Value.ToJS()
	}
	if g.SrcFactor.Specified {
		result["srcFactor"] = g.SrcFactor.Value.ToJS()
	}
	if g.DstFactor.Specified {
		result["dstFactor"] = g.DstFactor.Value.ToJS()
	}
	return result
}

// BlendState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendstate
type BlendState struct {
	Color BlendComponent
	Alpha BlendComponent
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BlendState) ToJS() any {
	return map[string]any{
		"color": g.Color.ToJS(),
		"alpha": g.Alpha.ToJS(),
	}
}

// ColorTargetState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucolortargetstate
type ColorTargetState struct {
	Format    TextureFormat
	Blend     opt.T[GPUBlendState]
	WriteMask opt.T[GPUColorWriteFlags]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ColorTargetState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.String()
	if g.Blend.Specified {
		result["blend"] = g.Blend.Value.ToJS()
	}
	if g.WriteMask.Specified {
		result["writeMask"] = g.WriteMask.Value.ToJS()
	}
	return result
}

// FragmentState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpufragmentstate
type FragmentState struct {
	Module     ShaderModule
	EntryPoint string
	Targets    []ColorTargetState
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g FragmentState) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
		"targets": gog.Map(g.Targets, func(target ColorTargetState) any {
			return target.ToJS()
		}),
	}
}

// RenderPipelineDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpipelinedescriptor
type RenderPipelineDescriptor struct {
	Layout       opt.T[GPUPipelineLayout]
	Vertex       VertexState
	Primitive    opt.T[GPUPrimitiveState]
	DepthStencil opt.T[GPUDepthStencilState]
	Multisample  opt.T[GPUMultisampleState]
	Fragment     opt.T[GPUFragmentState]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPipelineDescriptor) ToJS() any {
	result := make(map[string]any)
	if g.Layout.Specified {
		result["layout"] = g.Layout.Value.ToJS()
	} else {
		result["layout"] = "auto"
	}
	result["vertex"] = g.Vertex.ToJS()
	if g.Primitive.Specified {
		result["primitive"] = g.Primitive.Value.ToJS()
	}
	if g.DepthStencil.Specified {
		result["depthStencil"] = g.DepthStencil.Value.ToJS()
	}
	if g.Multisample.Specified {
		result["multisample"] = g.Multisample.Value.ToJS()
	}
	if g.Fragment.Specified {
		result["fragment"] = g.Fragment.Value.ToJS()
	}
	return result
}

// RenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#gpurenderpipeline
type RenderPipeline struct {
	jsValue js.Value
}

// GetBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpupipelinebase-getbindgrouplayout
func (g RenderPipeline) GetBindGroupLayout(index uint32) BindGroupLayout {
	jsLayout := g.jsValue.Call("getBindGroupLayout", index)
	return BindGroupLayout{
		jsValue: jsLayout,
	}
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPipeline) ToJS() any {
	return g.jsValue
}
