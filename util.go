package list

// Make makes a list using xs.
func Make(xs ...interface{}) *List {
	if len(xs) == 0 {
		return nil
	}
	return New(xs[0], func() *List {
		return Make(xs[1:]...)
	})
}

// Repeat makes a list with same element x.
func Repeat(x interface{}) *List {
	return New(x, func() *List {
		return Repeat(x)
	})
}

// Range makes a list which elements are x, x+s, x+2s...
func Range(x, s int) *List {
	return New(x, func() *List {
		return Range(x+s, s)
	})
}

// Concat concats the lists.
func Concat(l *List, ls ...*List) *List {
	if len(ls) == 0 {
		return l
	}
	if l == nil {
		return Concat(ls[0], ls[1:]...)
	}
	return New(l.head, func() *List {
		return Concat(l.Tail(), ls...)
	})
}

// Zip zips the lists by function f.
func Zip(f func(xs ...interface{}) interface{}, ls ...*List) *List {
	n := len(ls)
	hs := make([]interface{}, n)
	for i, l := range ls {
		if l == nil {
			return nil
		}
		hs[i] = l.head
	}
	return New(f(hs...), func() *List {
		for i, l := range ls {
			ls[i] = l.Tail()
		}
		return Zip(f, ls...)
	})
}
