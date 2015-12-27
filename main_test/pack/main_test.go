package pack

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
func TestAvarage2(t *testing.T) {
	var v float64
	v = Average([]float64{7, 6})
	if v != 12.5 {
		t.Error("Expected 12.5 , got ", v)
	}
}
