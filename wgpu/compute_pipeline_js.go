//go:build js

package wgpu

import (
	"syscall/js"
)

// ComputePipelineDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucomputepipelinedescriptor
type ComputePipelineDescriptor struct {
	Layout  PipelineLayout
	Compute ProgrammableStage
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ComputePipelineDescriptor) ToJS() any {
	result := make(map[string]any)
	result["layout"] = g.Layout.ToJS()
	result["compute"] = g.Compute.ToJS()
	return result
}

// ComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#gpucomputepipeline
type ComputePipeline struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ComputePipeline) ToJS() any {
	return g.jsValue
}
