package main

import "testing"

func TestSetHashReplace(t *testing.T) {
	data := `
xyz: value3
abc: value1
def: value2
`
	expected := make(Properties, 3)
	expected[0] = &Property{"xyz", "value3"}
	expected[1] = &Property{"abc", "value1"}
	expected[2] = &Property{"def", "value4"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "def", "value4")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetHashNew(t *testing.T) {
	data := `
xyz: value3
abc: value1
def: value2
`
	expected := make(Properties, 4)
	expected[0] = &Property{"xyz", "value3"}
	expected[1] = &Property{"abc", "value1"}
	expected[2] = &Property{"def", "value2"}
	expected[3] = &Property{"mno", "value4"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "mno", "value4")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetArrayReplace(t *testing.T) {
	data := `
array:
  - value1
  - value2
  - value3
`

	expected := make(Properties, 3)
	expected[0] = &Property{"array[0]", "value1"}
	expected[1] = &Property{"array[1]", "value4"}
	expected[2] = &Property{"array[2]", "value3"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "array[1]", "value4")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetArrayNew(t *testing.T) {
	data := `
array:
  - value1
  - value2
  - value3
`

	expected := make(Properties, 4)
	expected[0] = &Property{"array[0]", "value1"}
	expected[1] = &Property{"array[1]", "value2"}
	expected[2] = &Property{"array[2]", "value3"}
	expected[3] = &Property{"array[3]", "value4"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "array[3]", "value4")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetArrayNewNested(t *testing.T) {
	data := `
array:
  - value1
  - value2
  - value3
`

	expected := make(Properties, 4)
	expected[0] = &Property{"array[0]", "value1"}
	expected[1] = &Property{"array[1]", "value2"}
	expected[2] = &Property{"array[2]", "value3"}
	expected[3] = &Property{"matrix[0][0][0]", "1-1-1"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "matrix[0][0][0]", "1-1-1")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedHashReplace(t *testing.T) {
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
	expected[1] = &Property{"japan.tokyo.population", "14 millions"}
	expected[2] = &Property{"japan.kanagawa.capital", "yokohama"}
	expected[3] = &Property{"japan.kanagawa.population", "9 millions"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "japan.tokyo.population", "14 millions")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedHashNew(t *testing.T) {
	data := `
japan:
  tokyo:
    capital: tokyo
    population: 13 millions
  kanagawa:
    capital: yokohama
    population: 9 millions
`

	expected := make(Properties, 5)
	expected[0] = &Property{"japan.tokyo.capital", "tokyo"}
	expected[1] = &Property{"japan.tokyo.population", "13 millions"}
	expected[2] = &Property{"japan.kanagawa.capital", "yokohama"}
	expected[3] = &Property{"japan.kanagawa.population", "9 millions"}
	expected[4] = &Property{"japan.aichi.capital", "nagoya"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "japan.aichi.capital", "nagoya")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedHashNewInsert(t *testing.T) {
	data := `
japan:
  tokyo:
    capital: tokyo
    population: 13 millions
  kanagawa:
    capital: yokohama
    population: 9 millions
`

	expected := make(Properties, 5)
	expected[0] = &Property{"japan.tokyo.capital", "tokyo"}
	expected[1] = &Property{"japan.tokyo.population", "13 millions"}
	expected[2] = &Property{"japan.tokyo.flower", "cherry"}
	expected[3] = &Property{"japan.kanagawa.capital", "yokohama"}
	expected[4] = &Property{"japan.kanagawa.population", "9 millions"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "japan.tokyo.flower", "cherry")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedArrayReplace(t *testing.T) {
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
	expected[3] = &Property{"matrix[1][0]", "2-10"}
	expected[4] = &Property{"matrix[1][1]", "2-2"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "matrix[1][0]", "2-10")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedArrayNew(t *testing.T) {
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

	expected := make(Properties, 6)
	expected[0] = &Property{"matrix[0][0]", "1-1"}
	expected[1] = &Property{"matrix[0][1]", "1-2"}
	expected[2] = &Property{"matrix[0][2]", "1-3"}
	expected[3] = &Property{"matrix[1][0]", "2-1"}
	expected[4] = &Property{"matrix[1][1]", "2-2"}
	expected[5] = &Property{"matrix[2][0]", "3-1"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "matrix[2][0]", "3-1")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetNestedArrayNewInsert(t *testing.T) {
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

	expected := make(Properties, 6)
	expected[0] = &Property{"matrix[0][0]", "1-1"}
	expected[1] = &Property{"matrix[0][1]", "1-2"}
	expected[2] = &Property{"matrix[0][2]", "1-3"}
	expected[3] = &Property{"matrix[0][3]", "1-4"}
	expected[4] = &Property{"matrix[1][0]", "2-1"}
	expected[5] = &Property{"matrix[1][1]", "2-2"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "matrix[0][3]", "1-4")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetArrayOfHashReplace(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
  - name: Jane Doe
    age: 19
`

	expected := make(Properties, 4)
	expected[0] = &Property{"members[0].name", "John Doe"}
	expected[1] = &Property{"members[0].age", 30}
	expected[2] = &Property{"members[1].name", "Jane Doe"}
	expected[3] = &Property{"members[1].age", 19}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "members[0].age", 30)
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetArrayOfHashNew(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
  - name: Jane Doe
    age: 19
`

	expected := make(Properties, 5)
	expected[0] = &Property{"members[0].name", "John Doe"}
	expected[1] = &Property{"members[0].age", 28}
	expected[2] = &Property{"members[1].name", "Jane Doe"}
	expected[3] = &Property{"members[1].age", 19}
	expected[4] = &Property{"members[2].name", "Voldemort"}

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	after, setErr := Set(before, "members[2].name", "Voldemort")
	if setErr != nil {
		t.Fatalf("Unexpected error: %v", setErr)
	}
	result := List(after)
	assertProperties(t, result, expected)
}

func TestSetErrorOutOfBounds(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
    job:
      type: programmer
      income: 10000000
  - name: Jane Doe
    age: 19
    job:
      type: student
      income: 0
`

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, setErr := Set(before, "members[3].name", "abc")
	if setErr == nil {
		t.Errorf("Error did not occured despite array out of bounds.")
	}
}

func TestSetErrorNewArrayOutOfBounds(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
    job:
      type: programmer
      income: 10000000
  - name: Jane Doe
    age: 19
    job:
      type: student
      income: 0
`

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, setErr := Set(before, "array[1]", "abc")
	if setErr == nil {
		t.Errorf("Error did not occured despite new array out of bounds.")
	}
}

func TestSetErrorPathIndicatesArray(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
    job:
      type: programmer
      income: 10000000
  - name: Jane Doe
    age: 19
    job:
      type: student
      income: 0
`

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, setErr := Set(before, "members", "abc")
	if setErr == nil {
		t.Errorf("Error did not occured despite path indicates array.")
	}
}

func TestSetErrorPathIndicatesHash(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
    job:
      type: programmer
      income: 10000000
  - name: Jane Doe
    age: 19
    job:
      type: student
      income: 0
`

	before, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, setErr := Set(before, "members[0].job", "abc")
	if setErr == nil {
		t.Errorf("Error did not occured despite path indicates hash.")
	}
}
