# go-webgpu

> ## **IMPORTANT:** this fork has been moved to https://github.com/cogentcore/webgpu

Go bindings for [`wgpu-native`](https://github.com/gfx-rs/wgpu-native), a cross-platform, safe, graphics api. It runs natively on vulkan, metal, d3d12 and opengles. It also comes with web (JS) support based on https://github.com/mokiat/wasmgpu.

For more info check:
- [webgpu](https://gpuweb.github.io/gpuweb/)
- [wgsl](https://gpuweb.github.io/gpuweb/wgsl/)
- [webgpu-native](https://github.com/webgpu-native/webgpu-headers)

Included static libs are built via [github actions](./.github/workflows/build-wgpu.yml).

## Examples

|[boids][b]|[cube][c]|[triangle][t]|
:-:|:-:|:-:
| [![b-i]][b] | [![c-i]][c] | [![t-i]][t] |

[b-i]: https://raw.githubusercontent.com/rajveermalviya/go-webgpu/main/tests/boids/image-msaa.png
[b]: https://github.com/rajveermalviya/go-webgpu-examples/tree/main/boids
[c-i]: https://raw.githubusercontent.com/rajveermalviya/go-webgpu/main/tests/cube/image-msaa.png
[c]: https://github.com/rajveermalviya/go-webgpu-examples/tree/main/cube
[t-i]: https://raw.githubusercontent.com/rajveermalviya/go-webgpu/main/tests/triangle/image-msaa.png
[t]: https://github.com/rajveermalviya/go-webgpu-examples/tree/main/triangle

you can check out all the examples in [go-webgpu-examples repo](https://github.com/rajveermalviya/go-webgpu-examples)
