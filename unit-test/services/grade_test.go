package services_test

import (
	"testing"
	"unit-test/services"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score) // .services เรียกข้าม package
			// expected := c.expected

			if grade != c.expected {
				t.Errorf("got %v expected %v", grade, c.expected)
			}
		})
	}

	// ---- check some test ----
	// go test unit-test/services -run="TestCheckGrade/A" -v

	// t.Run("B", func(t *testing.T) {
	// 	grade := services.CheckGrade(70) // .services เรียกข้าม package
	// 	expectd := "B"

	// 	if grade != expectd {
	// 		t.Errorf("got %v expected %v", grade, expectd)
	// 	}
	// })
}

//  ----- terminal ------
//	1. cd to folder test
//  2. cli "go test"
//  3. check test "go test -v"

// 1. go test <module>/<package>
// 2. go test <module>/<package> -v ===> run every test in that file
// 2. go test <module>/<package> -v -run=<testname> ===> run some test

// go test ./... ===> run every test every where
// go test ./... -v ===> run every test every where
