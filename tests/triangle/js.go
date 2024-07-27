//go:build js

package main

import "github.com/rajveermalviya/go-webgpu/wgpu"

type window struct{}

func (w window) GetSize() (int, int) { return 512, 512 }

func main() {
	_, err := InitState(&window{}, &wgpu.SurfaceDescriptor{})
	if err != nil {
		panic(err)
	}
	select {}
}
