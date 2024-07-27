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

func (g PipelineLayoutDescriptor) toJS() any {
	return map[string]any{
		"bindGroupLayouts": mapSlice(g.BindGroupLayouts, func(layout BindGroupLayout) any {
			return layout.toJS()
		}),
	}
}

// PipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#gpupipelinelayout
type PipelineLayout struct {
	jsValue js.Value
}

func (g PipelineLayout) toJS() any {
	return g.jsValue
}

// VertexAttribute as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexattribute
type VertexAttribute struct {
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

func (g VertexAttribute) toJS() any {
	return map[string]any{
		"format":         toJS(g.Format),
		"offset":         g.Offset,
		"shaderLocation": g.ShaderLocation,
	}
}

// VertexBufferLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexbufferlayout
type VertexBufferLayout struct {
	ArrayStride uint64
	StepMode    VertexStepMode
	Attributes  []VertexAttribute
}

func (g VertexBufferLayout) toJS() any {
	result := make(map[string]any)
	result["arrayStride"] = g.ArrayStride
	result["stepMode"] = toJS(g.StepMode)
	result["attributes"] = mapSlice(g.Attributes, func(attrib VertexAttribute) any {
		return attrib.toJS()
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

func (g VertexState) toJS() any {
	return map[string]any{
		"module":     g.Module.toJS(),
		"entryPoint": g.EntryPoint,
		"buffers": mapSlice(g.Buffers, func(layout VertexBufferLayout) any {
			return layout.toJS()
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

func (g PrimitiveState) toJS() any {
	result := make(map[string]any)
	result["topology"] = toJS(g.Topology)
	result["stripIndexFormat"] = toJS(g.StripIndexFormat)
	result["frontFace"] = toJS(g.FrontFace)
	result["cullMode"] = toJS(g.CullMode)
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

func (g StencilFaceState) toJS() any {
	result := make(map[string]any)
	result["compare"] = toJS(g.Compare)
	result["failOp"] = toJS(g.FailOp)
	result["depthFailOp"] = toJS(g.DepthFailOp)
	result["passOp"] = toJS(g.PassOp)
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
	StencilReadMask     uint32
	StencilWriteMask    uint32
	DepthBias           int32
	DepthBiasSlopeScale float32
	DepthBiasClamp      float32
}

func (g DepthStencilState) toJS() any {
	result := make(map[string]any)
	result["format"] = toJS(g.Format)
	result["depthWriteEnabled"] = g.DepthWriteEnabled
	result["depthCompare"] = toJS(g.DepthCompare)
	result["stencilFront"] = g.StencilFront.toJS()
	result["stencilBack"] = g.StencilBack.toJS()
	result["stencilReadMask"] = g.StencilReadMask
	result["stencilWriteMask"] = g.StencilWriteMask
	result["depthBias"] = g.DepthBias
	result["depthBiasSlopeScale"] = g.DepthBiasSlopeScale
	result["depthBiasClamp"] = g.DepthBiasClamp
	return result
}

// MultisampleState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpumultisamplestate
type MultisampleState struct {
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled bool
}

func (g MultisampleState) toJS() any {
	result := make(map[string]any)
	result["count"] = g.Count
	result["mask"] = g.Mask
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

func (g BlendComponent) toJS() any {
	result := make(map[string]any)
	result["operation"] = toJS(g.Operation)
	result["srcFactor"] = toJS(g.SrcFactor)
	result["dstFactor"] = toJS(g.DstFactor)
	return result
}

// BlendState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendstate
type BlendState struct {
	Color BlendComponent
	Alpha BlendComponent
}

func (g BlendState) toJS() any {
	return map[string]any{
		"color": g.Color.toJS(),
		"alpha": g.Alpha.toJS(),
	}
}

// ColorTargetState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucolortargetstate
type ColorTargetState struct {
	Format    TextureFormat
	Blend     BlendState
	WriteMask ColorWriteMask
}

func (g ColorTargetState) toJS() any {
	result := make(map[string]any)
	result["format"] = toJS(g.Format)
	result["blend"] = g.Blend.toJS()
	result["writeMask"] = toJS(g.WriteMask)
	return result
}

// FragmentState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpufragmentstate
type FragmentState struct {
	Module     ShaderModule
	EntryPoint string
	Targets    []ColorTargetState
}

func (g FragmentState) toJS() any {
	return map[string]any{
		"module":     g.Module.toJS(),
		"entryPoint": g.EntryPoint,
		"targets": mapSlice(g.Targets, func(target ColorTargetState) any {
			return target.toJS()
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

func (g RenderPipelineDescriptor) toJS() any {
	result := make(map[string]any)
	result["layout"] = g.Layout.toJS()
	result["vertex"] = g.Vertex.toJS()
	result["primitive"] = g.Primitive.toJS()
	result["depthStencil"] = g.DepthStencil.toJS()
	result["multisample"] = g.Multisample.toJS()
	result["fragment"] = g.Fragment.toJS()
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

func (g RenderPipeline) toJS() any {
	return g.jsValue
}
