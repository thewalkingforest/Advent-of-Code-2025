package utils

import (
	"testing"
)

func Test_IndexAll_None(t *testing.T) {
	result := IndexAll("...............", "^")
	var expected []int
	if len(result) == len(expected) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != result[i] {
				t.Errorf("IndexAll(..) got incorrect index at %d", i)
			}
		}
	} else {
		t.Errorf("IndexAll(..) got incorrect amount of indexes. expected=%d, got=%d", len(expected), len(result))
	}
}

func Test_IndexAll_One(t *testing.T) {
	result := IndexAll("...^...........", "^")
	expected := []int{3}
	if len(result) == len(expected) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != result[i] {
				t.Errorf("IndexAll(..) got incorrect index at %d", i)
			}
		}
	} else {
		t.Errorf("IndexAll(..) got incorrect amount of indexes. expected=%d, got=%d", len(expected), len(result))
	}
}

func Test_IndexAll_Some(t *testing.T) {
	result := IndexAll("...^^......^...", "^")
	expected := []int{3, 4, 11}
	if len(result) != len(expected) {
		t.Errorf("IndexAll(..) got incorrect amount of indexes. expected=%d, got=%d", len(expected), len(result))
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("IndexAll(..) got incorrect index at %d", i)
		}
	}
}

func Test_IndexAll_One_Long(t *testing.T) {
	result := IndexAll("...^^......^...", "^^")
	expected := []int{3}
	if len(result) != len(expected) {
		t.Errorf("IndexAll(..) got incorrect amount of indexes. expected=%d, got=%d", len(expected), len(result))
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("IndexAll(..) got incorrect index at %d", i)
		}
	}
}

func Test_IndexAll_One_Long_NoOverlap(t *testing.T) {
	result := IndexAll("...^^^.....^...", "^^")
	expected := []int{3}
	if len(result) != len(expected) {
		t.Errorf("IndexAll(..) got incorrect amount of indexes. expected=%d, got=%d", len(expected), len(result))
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("IndexAll(..) got incorrect index at %d", i)
		}
	}
}
