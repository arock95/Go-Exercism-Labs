package strain

type Ints []int

type Strings []string

type Lists [][]int

func intDecider(i Ints, f func(int) bool) (Ints, Ints) {
	keepers, dumpers := Ints{}, Ints{}
	for _, v := range i {
		if f(v) {
			keepers = append(keepers, v)
		} else {
			dumpers = append(dumpers, v)
		}
	}
	if len(keepers) == 0 {
		keepers = nil
	}
	if len(dumpers) == 0 {
		dumpers = nil
	}
	return keepers, dumpers
}

func (i Ints) Keep(f func(int) bool) Ints {
	keepers, _ := intDecider(i, f)
	return keepers
}

func (i Ints) Discard(f func(int) bool) Ints {
	_, dumpers := intDecider(i, f)
	return dumpers
}

func listDecider(i Lists, f func([]int) bool) (Lists, Lists) {
	keepers, dumpers := Lists{}, Lists{}
	for _, v := range i {
		if f(v) {
			keepers = append(keepers, v)
		} else {
			dumpers = append(dumpers, v)
		}
	}
	if len(keepers) == 0 {
		keepers = nil
	}
	if len(dumpers) == 0 {
		dumpers = nil
	}
	return keepers, dumpers
}

func (l Lists) Keep(f func([]int) bool) Lists {
	keepers, _ := listDecider(l, f)
	return keepers
}

func stringDecider(s Strings, f func(string) bool) (Strings, Strings) {
	keepers, dumpers := Strings{}, Strings{}
	for _, v := range s {
		if f(v) {
			keepers = append(keepers, v)
		} else {
			dumpers = append(dumpers, v)
		}
	}
	if len(keepers) == 0 {
		keepers = nil
	}
	if len(dumpers) == 0 {
		dumpers = nil
	}
	return keepers, dumpers
}

func (s Strings) Keep(f func(string) bool) Strings {
	keepers, _ := stringDecider(s, f)
	return keepers
}
