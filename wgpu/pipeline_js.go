//go:build js

package wgpu

import (
	"syscall/js"
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
		"bindGroupLayouts": mapSlice(g.BindGroupLayouts, func(layout BindGroupLayout) any {
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
	StepMode    VertexStepMode
	Attributes  []VertexAttribute
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g VertexBufferLayout) ToJS() any {
	result := make(map[string]any)
	result["arrayStride"] = g.ArrayStride.ToJS()
	result["stepMode"] = g.StepMode.String()
	result["attributes"] = mapSlice(g.Attributes, func(attrib VertexAttribute) any {
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
		"buffers": mapSlice(g.Buffers, func(layout VertexBufferLayout) any {
			return layout.ToJS()
		}),
	}
}

// PrimitiveState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuprimitivestate
type PrimitiveState struct {
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace        FrontFace
	CullMode         CullMode
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g PrimitiveState) ToJS() any {
	result := make(map[string]any)
	result["topology"] = g.Topology.String()
	result["stripIndexFormat"] = g.StripIndexFormat.String()
	result["frontFace"] = g.FrontFace.String()
	result["cullMode"] = g.CullMode.String()
	return result
}

// StencilFaceState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustencilfacestate
type StencilFaceState struct {
	Compare     CompareFunction
	FailOp      StencilOperation
	DepthFailOp StencilOperation
	PassOp      StencilOperation
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g StencilFaceState) ToJS() any {
	result := make(map[string]any)
	result["compare"] = g.Compare.String()
	result["failOp"] = g.FailOp.String()
	result["depthFailOp"] = g.DepthFailOp.String()
	result["passOp"] = g.PassOp.String()
	return result
}

// DepthStencilState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpudepthstencilstate
type DepthStencilState struct {
	Format              TextureFormat
	DepthWriteEnabled   bool
	DepthCompare        CompareFunction
	StencilFront        StencilFaceState
	StencilBack         StencilFaceState
	StencilReadMask     StencilValue
	StencilWriteMask    StencilValue
	DepthBias           DepthBias
	DepthBiasSlopeScale float32
	DepthBiasClamp      float32
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g DepthStencilState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.String()
	result["depthWriteEnabled"] = g.DepthWriteEnabled
	result["depthCompare"] = g.DepthCompare.String()
	result["stencilFront"] = g.StencilFront.ToJS()
	result["stencilBack"] = g.StencilBack.ToJS()
	result["stencilReadMask"] = g.StencilReadMask.ToJS()
	result["stencilWriteMask"] = g.StencilWriteMask.ToJS()
	result["depthBias"] = g.DepthBias.ToJS()
	result["depthBiasSlopeScale"] = g.DepthBiasSlopeScale
	result["depthBiasClamp"] = g.DepthBiasClamp
	return result
}

// MultisampleState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpumultisamplestate
type MultisampleState struct {
	Count                  Size32
	Mask                   SampleMask
	AlphaToCoverageEnabled bool
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g MultisampleState) ToJS() any {
	result := make(map[string]any)
	result["count"] = g.Count.ToJS()
	result["mask"] = g.Mask.ToJS()
	result["alphaToCoverageEnabled"] = g.AlphaToCoverageEnabled
	return result
}

// BlendComponent as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendcomponent
type BlendComponent struct {
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g BlendComponent) ToJS() any {
	result := make(map[string]any)
	result["operation"] = g.Operation.String()
	result["srcFactor"] = g.SrcFactor.String()
	result["dstFactor"] = g.DstFactor.String()
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
	Blend     BlendState
	WriteMask ColorWriteMask
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ColorTargetState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.String()
	result["blend"] = g.Blend.ToJS()
	result["writeMask"] = g.WriteMask.String()
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
		"targets": mapSlice(g.Targets, func(target ColorTargetState) any {
			return target.ToJS()
		}),
	}
}

// RenderPipelineDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpipelinedescriptor
type RenderPipelineDescriptor struct {
	Layout       PipelineLayout
	Vertex       VertexState
	Primitive    PrimitiveState
	DepthStencil DepthStencilState
	Multisample  MultisampleState
	Fragment     FragmentState
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g RenderPipelineDescriptor) ToJS() any {
	result := make(map[string]any)
	result["layout"] = g.Layout.ToJS()
	result["vertex"] = g.Vertex.ToJS()
	result["primitive"] = g.Primitive.ToJS()
	result["depthStencil"] = g.DepthStencil.ToJS()
	result["multisample"] = g.Multisample.ToJS()
	result["fragment"] = g.Fragment.ToJS()
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
