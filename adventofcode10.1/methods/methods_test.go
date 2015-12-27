package methods

import "testing"

func TestCalculateFirst1(t *testing.T) {
	var str string = "1113222113"
	res := CalculateFirst(str)
	if res != "3113322113" {
		t.Error("Extected 3113322113,got", res)
	}

}
func TestCalculateFirst2(t *testing.T) {
	var str = "1113222113"

	res := CalculateFirst(str)
	res = CalculateFirst(res)
	if res != "132123222113" {
		t.Error("Expected 132123222113,got", res)
	}
}
