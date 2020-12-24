package set

// Set data structure for collecting unique symbols
type Set struct {
	arr []rune
	idx int
}

func (s *Set) push(e rune) {
	for _, elem := range s.arr {
		if e == elem {
			return
		}
	}
	s.arr[s.idx] = e
	s.idx++
}

func (s *Set) contains(e rune) bool {
	for _, elem := range s.arr {
		if e == elem {
			return true
		}
	}
	return false
}

// Size returns size of set - amount of elements in set
func (s *Set) Size() int {
	return s.idx
}

// Intersection calculates intersection between two sets
func (s *Set) Intersection(another Set) Set {
	intersectedSet := Set{make([]rune, 26), 0}
	for _, e := range s.arr {
		if another.contains(e) {
			intersectedSet.push(e)
		}
	}
	return intersectedSet
}

// SetFromString creates Set from string str
func SetFromString(str string) Set {
	s := Set{make([]rune, 26), 0}
	for _, i := range str {
		s.push(i)
	}
	return s
}
