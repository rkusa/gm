# Benchmarks

Benchmarks are done on a MacBook Pro (Early 2015, 2.9GHz Intel Core i5)

### Vec3

Initial tests showed that using SIMD for Vec3 is not faster. That is, Vec3
implementations are pure Go only.

function  | result
--------- | ----------
Cross     | 4.34 ns/op
Div       | 5.42 ns/op
Len       | 4.50 ns/op
Normalize | 18.6 ns/op
Sub       | 2.83 ns/op

### Vec4

function | Go             | SIMD
-------- | -------------- | --------------
Add      | 4.26 ns/op     | **2.78 ns/op**
Len      | 4.34 ns/op     | **2.77 ns/op**
Mul      | 2.76 ns/op     | 2.77 ns/op
Sub      | 4.28 ns/op     | **2.75 ns/op**

### Mat4

function | Go             | SIMD
-------- | -------------- | --------------
Mul      | 29.3 ns/op     | **9.34 ns/op**

function    | result
----------- | ----------
LookAt      | 94.4 ns/op
Perspective | 23.3 ns/op
Translate   | 9.46 ns/op

### Math32 (float32 math)

function | math (float64) | math32 (float32)
-------- | -------------- | ----------------
Abs      | 12.2 ns/op     | 12.5 ns/op
Cos      | 26.3 ns/op     | 25.9 ns/op
Sin      | 26.7 ns/op     | 26.1 ns/op
Sqrt     | **9.35 ns/op** | 12.4 ns/op [**!?**](https://github.com/rkusa/ml/issues/1)
Tan      | 30.2 ns/op     | **25.9 ns/op**
