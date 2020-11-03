package copy

import (
	"testing"

	"github.com/gotidy/copy"
	"github.com/jinzhu/copier"
	"github.com/ulule/deepcopier"
)

type internal1 struct {
	I int
}
type testStruct1 struct {
	S  string
	I  int
	BB []bool
	V  internal1
}

type internal2 struct {
	I int
}

type testStruct2 struct {
	S  string
	I  int
	BB []bool
	V  internal2
}

var src = testStruct1{
	S:  "string",
	I:  10,
	BB: []bool{true, false},
	V:  internal1{I: 5},
}

var dst testStruct2

func BenchmarkManualCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst = testStruct2{
			S:  src.S,
			I:  src.I,
			BB: src.BB,
			V:  internal2{I: src.V.I},
		}
	}
}

func BenchmarkCopiers(b *testing.B) {
	b.StopTimer()
	c := copy.New()
	c.Prepare(&dst, &src)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c.Copy(&dst, &src)
	}
}

func BenchmarkCopier(b *testing.B) {
	b.StopTimer()
	c := copy.New()
	copier := c.Get(&dst, &src)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		copier.Copy(&dst, &src)
	}
}

func BenchmarkJinzhuCopier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = copier.Copy(&dst, &src)
	}
}

func BenchmarkDeepcopier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = deepcopier.Copy(&src).To(&dst)
	}
}
