package main

import "testing"

func TestQueryHash(t *testing.T) {
	data := `
xyz: value3
abc: value1
def: value2
`
	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, queryErr := Query(s, "abc")
	if queryErr != nil {
		t.Fatalf("Unexpected error: %v", queryErr)
	}

	expected := "value1"
	if result != expected {
		t.Errorf("Got %v, want %s", result, expected)
	}
}

func TestQueryArray(t *testing.T) {
	data := `
array:
  - value1
  - value2
  - value3
`

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, queryErr := Query(s, "array[1]")
	if queryErr != nil {
		t.Fatalf("Unexpected error: %v", queryErr)
	}

	expected := "value2"
	if result != expected {
		t.Errorf("Got %v, want %s", result, expected)
	}
}

func TestQueryNestedHash(t *testing.T) {
	data := `
japan:
  tokyo:
    capital: tokyo
    population: 13 millions
  kanagawa:
    capital: yokohama
    population: 9 millions
`

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, queryErr := Query(s, "japan.kanagawa.capital")
	if queryErr != nil {
		t.Fatalf("Unexpected error: %v", queryErr)
	}

	expected := "yokohama"
	if result != expected {
		t.Errorf("Got %v, want %s", result, expected)
	}
}

func TestQueryNestedArray(t *testing.T) {
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

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, queryErr := Query(s, "matrix[0][2]")
	if queryErr != nil {
		t.Fatalf("Unexpected error: %v", queryErr)
	}

	expected := "1-3"
	if result != expected {
		t.Errorf("Got %v, want %s", result, expected)
	}
}

func TestQueryArrayOfHash(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
  - name: Jane Doe
    age: 19
`

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, queryErr := Query(s, "members[1].age")
	if queryErr != nil {
		t.Fatalf("Unexpected error: %v", queryErr)
	}

	expected := 19
	if result != expected {
		t.Errorf("Got %v, want %s", result, expected)
	}
}

func TestQueryError(t *testing.T) {
	data := `
members:
  - name: John Doe
    age: 28
    language:
      main: japanese
      sub:
        - english
        - french
  - name: Jane Doe
    age: 19
    language:
      main: english
      sub:
        - russian
`

	s, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if _, queryErr := Query(s, "member"); queryErr == nil {
		t.Errorf("Error did not occured despite path does not exists.")
	}

	if _, queryErr := Query(s, "members[0].language.sub[2]"); queryErr == nil {
		t.Errorf("Error did not occured despite array index is out of range.")
	}

	if _, queryErr := Query(s, "members[0].country"); queryErr == nil {
		t.Errorf("Error did not occured despite hash key is out of range.")
	}

	if _, queryErr := Query(s, "members[0].language.sub"); queryErr == nil {
		t.Errorf("Error did not occured despite path indicates array.")
	}

	if _, queryErr := Query(s, "members[0].language"); queryErr == nil {
		t.Errorf("Error did not occured despite path indicates hash.")
	}
}
