package tuple

// This data structure doesn't really have a use in golang, and
// it would be a pain in the ass to make a fully featured tuple
// class from scratch, so... here's a mediocre implementation

type Tuple[A, B any] struct {
	first  A
	second B
}

func NewTuple[A, B any](first A, second B) Tuple[A, B] {
	return Tuple[A, B]{
		first:  first,
		second: second,
	}
}

func (t *Tuple[A, B]) First() A {
	return t.first
}

func (t *Tuple[A, B]) Second() B {
	return t.second
}
