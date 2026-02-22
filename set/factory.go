package set

import (
	"errors"
	"fmt"
)

func FromSlice[T comparable](items []T) (Set[T], error) {
	s := New[T]()
	errs := make([]error, 0)
	for _, item := range items {
		if !s.Add(item) {
			errs = append(errs, fmt.Errorf("item %v already exists", item))
		}
	}

	if len(errs) > 0 {
		return Set[T]{}, errors.Join(errs...)
	}

	return s, nil
}

func FromMap[T comparable, K comparable, V any](items map[K]V, selector func(K, V) T) (Set[T], error) {
	s := New[T]()
	errs := make([]error, 0)
	for key, value := range items {
		item := selector(key, value)
		if !s.Add(item) {
			errs = append(errs, fmt.Errorf("item %v already exists", item))
		}
	}

	if len(errs) > 0 {
		return Set[T]{}, errors.Join(errs...)
	}

	return s, nil
}

func FromSliceSafe[T comparable](items []T) (SyncSet[T], error) {
	s, err := FromSlice(items)
	if err != nil {
		return SyncSet[T]{}, err
	}

	return s.ToSafe(), nil
}

func FromMapSafe[T comparable, K comparable, V any](items map[K]V, selector func(K, V) T) (SyncSet[T], error) {
	s, err := FromMap(items, selector)
	if err != nil {
		return SyncSet[T]{}, err
	}

	return s.ToSafe(), nil
}
