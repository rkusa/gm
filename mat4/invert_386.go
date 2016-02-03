package mat4

func invertSIMD(out *Mat4) bool {
	tmp := &Mat4{}
	return invertSIMD386(out, tmp)
}

func invertSIMD386(out, tmp *Mat4) bool
