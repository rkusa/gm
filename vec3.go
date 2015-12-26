package ml

type Vec3 [3]float32

// Sub two 3-dimensional vectors
func (out *Vec3) Sub(rhs *Vec3) {
	vec3SubSIMD(out, rhs)
}
