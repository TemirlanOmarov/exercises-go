package problem3

import "fmt"

type Set struct {
	elements map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		elements: make(map[interface{}]bool),
	}
}

func (s *Set) Add(value interface{}) {
	s.elements[value] = true
}

func (s *Set) Remove(value interface{}) {
	delete(s.elements, value)
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []interface{} {
	keys := make([]interface{}, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set) Has(value interface{}) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for elem := range s.elements {
		newSet.Add(elem)
	}
	return newSet
}

func (s *Set) Difference(other *Set) *Set {
	result := NewSet()
	for elem := range s.elements {
		if !other.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

func (s *Set) IsSubset(other *Set) bool {
	for elem := range s.elements {
		if !other.Has(elem) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	result := NewSet()
	for _, set := range sets {
		for elem := range set.elements {
			result.Add(elem)
		}
	}
	return result
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	result := sets[0].Copy()
	for _, set := range sets[1:] {
		for elem := range result.elements {
			if !set.Has(elem) {
				result.Remove(elem)
			}
		}
	}
	return result
}

func main() {
	setA := NewSet()
	setA.Add(1)
	setA.Add(2)
	setA.Add(3)

	setB := NewSet()
	setB.Add(3)
	setB.Add(4)
	setB.Add(5)

	diff := setA.Difference(setB)
	fmt.Println("Difference (A - B):", diff.List())

	union := Union(setA, setB)
	fmt.Println("Union (A ∪ B):", union.List())

	intersect := Intersect(setA, setB)
	fmt.Println("Intersection (A ∩ B):", intersect.List())

	fmt.Println("Is B a subset of A?", setB.IsSubset(setA))
}
