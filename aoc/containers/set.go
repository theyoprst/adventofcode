package containers

type Set[K comparable] map[K]struct{}

func NewSet[K comparable](keys ...K) Set[K] {
	s := Set[K]{}
	s.Add(keys...)
	return s
}

func (s Set[K]) Has(key K) bool {
	_, ok := s[key]
	return ok
}

func (s Set[K]) Add(keys ...K) Set[K] {
	if s == nil {
		s = NewSet[K]()
	}
	for _, key := range keys {
		s[key] = struct{}{}
	}
	return s
}

func (s Set[K]) Remove(keys ...K) {
	for _, key := range keys {
		delete(s, key)
	}
}

func (s Set[K]) Any() K {
	for key := range s {
		return key
	}
	panic("No values in the set")
}

func (s Set[K]) Slice() []K {
	slice := make([]K, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}
