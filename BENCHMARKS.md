# Benchmarks

Benchmarks are done on a MacBook Pro (Early 2015, 2.9GHz Intel Core i5)

### Vec3

function | scalar         | SIMD
-------- | -------------- | --------------
Sub      | 3.76 ns/op     | **2.77 ns/op**

### Vec4

function | scalar         | SIMD
-------- | -------------- | --------------
Add      | 4.26 ns/op     | **2.78 ns/op**
Mul      | **2.76 ns/op** | 2.77 ns/op

### Mat4

function | scalar         | SIMD
-------- | -------------- | --------------
Mul      | 29.3 ns/op     | **9.34 ns/op**

function    | result
----------- | ----------
Perspective | 23.3 ns/op

### Math32 (float32 math)

function | math (float64) | math32 (float32)
-------- | -------------- | ----------------
Abs      | 12.2 ns/op     | 12.5 ns/op
Sqrt     | 10.1 ns/op     | 10.3 ns/op
Tan      | 30.2 ns/op     | **25.9 ns/op**
