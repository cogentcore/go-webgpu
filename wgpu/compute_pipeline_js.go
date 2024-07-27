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

func (g ComputePipelineDescriptor) toJS() any {
	result := make(map[string]any)
	result["layout"] = g.Layout.toJS()
	result["compute"] = g.Compute.toJS()
	return result
}

// ComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#gpucomputepipeline
type ComputePipeline struct {
	jsValue js.Value
}

func (g ComputePipeline) toJS() any {
	return g.jsValue
}
