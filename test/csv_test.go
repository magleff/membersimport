package main_test

import (
	main "github.com/supercoopbdx/membersimport"
	"testing"
)

/*
	MOCKS
*/

type SupercoopMailMock struct {
	GenerateCalled bool
}

func (s *SupercoopMailMock) Generate(firstName string, lastName string) string {
	s.GenerateCalled = true

	if len(firstName) == 0 || len(lastName) == 0 {
		return ""
	} else {
		return "dummy.response@supercoop.fr"
	}
}

/*
	TESTS
*/

func TestProcessRecord(t *testing.T) {
	// Arrange
	generator := &SupercoopMailMock{}

	record := &main.Member{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "06 01 02 03   04",
	}

	// Act
	main.ProcessRecord(generator, record)

	// Assert
	if record.SupercoopMail != "dummy.response@supercoop.fr" {
		t.Error("ProcessRecord: expected SupercoopMail to have been processed")
	}
	if record.Phone != "0601020304" {
		t.Error("ProcessRecord: expected Phone to have been processed")
	}
}
