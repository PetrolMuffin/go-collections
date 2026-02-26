package set

import (
	"iter"
	"maps"
)

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{
		items: make(map[T]struct{}),
	}
}

func NewSized[T comparable](capacity int) Set[T] {
	return Set[T]{
		items: make(map[T]struct{}, capacity),
	}
}

func (s *Set[T]) All() iter.Seq[T] {
	return maps.Keys(s.items)
}

func (s *Set[T]) Add(item T) bool {
	if s.Has(item) {
		return false
	}

	s.items[item] = struct{}{}
	return true
}

func (s *Set[T]) Remove(item T) bool {
	if !s.Has(item) {
		return false
	}

	delete(s.items, item)
	return true
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.items))
	for item := range s.items {
		slice = append(slice, item)
	}
	return slice
}

func (s *Set[T]) ToSafe() SyncSet[T] {
	return SyncSet[T]{
		set: s,
	}
}
