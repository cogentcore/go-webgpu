//go:build js

package wgpu

import "syscall/js"

// ProgrammableStage as described:
// https://gpuweb.github.io/gpuweb/#gpuprogrammablestage
type ProgrammableStage struct {
	Module     ShaderModule
	EntryPoint string
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ProgrammableStage) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
	}
}

// ShaderModuleDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpushadermoduledescriptor
type ShaderModuleDescriptor struct {
	Code string
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ShaderModuleDescriptor) ToJS() any {
	return map[string]any{
		"code": g.Code,
	}
}

// ShaderModule as described:
// https://gpuweb.github.io/gpuweb/#gpushadermodule
type ShaderModule struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g ShaderModule) ToJS() any {
	return g.jsValue
}
