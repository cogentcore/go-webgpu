//go:build js

package wgpu

import (
	"syscall/js"
)

// ComputePassDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucomputepassdescriptor
type ComputePassDescriptor struct{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g *ComputePassDescriptor) ToJS() any {
	return map[string]any{}
}

// ComputePassEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpucomputepassencoder
type ComputePassEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ComputePassEncoder) ToJS() any {
	return g.jsValue
}

// SetPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-setpipeline
func (g ComputePassEncoder) SetPipeline(pipeline ComputePipeline) {
	g.jsValue.Call("setPipeline", pipeline.ToJS())
}

// SetBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubindingcommandsmixin-setbindgroup
func (g ComputePassEncoder) SetBindGroup(index uint32, bindGroup BindGroup, dynamicOffsets []uint32) {
	params := make([]any, 3)
	params[0] = index
	params[1] = bindGroup.ToJS()
	params[2] = dynamicOffsets
	g.jsValue.Call("setBindGroup", params...)
}

// DispatchWorkgroups as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-dispatchworkgroups
func (g ComputePassEncoder) DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ uint32) {
	params := make([]any, 3)
	params[0] = workgroupCountX
	if workgroupCountY > 0 {
		params[1] = workgroupCountY
	} else {
		params[1] = js.Undefined()
	}
	if workgroupCountZ > 0 {
		params[2] = workgroupCountZ
	} else {
		params[2] = js.Undefined()
	}
	g.jsValue.Call("dispatchWorkgroups", params...)
}

// End as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-end
func (g ComputePassEncoder) End() {
	g.jsValue.Call("end")
}
