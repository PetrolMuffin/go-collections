package set

import (
	"iter"
	"sync"
)

type SyncSet[T comparable] struct {
	set *Set[T]
	mu  sync.RWMutex
}

func NewSync[T comparable]() SyncSet[T] {
	baseSet := New[T]()
	return SyncSet[T]{
		set: &baseSet,
	}
}

func NewSyncSized[T comparable](capacity int) SyncSet[T] {
	baseSet := NewSized[T](capacity)
	return SyncSet[T]{
		set: &baseSet,
	}
}

func (s *SyncSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		s.mu.RLock()
		defer s.mu.RUnlock()
		s.set.All()
	}
}

func (s *SyncSet[T]) Add(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Add(item)
}

func (s *SyncSet[T]) Remove(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.set.Remove(item)
}

func (s *SyncSet[T]) Has(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Has(item)
}

func (s *SyncSet[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Len()
}

func (s *SyncSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Clear()
}

func (s *SyncSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.ToSlice()
}
