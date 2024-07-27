//go:build js

package wgpu

import "syscall/js"

// ProgrammableStage as described:
// https://gpuweb.github.io/gpuweb/#gpuprogrammablestage
type ProgrammableStage struct {
	Module     *ShaderModule
	EntryPoint string
}

func (g ProgrammableStage) toJS() any {
	return map[string]any{
		"module":     pointerToJS(g.Module),
		"entryPoint": g.EntryPoint,
	}
}

// ShaderModuleDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpushadermoduledescriptor
type ShaderModuleDescriptor struct {
	Label          string
	WGSLDescriptor *ShaderModuleWGSLDescriptor
}

func (g ShaderModuleDescriptor) toJS() any {
	return map[string]any{
		"code": g.WGSLDescriptor.Code,
	}
}

// ShaderModule as described:
// https://gpuweb.github.io/gpuweb/#gpushadermodule
type ShaderModule struct {
	jsValue js.Value
}

func (g ShaderModule) toJS() any {
	return g.jsValue
}

func (g ShaderModule) Release() {} // no-op
