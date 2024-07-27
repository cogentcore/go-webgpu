package wgpu

import "strconv"

// This file contains common types and constants

const (
	ArrayLayerCountUndefined        = 0xffffffff
	CopyStrideUndefined             = 0xffffffff
	LimitU32Undefined        uint32 = 0xffffffff
	LimitU64Undefined        uint64 = 0xffffffffffffffff
	MipLevelCountUndefined          = 0xffffffff
	WholeMapSize                    = ^uint(0)
	WholeSize                       = 0xffffffffffffffff
)

type Version uint32

func (v Version) String() string {
	return "0x" + strconv.FormatUint(uint64(v), 8)
}

type Limits struct {
	MaxTextureDimension1D                     uint32
	MaxTextureDimension2D                     uint32
	MaxTextureDimension3D                     uint32
	MaxTextureArrayLayers                     uint32
	MaxBindGroups                             uint32
	MaxBindingsPerBindGroup                   uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage          uint32
	MaxSamplersPerShaderStage                 uint32
	MaxStorageBuffersPerShaderStage           uint32
	MaxStorageTexturesPerShaderStage          uint32
	MaxUniformBuffersPerShaderStage           uint32
	MaxUniformBufferBindingSize               uint64
	MaxStorageBufferBindingSize               uint64
	MinUniformBufferOffsetAlignment           uint32
	MinStorageBufferOffsetAlignment           uint32
	MaxVertexBuffers                          uint32
	MaxBufferSize                             uint64
	MaxVertexAttributes                       uint32
	MaxVertexBufferArrayStride                uint32
	MaxInterStageShaderComponents             uint32
	MaxInterStageShaderVariables              uint32
	MaxColorAttachments                       uint32
	MaxColorAttachmentBytesPerSample          uint32
	MaxComputeWorkgroupStorageSize            uint32
	MaxComputeInvocationsPerWorkgroup         uint32
	MaxComputeWorkgroupSizeX                  uint32
	MaxComputeWorkgroupSizeY                  uint32
	MaxComputeWorkgroupSizeZ                  uint32
	MaxComputeWorkgroupsPerDimension          uint32

	MaxPushConstantSize uint32
}

// Color as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpucolor
type Color struct {
	R, G, B, A float64
}

type Origin3D struct {
	X, Y, Z uint32
}

type ImageCopyTexture struct {
	Texture  *Texture
	MipLevel uint32
	Origin   Origin3D
	Aspect   TextureAspect
}

type TextureDataLayout struct {
	Offset       uint64
	BytesPerRow  uint32
	RowsPerImage uint32
}

type Extent3D struct {
	Width              uint32
	Height             uint32
	DepthOrArrayLayers uint32
}

type InstanceDescriptor struct {
	Backends           InstanceBackend
	Dx12ShaderCompiler Dx12Compiler
	DxilPath           string
	DxcPath            string
}

// RequestAdapterOptions as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurequestadapteroptions
type RequestAdapterOptions struct {
	CompatibleSurface    *Surface
	PowerPreference      PowerPreference
	ForceFallbackAdapter bool
	BackendType          BackendType
}

type RequiredLimits struct {
	Limits Limits
}

type DeviceDescriptor struct {
	Label              string
	RequiredFeatures   []FeatureName
	RequiredLimits     *RequiredLimits
	DeviceLostCallback DeviceLostCallback
	TracePath          string
}

type DeviceLostCallback func(reason DeviceLostReason, message string)

// TextureDescriptor as described:
// https://gpuweb.github.io/gpuweb/#gputexturedescriptor
type TextureDescriptor struct {
	Label         string
	Usage         TextureUsage
	Dimension     TextureDimension
	Size          Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
}

// BufferDescriptor as described:
// https://gpuweb.github.io/gpuweb/#gpubufferdescriptor
type BufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}

// RenderPassColorAttachment as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpasscolorattachment
type RenderPassColorAttachment struct {
	View          *TextureView
	ResolveTarget *TextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

type CommandEncoderDescriptor struct {
	Label string
}

type TextureViewDescriptor struct {
	Label           string
	Format          TextureFormat
	Dimension       TextureViewDimension
	BaseMipLevel    uint32
	MipLevelCount   uint32
	BaseArrayLayer  uint32
	ArrayLayerCount uint32
	Aspect          TextureAspect
}

type CommandBufferDescriptor struct {
	Label string
}
