# gm

... high performance `float32` matrix and vector math library using SIMD for graphics software.

- use SIMD instructions (at least where it makes sense)
- mutate matrices/vectors in place
- use `float32` math (instead of converting values to `float64`)

[Documentation](https://godoc.org/github.com/rkusa/gm) | [Benchmarks](BENCHMARKS.md)

**Status:**
- Feature set is incomplete. I am implementing methods once I need them. PRs are welcome (also, feel free to kindly ask for missing functions).
- API may change.
- ARM falls back to pure go implementations, because of [golang/go#7300](https://github.com/golang/go/issues/7300)
- Any feedback is welcome (regarding performance, implementations, API, ...)

