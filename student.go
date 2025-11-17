package main

type Student struct {
	FIO    string
	Grades []int
	Avg    float64
}

func (s *Student) calculateAverage() {
	if len(s.Grades) == 0 {
		s.Avg = 0.0
		return
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Avg = float64(sum) / float64(len(s.Grades))
}
