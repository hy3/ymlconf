package main

import "testing"

func assertPathChain(t *testing.T, result []interface{}, expected []interface{}) {
	if len(result) != len(expected) {
		t.Fatalf("Result has %d properties, want %d", len(result), len(expected))
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("result[%d] = %v, want %v", i, result[i], expected[i])
		}
	}
}

func TestParsePathString(t *testing.T) {
	path := "path.to.property"

	expected := make([]interface{}, 3)
	expected[0] = "path"
	expected[1] = "to"
	expected[2] = "property"

	result := ParsePathString(path)
	assertPathChain(t, result, expected)
}

func TestParsePathStringArray(t *testing.T) {
	path := "path.to.arrays[1]"

	expected := make([]interface{}, 4)
	expected[0] = "path"
	expected[1] = "to"
	expected[2] = "arrays"
	expected[3] = 1

	result := ParsePathString(path)
	assertPathChain(t, result, expected)
}

func TestParsePathStringNestedArray(t *testing.T) {
	path := "path.to.matrix[1][3]"

	expected := make([]interface{}, 5)
	expected[0] = "path"
	expected[1] = "to"
	expected[2] = "matrix"
	expected[3] = 1
	expected[4] = 3

	result := ParsePathString(path)
	assertPathChain(t, result, expected)
}

func TestParsePathStringArrayOfHash(t *testing.T) {
	path := "path.to.arrays[1].member"

	expected := make([]interface{}, 5)
	expected[0] = "path"
	expected[1] = "to"
	expected[2] = "arrays"
	expected[3] = 1
	expected[4] = "member"

	result := ParsePathString(path)
	assertPathChain(t, result, expected)
}
