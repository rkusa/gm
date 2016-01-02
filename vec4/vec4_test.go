package vec4

import (
	"reflect"
	"testing"

	"github.com/rkusa/gm/mat4"
)

func TestNew(t *testing.T) {
	lhs := New(1, 2, 3, 4)

	if !reflect.DeepEqual(lhs, &Vec4{1, 2, 3, 4}) {
		t.Fatalf("New wrong result, got %v", lhs)
	}
}

func TestClone(t *testing.T) {
	a := &Vec4{1, 2, 3, 4}
	b := a.Clone()

	if a == b {
		t.Fatalf("Clone must create a new instance")
	}

	if !reflect.DeepEqual(a, &Vec4{1, 2, 3, 4}) {
		t.Fatalf("Clone must not change values")
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Clone must keep values")
	}
}

func TestTransform(t *testing.T) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := mat4.New(1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16)

	lhs.Transform(rhs)

	expectation := &Vec4{30, 70, 110, 150}
	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Transform wrong result, got %v", lhs)
	}
}
