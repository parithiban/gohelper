package gohelper

import "testing"

func Test_FindInSlice(t *testing.T) {
	stringSlice := []string{"Arthur", "Alexander", "Asher", "Anouar", "Avner"}
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.2, 3.4, 5.6, 7.8, 9.0}
	boolSlice := []bool{true, false}

	tt := []struct {
		name     string
		list     interface{}
		search   interface{}
		expected bool
	}{
		{name: "checks that string data available in slice", list: stringSlice, search: "Arthur", expected: true},
		{name: "checks that string data is not available in slice", list: stringSlice, search: "Burrell", expected: false},
		{name: "checks that int data available in slice", list: intSlice, search: 1, expected: true},
		{name: "checks that int data is not available in slice", list: intSlice, search: 0, expected: false},
		{name: "checks that float data available in slice", list: floatSlice, search: 1.2, expected: true},
		{name: "checks that float data is not available in slice", list: floatSlice, search: 123.456, expected: false},
		{name: "checks that bool data available in slice", list: boolSlice, search: true, expected: true},
		{name: "checks that bool data is not available in slice", list: boolSlice, search: nil, expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			contains := FindInSlice(tc.list, tc.search)

			if contains != tc.expected {
				t.Fatalf("Output should be %t, but got %t", tc.expected, contains)
			}
		})
	}
}
func Test_FindInStruct(t *testing.T) {

	type Person struct {
		Name string
		Age  int
	}

	type PersonPointer struct {
		Name *string
		Age  int
	}

	user1 := "Toby Hyde"
	user2 := "Donald K. Sullivan"
	user1Age := 25
	user2Age := 35

	person1 := Person{
		user1,
		user1Age,
	}

	person2 := Person{
		user2,
		user2Age,
	}

	var persons []Person

	persons = append(persons, person1)
	persons = append(persons, person2)

	pointerPerson1 := PersonPointer{
		&user1,
		user1Age,
	}

	pointerPerson2 := PersonPointer{
		&user2,
		user2Age,
	}

	var pointerPersons []PersonPointer

	pointerPersons = append(pointerPersons, pointerPerson1)
	pointerPersons = append(pointerPersons, pointerPerson2)

	var pointerStruct []*Person

	pointerStruct = append(pointerStruct, &person1)
	pointerStruct = append(pointerStruct, &person2)

	var pointerStructAndValue []*PersonPointer

	pointerStructAndValue = append(pointerStructAndValue, &pointerPerson1)
	pointerStructAndValue = append(pointerStructAndValue, &pointerPerson2)

	tt := []struct {
		name     string
		structs  interface{}
		value    interface{}
		key      string
		expected bool
	}{
		{name: "find string value exists in slice of structs", structs: persons, value: "Donald K. Sullivan", key: "Name", expected: true},
		{name: "find int value exists in slice of structs", structs: persons, value: 35, key: "Age", expected: true},
		{name: "find value that not exists in slice of structs", structs: persons, value: "Arthur", key: "Name", expected: false},
		{name: "find pointer value that exists in slice of structs", structs: pointerPersons, value: "Donald K. Sullivan", key: "Name", expected: true},
		{name: "find pointer value that not exists in slice of structs", structs: pointerPersons, value: "Alexander", key: "Name", expected: false},
		{name: "find value that exists in slice of pointer structs", structs: pointerStruct, value: "Donald K. Sullivan", key: "Name", expected: true},
		{name: "find value that not exists in slice of pointer structs", structs: pointerStruct, value: "Alexander", key: "Age", expected: false},
		{name: "find pointer value that exists in slice of pointer structs", structs: pointerStructAndValue, value: "Donald K. Sullivan", key: "Name", expected: true},
		{name: "find pointer value that not exists in slice of pointer structs", structs: pointerStructAndValue, value: 50, key: "Age", expected: false},
		{name: "find key not exists in slice of struct", structs: pointerStructAndValue, value: 50, key: "Email", expected: false},
		{name: "find value that exists in a struct", structs: person1, value: "Toby Hyde", key: "Name", expected: true},
		{name: "find int value exists in struct", structs: person1, value: 25, key: "Age", expected: true},
		{name: "find value that not exists in a struct", structs: person1, value: "Arthur", key: "Name", expected: false},
		{name: "find pointer value that exists in a struct", structs: pointerPerson1, value: "Toby Hyde", key: "Name", expected: true},
		{name: "find pointer value that not exists in a struct", structs: pointerPerson1, value: "Arthur", key: "Name", expected: false},
		{name: "find value that exists in a pointer struct", structs: &person1, value: "Toby Hyde", key: "Name", expected: true},
		{name: "find value that not exists in a pointer struct", structs: &person1, value: "Arthur", key: "Name", expected: false},
		{name: "find pointer value that exists in pointer struct", structs: &pointerPerson1, value: "Toby Hyde", key: "Name", expected: true},
		{name: "find pointer value that not exists in pointer structs", structs: &pointerPerson1, value: 50, key: "Age", expected: false},
		{name: "find key not exists struct", structs: pointerStructAndValue, value: 50, key: "Email", expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			contains := FindInStruct(tc.structs, tc.value, tc.key)

			if contains != tc.expected {
				t.Fatalf("Output should be %t, but got %t", tc.expected, contains)
			}
		})
	}
}

func Test_EmailValidation(t *testing.T) {
	tt := []struct {
		name     string
		email    string
		expected bool
	}{
		{name: "check email is valid", email: "email@domain.com", expected: true},
		{name: "Email contains dot with subdomain", email: "email@subdomain.domain.com", expected: true},
		{name: "Domain is valid IP address", email: "email@123.123.123.123", expected: true},
		{name: "Digits in address are valid", email: "1234567890@domain.com", expected: true},
		{name: "Missing @ sign and domain", email: "plainaddress", expected: false},
		{name: "Missing username", email: "@domain.com", expected: false},
		{name: "Two @ sign", email: "email@domain@domain.com", expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			expected := EmailValidation(tc.email)

			if expected != tc.expected {
				t.Fatalf("Output should be %t, but got %t", tc.expected, expected)
			}
		})
	}
}
