package main

import (
	"reflect"
	"testing"
)

func TestSumRange(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{"Sum 0 to 10", 0, 10, 55},
		{"Sum 1 to 5", 1, 5, 15},
		{"Sum single number", 5, 5, 5},
		{"Sum negative range", -2, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumRange(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("SumRange(%d, %d) = %d, want %d", tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestGetZeroValues(t *testing.T) {
	n, s, b, f, arr, m := GetZeroValues()

	if n != 0 {
		t.Errorf("Expected zero value for int: 0, got %d", n)
	}
	if s != "" {
		t.Errorf("Expected zero value for string: \"\", got %q", s)
	}
	if b != false {
		t.Errorf("Expected zero value for bool: false, got %t", b)
	}
	if f != 0.0 {
		t.Errorf("Expected zero value for float64: 0.0, got %f", f)
	}
	if arr != nil {
		t.Errorf("Expected zero value for slice: nil, got %v", arr)
	}
	if m != nil {
		t.Errorf("Expected zero value for map: nil, got %v", m)
	}
}

func TestCreateVariables(t *testing.T) {
	n, s, b, f, arr, m := CreateVariables()

	expectedN := 100
	expectedS := "hello"
	expectedB := true
	expectedF := 3.14
	expectedArr := []int{1, 2, 3}
	expectedM := map[string]int{"apple": 1, "banana": 2}

	if n != expectedN {
		t.Errorf("Expected int: %d, got %d", expectedN, n)
	}
	if s != expectedS {
		t.Errorf("Expected string: %q, got %q", expectedS, s)
	}
	if b != expectedB {
		t.Errorf("Expected bool: %t, got %t", expectedB, b)
	}
	if f != expectedF {
		t.Errorf("Expected float64: %f, got %f", expectedF, f)
	}
	if !reflect.DeepEqual(arr, expectedArr) {
		t.Errorf("Expected slice: %v, got %v", expectedArr, arr)
	}
	if !reflect.DeepEqual(m, expectedM) {
		t.Errorf("Expected map: %v, got %v", expectedM, m)
	}
}