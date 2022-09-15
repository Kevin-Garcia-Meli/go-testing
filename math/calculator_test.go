package main

import "testing"

type subsTest struct {
	arg1, arg2, expect int
}

var subsTests = []subsTest{
	{3, 2, 1},
	{10, 2, 8},
	{20, 5, 15},
}

func TestSubtract(t *testing.T) {
	for _, test := range subsTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expect {
			t.Errorf("Output %q not equal to expected %q", output, test.expect)
		}
	}
}

//--------------------------------------------

type multi struct {
	arg1, arg2, expect int
}

var multiTests = []subsTest{
	{3, 2, 6},
	{10, 2, 20},
	{20, 5, 100},
}

func TestMultiply(t *testing.T) {
	for _, test := range subsTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expect {
			t.Errorf("Output %q not equal to expected %q", output, test.expect)
		}
	}
}

//--------------------------------------------

type divi struct {
	arg1, arg2, expect int
}

var diviTests = []subsTest{
	{4, 2, 2},
	{10, 2, 5},
	{20, 5, 4},
}

func TestDiv(t *testing.T) {
	for _, test := range subsTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expect {
			t.Errorf("Output %q not equal to expected %q", output, test.expect)
		}
	}
}

//-------------------------Benchark

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(4, 3)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(4, 6)
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(4, 6)
	}
}
