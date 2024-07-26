//go:build js

package wgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
)

// ComputePassDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucomputepassdescriptor
type ComputePassDescriptor struct{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ComputePassDescriptor) ToJS() any {
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
func (g ComputePassEncoder) SetBindGroup(index Index32, bindGroup BindGroup, dynamicOffsets []BufferDynamicOffset) {
	params := make([]any, 3)
	params[0] = index.ToJS()
	params[1] = bindGroup.ToJS()
	params[2] = gog.Map(dynamicOffsets, func(offset BufferDynamicOffset) any {
		return offset.ToJS()
	})
	g.jsValue.Call("setBindGroup", params...)
}

// DispatchWorkgroups as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-dispatchworkgroups
func (g ComputePassEncoder) DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ Size32) {
	params := make([]any, 3)
	params[0] = workgroupCountX.ToJS()
	if workgroupCountY > 0 {
		params[1] = workgroupCountY.ToJS()
	} else {
		params[1] = js.Undefined()
	}
	if workgroupCountZ > 0 {
		params[2] = workgroupCountZ.ToJS()
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
