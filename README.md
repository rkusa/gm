# gm

... (incomplete) high performance matrix and vector library (based on float32) using SIMD for graphics software.

- mutate matrices/vectors in place
- reimplement (only the ones needed) std math `float32` (instead of converting values to `float64`)
- both scalar and SIMD calculations are implemented, but only the faster one is exported

[Documentation](https://godoc.org/github.com/rkusa/gm) | [Benchmarks](BENCHMARKS.md)

**Status:**
- SIMD currently only implemented for amd64. Implementations for other architectures will follow soon.
- Feature set is very incomplete. I am implementing methods once I need them. PRs are welcome (feel also free to kindly ask for missing functions).
- API may change.

