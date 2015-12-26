# ml

... incomplete high performacne matrix and vector library (based on float32) using SIMD (at least where SIMD is faster than simple scalar calculations; see benchmarks below)

- mutate matrices/vectors in place
- reimplement (only the ones needed) std math `float32` (instead of converting them to `float64`)
- both scalar and SIMD calculations are implemented, but only the faster one is exported

[Benchmarks](BENCHMARKS.md)

## Implemented

Really empty for now; just getting this project started. New functions will be implemented once I need them.

#### Vec3

function | scalar | amd64 | 386 | arm
-------- | ------ | ----- | --- | -----
Sub      | ✓      | ✓     | ✗   | ✗

#### Vec4

function | scalar | amd64 | 386 | arm
-------- | ------ | ----- | --- | -----
Add      | ✓      | ✓     | ✗   | ✗
Mul      | ✓      | ✓     | ✗   | ✗

#### Mat4

function | scalar | amd64 | 386 | arm
-------- | ------ | ----- | --- | -----
Mul      | ✓      | ✓     | ✗   | ✗


