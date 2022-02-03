package main

type Set struct {
	set map[byte]bool
}

func (set *Set) exists(key byte) bool {
	return set.set[key]
}

func (set *Set) add(key byte) {
	set.set[key] = true
}

func NewSet() Set {
	return Set{make(map[byte]bool)}
}
