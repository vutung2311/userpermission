package usecase_test

import (
	"reflect"
	"testing"

	"userpermission/internal/usecase"
)

func TestProcessPermissionInputWithoutDynamicModification(t *testing.T) {
	input := []string{
		"6",
		"A F",
		"A B",
		"A C E",
		"A",
		"D",
		"A C",
		"A B",
		"CEO",
		"CEO",
		"1",
		"1",
		"1",
		"2",
	}

	expectedOutput := []string{
		"A, B, C, D, E, F",
		"A, B, C, D",
		"A, B, C, E",
		"A",
		"D",
		"A, C",
		"A, B",
	}

	actualOutput, err := usecase.ProcessPermissionInput(input)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Error("actual result is not equal to expected result")
	}
}

func TestProcessPermissionInputWithDynamicModification(t *testing.T) {
	input := []string{
		"6",
		"A F",
		"A B",
		"A C E",
		"A",
		"D",
		"A C",
		"A B",
		"CEO",
		"CEO",
		"1",
		"1",
		"1",
		"2",
		"ADD 2 X",
		"QUERY 2",
		"QUERY CEO",
		"REMOVE 2 X",
		"QUERY 2",
		"QUERY CEO",
	}

	expectedOutput := []string{
		"A, B, C, D, E, F",
		"A, B, C, D",
		"A, B, C, E",
		"A",
		"D",
		"A, C",
		"A, B",
		"A, B, C, E, X",
		"A, B, C, D, E, F, X",
		"A, B, C, E",
		"A, B, C, D, E, F",
	}

	actualOutput, err := usecase.ProcessPermissionInput(input)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Error("actual result is not equal to expected result")
	}
}
