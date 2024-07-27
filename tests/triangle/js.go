//go:build js

package main

import (
	"time"

	"github.com/rajveermalviya/go-webgpu/wgpu"
)

type window struct{}

func (w window) GetSize() (int, int) { return 512, 512 }

func main() {
	s, err := InitState(&window{}, &wgpu.SurfaceDescriptor{})
	if err != nil {
		panic(err)
	}
	defer s.Destroy()
	ticker := time.NewTicker(time.Second / 60)
	for range ticker.C {
		s.Render()
	}
}
