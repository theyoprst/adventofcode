package containers

type Set[K comparable] map[K]struct{}

func NewSet[K comparable]() Set[K] {
	return Set[K]{}
}

func (s Set[K]) Has(key K) bool {
	_, ok := s[key]
	return ok
}

func (s Set[K]) Add(key K) {
	s[key] = struct{}{}
}

func (s Set[K]) Any() K {
	for key := range s {
		return key
	}
	panic("No values in the set")
}
