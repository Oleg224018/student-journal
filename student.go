package main

type Student struct {
	Name   string
	Grades []float64
}

func (s *Student) Average() float64 {
	return average(s.Grades)
}
