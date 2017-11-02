package main

import "testing"

func assertProperties(t *testing.T, result Properties, expected Properties) {
	if len(result) != len(expected) {
		t.Fatalf("Result has %d properties, want %d", len(result), len(expected))
	}
	for i := 0; i < len(result); i++ {
		if result[i].Path != expected[i].Path {
			t.Errorf("result[%d].path = %s, want %s", i, result[i].Path, expected[i].Path)
		}
		if result[i].Value != expected[i].Value {
			t.Errorf("result[%d].value = %v, want %v", i, result[i].Value, expected[i].Value)
		}
	}
}

func TestListHash(t *testing.T) {
	data := `
xyz: value3
abc: value1
def: value2
`
	expected := make(Properties, 3)
	expected[0] = &Property{"xyz", "value3"}
	expected[1] = &Property{"abc", "value1"}
	expected[2] = &Property{"def", "value2"}

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := List(s)
	assertProperties(t, result, expected)
}

func TestListArray(t *testing.T) {
	data := `
array:
  - value1
  - value2
  - value3
`

	expected := make(Properties, 3)
	expected[0] = &Property{"array[0]", "value1"}
	expected[1] = &Property{"array[1]", "value2"}
	expected[2] = &Property{"array[2]", "value3"}

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := List(s)
	assertProperties(t, result, expected)
}

func TestListNestedHash(t *testing.T) {
	data := `
japan:
  tokyo:
    capital: tokyo
    population: 13 millions
  kanagawa:
    capital: yokohama
    population: 9 millions
`

	expected := make(Properties, 4)
	expected[0] = &Property{"japan.tokyo.capital", "tokyo"}
	expected[1] = &Property{"japan.tokyo.population", "13 millions"}
	expected[2] = &Property{"japan.kanagawa.capital", "yokohama"}
	expected[3] = &Property{"japan.kanagawa.population", "9 millions"}

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := List(s)
	assertProperties(t, result, expected)
}

func TestListNestedArray(t *testing.T) {
	data := `
matrix:
  -
    - 1-1
    - 1-2
    - 1-3
  -
    - 2-1
    - 2-2
`

	expected := make(Properties, 5)
	expected[0] = &Property{"matrix[0][0]", "1-1"}
	expected[1] = &Property{"matrix[0][1]", "1-2"}
	expected[2] = &Property{"matrix[0][2]", "1-3"}
	expected[3] = &Property{"matrix[1][0]", "2-1"}
	expected[4] = &Property{"matrix[1][1]", "2-2"}

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := List(s)
	assertProperties(t, result, expected)
}

func TestListArrayOfHash(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
  - name: Jane Doe
    age: 19
`

	expected := make(Properties, 4)
	expected[0] = &Property{"members[0].name", "John Doe"}
	expected[1] = &Property{"members[0].age", 28}
	expected[2] = &Property{"members[1].name", "Jane Doe"}
	expected[3] = &Property{"members[1].age", 19}

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := List(s)
	assertProperties(t, result, expected)
}
