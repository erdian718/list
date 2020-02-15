// Package list implements the lazy list data structure.
package list

// List is a lazy list.
type List struct {
	head  interface{}
	tail  *List
	tailf func() *List
}

// New creates a new list.
func New(head interface{}, tailf func() *List) *List {
	return &List{
		head:  head,
		tailf: tailf,
	}
}

// Head returns the head of the list.
func (a *List) Head() interface{} {
	return a.head
}

// Tail returns the tail of the list.
func (a *List) Tail() *List {
	if a.tailf != nil {
		a.tail = a.tailf()
		a.tailf = nil
	}
	return a.tail
}

// Len returns the number of elements of the list.
func (a *List) Len() int {
	n := 0
	for l := a; l != nil; l = l.Tail() {
		n++
	}
	return n
}

// Force calculates the elements of the list immediately, no lazy.
func (a *List) Force() *List {
	for l := a; l != nil; l = l.Tail() {
	}
	return a
}

// Each applies the given f function to each element of the list.
func (a *List) Each(f func(interface{})) *List {
	for l := a; l != nil; l = l.Tail() {
		f(l.head)
	}
	return a
}

// All returns whether all elements of the list match the given f function.
func (a *List) All(f func(interface{}) bool) bool {
	for l := a; l != nil; l = l.Tail() {
		if !f(l.head) {
			return false
		}
	}
	return true
}

// Any returns whether any element of the list matches the given f function.
func (a *List) Any(f func(interface{}) bool) bool {
	for l := a; l != nil; l = l.Tail() {
		if f(l.head) {
			return true
		}
	}
	return false
}

// Cons prepends the element x to the list.
func (a *List) Cons(x interface{}) *List {
	return New(x, func() *List {
		return a
	})
}

// Map applies the given f function to each element of the list and returns the new list.
func (a *List) Map(f func(interface{}) interface{}) *List {
	if a == nil {
		return nil
	}
	return New(f(a.head), func() *List {
		return a.Tail().Map(f)
	})
}

// Filter filters the elements of the list to match the given f function and returns the new list.
func (a *List) Filter(f func(interface{}) bool) *List {
	for l := a; l != nil; l = l.Tail() {
		if f(l.head) {
			return New(l.head, func() *List {
				return l.Tail().Filter(f)
			})
		}
	}
	return nil
}

// Fold applies the f function to each element of the list, threading an accumulator argument a through the computation.
func (a *List) Fold(r interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	for l := a; l != nil; l = l.Tail() {
		r = f(r, l.head)
	}
	return r
}

// Take takes the first n elements.
func (a *List) Take(n int) *List {
	if n <= 0 || a == nil {
		return nil
	}
	return New(a.head, func() *List {
		return a.Tail().Take(n - 1)
	})
}

// Drop drops the first n elements.
func (a *List) Drop(n int) *List {
	for i := 0; i < n && a != nil; i++ {
		a = a.Tail()
	}
	return a
}

// Cut cuts the last x elements.
func (a *List) Cut(n int) *List {
	if n <= 0 {
		return a
	}
	return cutn(a, a.Drop(n))
}

// TakeWhile takes all elements of the list as long as f returns true.
func (a *List) TakeWhile(f func(interface{}) bool) *List {
	if a == nil {
		return nil
	}
	if f(a.head) {
		return New(a.head, func() *List {
			return a.Tail().TakeWhile(f)
		})
	}
	return nil
}

// DropWhile drops all elements of the list as long as f returns true.
func (a *List) DropWhile(f func(interface{}) bool) *List {
	for l := a; l != nil; l = l.Tail() {
		if !f(l.head) {
			return l
		}
	}
	return nil
}

// CutWhile cuts all elements of the list as long as f returns true.
func (a *List) CutWhile(f func(interface{}) bool) *List {
	return cutf(a, a, f)
}

func cutn(a, b *List) *List {
	if b == nil {
		return nil
	}
	return New(a.head, func() *List {
		return cutn(a.Tail(), b.Tail())
	})
}

func cutf(a, b *List, f func(interface{}) bool) *List {
	if a == b {
		b = b.DropWhile(f)
	}
	if b == nil {
		return nil
	}
	return New(a.head, func() *List {
		if a == b {
			b = b.Tail()
		}
		return cutf(a.Tail(), b, f)
	})
}
