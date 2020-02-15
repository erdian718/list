package list_test

import (
	"testing"

	"github.com/ofunc/list"
)

func BenchmarkLen(b *testing.B) {
	list.Repeat(0).Take(b.N).Len()
}

func BenchmarkForce(b *testing.B) {
	list.Repeat(0).Take(b.N).Force()
}

func BenchmarkEach(b *testing.B) {
	list.Repeat(0).Take(b.N).Each(func(interface{}) {})
}

func BenchmarkAll(b *testing.B) {
	list.Repeat(0).Take(b.N).All(func(interface{}) bool {
		return true
	})
}

func BenchmarkAny(b *testing.B) {
	list.Repeat(0).Take(b.N).Any(func(interface{}) bool {
		return false
	})
}

func BenchmarkMap(b *testing.B) {
	list.Repeat(0).Map(func(x interface{}) interface{} {
		return x
	}).Take(b.N).Force()
}

func BenchmarkFilter(b *testing.B) {
	list.Repeat(0).Filter(func(x interface{}) bool {
		return true
	}).Take(b.N).Force()
}

func BenchmarkFold(b *testing.B) {
	list.Repeat(0).Take(b.N).Fold(0, func(r, x interface{}) interface{} {
		return r
	})
}

func BenchmarkTake(b *testing.B) {
	list.Repeat(0).Take(b.N).Force()
}

func BenchmarkDrop(b *testing.B) {
	list.Repeat(0).Drop(b.N)
}

func BenchmarkCut(b *testing.B) {
	list.Repeat(0).Cut(b.N)
}

func BenchmarkTakeWhile(b *testing.B) {
	list.Series(0, 1).TakeWhile(func(x interface{}) bool {
		return x.(int) < b.N
	}).Force()
}

func BenchmarkDropWhile(b *testing.B) {
	list.Series(0, 1).DropWhile(func(x interface{}) bool {
		return x.(int) < b.N
	})
}

func BenchmarkCutWhile(b *testing.B) {
	n := b.N / 2
	list.Series(0, 1).Take(b.N).CutWhile(func(x interface{}) bool {
		return x.(int) > n
	}).Force()
}
