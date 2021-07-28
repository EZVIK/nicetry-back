package main

type TANK interface {
	Fire() int

	Move(d, step int) int
}

type T1 struct {
}

type T2 struct {
}

func (t *T1) Fire() int {
	return 0
}

func (t *T2) Fire() int {
	return 0
}

func (t *T1) Move(d, step int) int {
	return 0
}

func (t *T2) Move(d, step int) int {
	return 0
}
