package uniformrandom

import (
	"testing"
)

func TestRandomFloat(t *testing.T) {
	proper := []float32{
		0.5430998,
		0.40631828,
		62.147213,
		0.058990162,
	}

	s := Stream{}
	s.SetSeed(72)
	results := []float32 {
		s.RandomFloat(0, 1),
		s.RandomFloat(0, 1),
		s.RandomFloat(0, 100),
		s.RandomFloat(0, 1),
	}

	for i := range proper {
		if proper[i] != results[i] {
			t.Errorf("Improper float result %f != %f", proper[i], results[i])
		}
	}
}

func TestRandomInt(t *testing.T) {
	proper := []int{
		6,
		9,
		95,
		8,
	}

	s := Stream{}
	s.SetSeed(555)

	results := []int{
		s.RandomInt(0, 10),
		s.RandomInt(0, 10),
		s.RandomInt(0, 100),
		s.RandomInt(0, 10),
	}

	for i := range proper {
		if proper[i] != results[i] {
			t.Errorf("Improper int result %d != %d", proper[i], results[i])
		}
	}
}
