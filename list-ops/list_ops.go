package listops

// IntList is a slice of integers with custom functions
type IntList []int
type predFunc func(n int) bool
type binFunc func(x, y int) int
type unaryFunc func(x int) int

// Append combines two IntLists
func (i IntList) Append(toAdd IntList) IntList {
	i = append(i, toAdd...)
	return i
}

// Concat combines a variable number of IntLists
func (i IntList) Concat(toAdd []IntList) IntList {
	for _, eachList := range toAdd {
		i = append(i, eachList...)
	}
	return i
}

// Reverse returns an IntList in reverse order
func (i IntList) Reverse() IntList {
	rev := IntList{}

	for itr := len(i) - 1; itr >= 0; itr-- {
		rev = append(rev, i[itr])
	}
	return rev

}

// Length returns the length of an IntList
func (i IntList) Length() int {
	return len(i)
}

// Filter returns only certain elements of an IntList based on a predicate function
func (i IntList) Filter(filter predFunc) IntList {
	filterSlice := IntList{}
	for _, v := range i {
		if filter(v) {
			filterSlice = append(filterSlice, v)
		}
	}
	return filterSlice
}

// Foldl reduces the list starting from the left
func (i IntList) Foldl(f binFunc, initial int) int {
	for _, v := range i {
		initial = (f(initial, v))
	}
	return initial
}

// Foldr reduces the list starting from the right
func (i IntList) Foldr(f binFunc, initial int) int {
	for itr := len(i) - 1; itr >= 0; itr-- {
		initial = f(i[itr], initial)
	}
	return initial
}

// Map performs an action on each element of an IntList
func (i IntList) Map(f unaryFunc) IntList {
	for itr, v := range i {
		i[itr] = f(v)
	}
	return i
}
