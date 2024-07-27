//go:build js

package wgpu

import "syscall/js"

// ProgrammableStage as described:
// https://gpuweb.github.io/gpuweb/#gpuprogrammablestage
type ProgrammableStage struct {
	Module     ShaderModule
	EntryPoint string
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ProgrammableStage) toJS() any {
	return map[string]any{
		"module":     g.Module.toJS(),
		"entryPoint": g.EntryPoint,
	}
}

// ShaderModuleDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpushadermoduledescriptor
type ShaderModuleDescriptor struct {
	Code string
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ShaderModuleDescriptor) toJS() any {
	return map[string]any{
		"code": g.Code,
	}
}

// ShaderModule as described:
// https://gpuweb.github.io/gpuweb/#gpushadermodule
type ShaderModule struct {
	jsValue js.Value
}

// toJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ShaderModule) toJS() any {
	return g.jsValue
}
