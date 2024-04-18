package tax

import "testing"

func TestCalculateTex(t *testing.T) {
	amount := 500.0
	expected := 5.0

	got := CalculateTax(amount)

	if got != expected {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

type testCase struct {
	amount, expected float64
}

func TestCalculateTexBatch(t *testing.T) {
	cases := []testCase{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1001.0, 10.0},
	}

	for _, c := range cases {
		got := CalculateTax(c.amount)

		if got != c.expected {
			t.Errorf("Expected %f, got %f", c.expected, got)
		}
	}
}

// go test -bench=.
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// go test -fuzz=. -fuzztime=3s
func FuzzCalculateTex(f *testing.F) {
	seed := []float64{-1.0, -2, 0, 500.0, 1000.0, 50000.0}

	for _, v := range seed {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		got := CalculateTax(amount)

		if amount <= 0 && got != 0 {
			t.Errorf("expected 0, got %f", got)
		}

		if amount > 20000 && got != 20 {
			t.Errorf("expected 20, got %f", got)
		}
	})
}
