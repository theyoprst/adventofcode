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

// Intersection returns a new set with keys that are present in both sets.
func (s Set[K]) Intersection(other Set[K]) Set[K] {
	intersection := NewSet[K]()
	if len(s) > len(other) {
		s, other = other, s // speed up the loop
	}
	for key := range s {
		if other.Has(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

// Difference returns a new set with keys that are present in the first set but not in the second.
func (s Set[K]) Difference(other Set[K]) Set[K] {
	difference := NewSet[K]()
	for key := range s {
		if !other.Has(key) {
			difference.Add(key)
		}
	}
	return difference
}

// Update adds all keys from the other set to the current set.
func (s Set[K]) Update(other Set[K]) {
	for key := range other {
		s[key] = struct{}{}
	}
}

// Union returns a new set with keys that are present in either set.
func (s Set[K]) Union(other Set[K]) Set[K] {
	union := NewSet[K]()
	union.Update(s)
	union.Update(other)
	return union
}
